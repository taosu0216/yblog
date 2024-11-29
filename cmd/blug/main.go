package main

import (
	"blug/daemonSet"
	"blug/internal/conf"
	"blug/internal/pkg"
	"blug/internal/pkg/aiService"
	"blug/internal/pkg/async"
	"blug/internal/pkg/auth"
	"blug/internal/pkg/loggers"
	"blug/internal/pkg/markdown"
	"context"
	"flag"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml.template")
	Name = "blugSvc"
	Version = "dev/v1"
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	currentDir := pkg.GetRootLocation()
	logger := loggers.InitLog(currentDir)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	log.Info(bc.Aiservice.Baseurl, bc.Aiservice.Apikey, bc.Aiservice.Model)
	aiService.InitAiservice(bc.Aiservice.Baseurl, bc.Aiservice.Apikey, bc.Aiservice.Model)
	auth.InitAuth(bc.Auth.Jwtkey)
	markdown.InitRenderer()
	// 加入链路追踪的配置
	err := setTracerProvider(bc.Trace.Endpoint)
	if err != nil {
		panic(err)
	}
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger, bc.Auth)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	async.InitAsynq(bc.Data)
	ctx := context.Background()
	go daemonSet.InitDaemonSet(ctx)

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}

// Set global trace provider 设置链路追逐的方法
func setTracerProvider(url string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("env", "dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
