package pkg

import "errors"

const (
	UserLoginFailedMsg  = "login failed"
	UserLoginSuccessMsg = "login success"

	UserNotFoundMsg    = "user not found"
	ArticleNotFoundMsg = "article not found"

	InternalErrMsg = "internal error"

	UploadSuccessMsg = "upload success"
	UploadFailedMsg  = "upload failed"

	RateLimitMsg = "rate limit"
	BlackListMsg = "you are banned"
)

// redis
const (
	LinkListKey    = "link_list"
	UserListKey    = "user_list"
	ArticleListKey = "article_list"
	ArticleMapKey  = "article_map"
)

const (
	NotFoundCode    = 404
	InternalErrCode = 500
)

var (
	UserNotFoundErr     = errors.New("user not found")
	ArticleNotFoundErr  = errors.New("article not found")
	InternalErr         = errors.New("internal error")
	AuthFailedErr       = errors.New("auth failed")
	PermissionDeniedErr = errors.New("permission denied")
	RateLimitErr        = errors.New("rate limit")
)
