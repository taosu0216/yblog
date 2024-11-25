package service

import (
	pb "blug/api/blug/v1"
	"context"
)

func (s *BlugService) CreateNewFriendLink(ctx context.Context, req *pb.CreateNewFriendLinkReq) (*pb.CreateNewFriendLinkResp, error) {
	res, err := s.uc.CreateNewFriendLinkReq(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateNewFriendLinkResp{Message: res.Message, Check: res.Check}, nil
}

func (s *BlugService) GetFriendLinkList(ctx context.Context, req *pb.GetFriendLinkListReq) (*pb.GetFriendLinkListResp, error) {
	res, err := s.uc.GetFriendLinkListReq(ctx, req)
	if err != nil {
		return &pb.GetFriendLinkListResp{}, err
	}
	return res, nil
}
