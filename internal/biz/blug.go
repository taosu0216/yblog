package biz

import (
	v1 "blug/api/blug/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

type BlugRepo interface {
	// friend service
	CreateNewFriendLinkInData(ctx context.Context, req *v1.CreateNewFriendLinkReq) error
	GetFriendLinkListInData(ctx context.Context) (*v1.GetFriendLinkListResp, error)

	// user service
	RegisterUserInData(ctx context.Context, req *v1.RegisterUserReq) error
	CheckUserInDB(ctx context.Context, req *v1.UserLoginReq) (bool, error)
	GetUserListInData(ctx context.Context) ([]string, error)

	// article service
	UploadArticleInData(title, desc, category, tags, url string, ctx transporthttp.Context) error
	GetArticleListInData(ctx context.Context, offset int) (*v1.GetArticleListResp, error)
	GetArticleInData(ctx context.Context, title string) (*v1.GetArticleByTitleResp, error)
}

type BlugUsecase struct {
	repo BlugRepo
	Log  *log.Helper
}

func NewBlugUsecase(repo BlugRepo, logger log.Logger) *BlugUsecase {
	return &BlugUsecase{repo: repo, Log: log.NewHelper(logger)}
}
