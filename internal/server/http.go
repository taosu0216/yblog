package server

import (
	v1 "blug/api/blug/v1"
	"blug/internal/conf"
	"blug/internal/pkg"
	"blug/internal/pkg/auth"
	"blug/internal/pkg/cache"
	"blug/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strings"
	"sync"

	netHttp "net/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, bs *service.BlugService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//logging.Server(logger),
			validate.Validator(), // 接口访问的参数校验
			selector.Server(
				auth.NewAuthServer(),
			).Match(NewWhiteListMatcher()).Build(),
			tracing.Server(), // 新增 tracing
			MiddlewareCors(), // 跨域请求头
		),
		http.Filter(rateLimitFilter()),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterBlugHTTPServer(srv, bs)

	{
		// 手动添加的路由
		route := srv.Route("/")
		// 上传文章
		{
			route.POST("/api/articles/upload", bs.Upload)
		}
		// 前端页面
		{
			// 首页 { 头像 | 文章列表 }
			route.GET("/static/index/page", bs.GetIndex)
			//route.GET("/static/avatar", bs.GetAvatar)
			route.GET("/static/articles/{article}", bs.GetArticle)
		}
	}

	return srv
}

// NewWhiteListMatcher 设置白名单，不需要 token 验证的接口
func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/blug.v1.Blug/RegisterUser"] = struct{}{}
	whiteList["/blug.v1.Blug/UserLogin"] = struct{}{}
	whiteList["/blug.v1.Blug/CreateNewFriendLink"] = struct{}{}
	whiteList["/blug.v1.Blug/GetFriendLinkList"] = struct{}{}
	whiteList["/blug.v1.Blug/GetArticleList"] = struct{}{}
	whiteList["/blug.v1.Blug/GetArticleByTitle"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		ip := ""
		if header, ok := transport.FromServerContext(ctx); ok {
			ip = header.RequestHeader().Get("X-Forwarded-For")
			if ip == "" {
				ip = header.RequestHeader().Get("X-RemoteAddr")
			}
		} else {
			ip = header.RequestHeader().Get("X-RemoteAddr")
		}
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func rateLimitFilter() http.FilterFunc {
	ipCounterFunc := cache.IpCounter()
	blackIpList := sync.Map{}
	var mu sync.Mutex
	for _, v := range pkg.GetBlackList() {
		blackIpList.Store(v, true)
	}
	return func(next netHttp.Handler) netHttp.Handler {
		return netHttp.HandlerFunc(func(w netHttp.ResponseWriter, req *netHttp.Request) {
			ip := ""
			ip = req.Header.Get("X-Real-Ip")
			log.Info("1. X-Real-Ip: ", ip)
			if ip == "" {
				ip = req.Header.Get("X-Forwarded-For")
				log.Info("2. X-Forwarded-For: ", ip)
				if ip == "" {
					ip = strings.Split(req.RemoteAddr, ":")[0]
					log.Info("3. RemoteAddr: ", ip)
				}
			}
			if value, ok := blackIpList.Load(ip); ok && value == true {
				netHttp.Error(w, pkg.BlackListMsg, 403)
				return
			}
			log.Infof("req ip:%s, operation is:%s", ip, req.URL.Path)
			req.Header.Add("X-RemoteAddr", strings.Split(req.RemoteAddr, ":")[0])
			sum := ipCounterFunc(ip)
			if sum >= 30 {
				mu.Lock()
				if _, exists := blackIpList.Load(ip); !exists {
					blackIpList.Store(ip, true)
					pkg.UpdateBlackListFile(ip)
				}
				mu.Unlock()
				netHttp.Error(w, pkg.RateLimitMsg, 429)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func blackListIpFilter() http.FilterFunc {
	return func(next netHttp.Handler) netHttp.Handler {
		return netHttp.HandlerFunc(func(w netHttp.ResponseWriter, req *netHttp.Request) {
			req.Header.Add("X-RemoteAddr", strings.Split(req.RemoteAddr, ":")[0])
			next.ServeHTTP(w, req)
		})
	}
}

// MiddlewareCors 设置跨域请求头
func MiddlewareCors() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if ts, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := ts.(http.Transporter); ok {
					ht.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
					ht.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
					ht.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
					ht.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,"+
						"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")
				}
			}
			return handler(ctx, req)
		}
	}
}
func logIP() func(handler middleware.Handler) middleware.Handler {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			ip := ""
			if header, ok := transport.FromServerContext(ctx); ok {
				ip = header.RequestHeader().Get("X-Forwarded-For")
				if ip == "" {
					ip = header.RequestHeader().Get("X-RemoteAddr")
				}
			} else {
				ip = header.RequestHeader().Get("X-RemoteAddr")
			}
			reply, err = handler(ctx, req)
			return
		}
	}
}
func ipCounter() func(handler middleware.Handler) middleware.Handler {
	ipCounterFunc := cache.IpCounter()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			ip := ctx.Value("ip")
			ipCounterFunc(ip.(string))
			reply, err = handler(ctx, req)
			return
		}
	}
}
