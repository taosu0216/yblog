package data

import (
	"blug/internal/conf"
	"blug/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBlugRepo)

// Data .
type Data struct {
	DB           *ent.Client
	UserCache    *redis.Client
	FriendCache  *redis.Client
	ArticleCache *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	entClient := newDB(c, logger)
	userCache := newUserCache(c, logger)
	friendCache := newFriendCache(c, logger)
	articleCache := newArticleCache(c, logger)
	return &Data{
		DB:           entClient,
		UserCache:    userCache,
		FriendCache:  friendCache,
		ArticleCache: articleCache,
	}, cleanup, nil
}

func newDB(c *conf.Data, logger log.Logger) *ent.Client {
	cli, err := ent.Open("postgres", c.Database.Source)

	if err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to postgres: %v", err)
		panic(err)
	}
	return cli
}

func newUserCache(c *conf.Data, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.User.Addr,
		Username:     c.User.User,
		Password:     c.User.Password,
		DB:           int(c.User.Db),
		WriteTimeout: c.User.WriteTimeout.AsDuration(),
		ReadTimeout:  c.User.ReadTimeout.AsDuration(),
	})
	return rdb
}
func newFriendCache(c *conf.Data, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Friend.Addr,
		Username:     c.User.User,
		Password:     c.Friend.Password,
		DB:           int(c.Friend.Db),
		WriteTimeout: c.Friend.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Friend.ReadTimeout.AsDuration(),
	})
	return rdb
}
func newArticleCache(c *conf.Data, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Article.Addr,
		Username:     c.User.User,
		Password:     c.Article.Password,
		DB:           int(c.Article.Db),
		WriteTimeout: c.Article.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Article.ReadTimeout.AsDuration(),
	})
	return rdb
}
