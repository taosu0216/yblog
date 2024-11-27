package async

import (
	"blug/internal/conf"
	"blug/internal/data/ent"
	"blug/internal/pkg"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type AsynqServer struct {
	*asynq.Server
	RedisClient *redis.Client
	Pg          *ent.Client
}

var asynqClient *asynq.Client
var asynqServer AsynqServer
var asyncInspector *asynq.Inspector

func InitAsynq(bc *conf.Data, logger log.Logger) {
	initAsynqClient(bc.Article, logger)
	initAsynqServer(bc, logger)
	initAsyncInspector(bc.Article, logger)
}

func initAsynqClient(bc *conf.Data_Redis, logger log.Logger) {
	asynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: bc.Addr, Username: bc.User, Password: bc.Password})
}

func initAsynqServer(bc *conf.Data, logger log.Logger) {
	asynqServer.Server = asynq.NewServer(
		asynq.RedisClientOpt{Addr: bc.Article.Addr, Username: bc.Article.User, Password: bc.Article.Password},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				pkg.DEFAULT_QUEUE: 1,
				pkg.DEATH_QUEUE:   1,
			},
			// See the godoc for other configuration options
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
		log.NewHelper(logger).Fatalf("failed opening connection to postgres: %v", err)
		panic(err)
	}
}

func initAsyncInspector(bc *conf.Data_Redis, logger log.Logger) {
	asyncInspector = asynq.NewInspector(asynq.RedisClientOpt{Addr: bc.Addr, Username: bc.User, Password: bc.Password})
}

func GetAsynqClient() *asynq.Client {
	return asynqClient
}

func GetAsynqServer() AsynqServer {
	return asynqServer
}

func GetAsyncInspector() *asynq.Inspector {
	return asyncInspector
}
