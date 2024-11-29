package async

import (
	"blug/internal/conf"
	"blug/internal/data/ent"
	"blug/internal/pkg"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type AsynqClient struct {
	AsynqClient *asynq.Client
	Pg          *ent.Client
}

type AsynqServer struct {
	AsynqDefaultServer *asynq.Server
	AsynqDeathServer   *asynq.Server
	RedisClient        *redis.Client
	Pg                 *ent.Client
}

type AsynqInspector struct {
	AsynqInspector *asynq.Inspector
}

var asynqClient AsynqClient
var asynqServer AsynqServer
var asyncInspector AsynqInspector

func InitAsynq(bc *conf.Data) {
	initAsynqClient(bc)
	initAsynqServer(bc)
	initAsyncInspector(bc.Article)
}

func initAsynqClient(bc *conf.Data) {
	asynqClient = AsynqClient{}
	asynqClient.AsynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: bc.Article.Addr, Username: bc.Article.User, Password: bc.Article.Password})
	var err error
	asynqClient.Pg, err = ent.Open("postgres", bc.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		panic(err)
	}
}

func initAsynqServer(bc *conf.Data) {
	asynqServer = AsynqServer{}
	asynqServer.AsynqDefaultServer = asynq.NewServer(
		asynq.RedisClientOpt{Addr: bc.Article.Addr, Username: bc.Article.User, Password: bc.Article.Password},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				pkg.DEFAULT_QUEUE: 1,
			},
		},
	)
	asynqServer.AsynqDeathServer = asynq.NewServer(
		asynq.RedisClientOpt{Addr: bc.Article.Addr, Username: bc.Article.User, Password: bc.Article.Password},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				pkg.DEATH_QUEUE: 1,
			},
		},
	)
	asynqServer.RedisClient = redis.NewClient(&redis.Options{
		Addr:         bc.Article.Addr,
		Username:     bc.Article.User,
		Password:     bc.Article.Password,
		DB:           int(bc.Article.Db),
		WriteTimeout: bc.Article.WriteTimeout.AsDuration(),
		ReadTimeout:  bc.Article.ReadTimeout.AsDuration(),
	})

	var err error
	asynqServer.Pg, err = ent.Open("postgres", bc.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		panic(err)
	}
}

func initAsyncInspector(bc *conf.Data_Redis) {
	asyncInspector = AsynqInspector{}
	asyncInspector.AsynqInspector = asynq.NewInspector(asynq.RedisClientOpt{Addr: bc.Addr, Username: bc.User, Password: bc.Password})
}

func GetAsynqClient() AsynqClient {
	return asynqClient
}

func GetAsynqServer() AsynqServer {
	return asynqServer
}

func GetAsyncInspector() AsynqInspector {
	return asyncInspector
}
