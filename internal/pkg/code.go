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

const (
	TASK_FLUSH_CACHE      = "task:flush-cache"
	TASK_GET_MACHINE_INFO = "task:get-machine-info"
	// todo: 模仿getMachineInfoFunc 制作获取网络带宽任务

	STATUS_INITIALIZING = "initializing"
	STATUS_RUNNING      = "running"
	STATUS_SUCCESS      = "success"
	STATUS_FAIL         = "fail"
	STATUS_CANCEL       = "cancel"

	DEFAULT_QUEUE = "default"
	DEATH_QUEUE   = "death"
)
