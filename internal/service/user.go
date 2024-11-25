package service

import (
	v1 "blug/api/blug/v1"
	"context"
)

func (s *BlugService) RegisterUser(ctx context.Context, req *v1.RegisterUserReq) (*v1.RegisterUserResp, error) {
	s.uc.Log.Infof("RegisterUserInData: %v", req)
	err := s.uc.RegisterUser(ctx, req)
	if err != nil {
		return &v1.RegisterUserResp{Message: "register failed"}, err
	}
	return &v1.RegisterUserResp{Message: "register success"}, nil
}

func (s *BlugService) UserLogin(ctx context.Context, req *v1.UserLoginReq) (*v1.UserLoginResp, error) {
	token, msg := s.uc.UserLogin(ctx, req)
	return &v1.UserLoginResp{Message: msg, Token: token}, nil
}

func (s *BlugService) UserList(ctx context.Context, req *v1.UserListReq) (*v1.UserListResp, error) {
	list, err := s.uc.GetUserList(ctx)
	if err != nil {
		return &v1.UserListResp{Username: []string{}}, err
	}
	return &v1.UserListResp{Username: list}, nil
}
