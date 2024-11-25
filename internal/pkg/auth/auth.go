package auth

import (
	"blug/internal/pkg"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

var jwtKey string

func InitAuth(key string) {
	jwtKey = key
}

type MyClaims struct {
	UserName string `json:"username"`
	jwtv4.StandardClaims
}

func Auth(username string) (string, error) {
	myClaims := MyClaims{
		username,
		jwtv4.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), //设置JWT过期时间,此处设置为2小时
			Issuer:    "blug.yb",                            //设置签发人
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	//加盐
	return claims.SignedString([]byte(jwtKey))
}

// CheckJWT 解析
func CheckJWT(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		result := make(map[string]interface{}, 2)
		result["username"] = claims["username"]
		result["differentiate"] = claims["differentiate"]
		return result, nil
	} else {
		return nil, errors.New("token type error")
	}
}

// NewAuthServer jwt Server中间件
func NewAuthServer() func(handler middleware.Handler) middleware.Handler {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var jwtToken string
			if md, ok := metadata.FromIncomingContext(ctx); ok {
				jwtToken = md.Get("x-md-global-jwt")[0]
			} else if header, ok := transport.FromServerContext(ctx); ok {
				jwtToken = strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)[1]
			} else {
				// 缺少可认证的token，返回错误
				return nil, pkg.AuthFailedErr
			}
			token, err := CheckJWT(jwtToken)
			if err != nil {
				// 缺少合法的token，返回错误
				return nil, pkg.AuthFailedErr
			}
			ctx = context.WithValue(ctx, "username", token["username"])
			ctx = context.WithValue(ctx, "differentiate", token["differentiate"])
			reply, err = handler(ctx, req)
			return
		}
	}
}

func IsRoot(ctx context.Context) bool {
	name := ctx.Value("username")
	if name != nil {
		if name.(string) == "root" {
			return true
		}
	}
	return false
}
