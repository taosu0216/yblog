package main

import (
	"blug/daemonSet"
	"blug/internal/conf"
	"blug/internal/pkg/async"
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"os"
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
	flag.StringVar(&flagconf, "conf", "../../configs/worker", "config path, eg: -conf config.yaml.template")
	Name = "blugSvc"
	Version = "dev/v1"
}

func main() {
	flag.Parse()

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
	async.InitAsynq(bc.Data)
	if bc.Daemon.IsInit {
		ctx := context.Background()
		daemonSet.InitDaemonSet(ctx, bc.Daemon.IsController, int(bc.Daemon.Skip))
	}
}
