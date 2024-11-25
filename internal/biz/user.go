package biz

import (
	v1 "blug/api/blug/v1"
	"blug/internal/pkg"
	"blug/internal/pkg/auth"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

func (uc *BlugUsecase) RegisterUser(ctx context.Context, req *v1.RegisterUserReq) error {
	return uc.repo.RegisterUserInData(ctx, req)
}

func (uc *BlugUsecase) UserLogin(ctx context.Context, req *v1.UserLoginReq) (string, string) {
	isUser, err := uc.repo.CheckUserInDB(ctx, req)
	if err != nil {
		if errors.Is(err, pkg.UserNotFoundErr) {
			return "", pkg.UserNotFoundMsg
		}
		return "", pkg.UserLoginFailedMsg
	}
	if isUser {
		token, err := auth.Auth(req.Username)
		if err != nil {
			return "", pkg.InternalErrMsg
		}
		return token, pkg.UserLoginSuccessMsg
	} else {
		return "", pkg.UserLoginFailedMsg
	}
}

func (uc *BlugUsecase) GetUserList(ctx context.Context) ([]string, error) {
	value := ctx.Value("username")
	if value != "root" {
		log.Infof("username is: %v", value)
		return nil, pkg.PermissionDeniedErr
	}
	return uc.repo.GetUserListInData(ctx)
}
