package biz

import (
	v1 "blug/api/blug/v1"
	"blug/internal/pkg/aiService"
	"context"
)

func (uc *BlugUsecase) CreateNewFriendLinkReq(ctx context.Context, req *v1.CreateNewFriendLinkReq) (*v1.CreateNewFriendLinkResp, error) {
	uc.Log.Infof("CreateNewFriendLinkReq: %v", req)
	status, msg, err := aiService.VerifyFriendLink(req.Link, req.Title, req.Desc)
	if err != nil || status == 0 {
		uc.Log.Errorf("CreateNewFriendLinkReq VerifyFriendLink err: %v", err)
		return &v1.CreateNewFriendLinkResp{
			Message: "failed",
			Check: &v1.Result{
				Status: "failed",
				Msg:    msg,
			},
		}, err
	}

	return &v1.CreateNewFriendLinkResp{
			Message: "success",
			Check: &v1.Result{
				Status: "success",
				Msg:    "success",
			},
		}, uc.repo.CreateNewFriendLinkInData(ctx, &v1.CreateNewFriendLinkReq{
			Title:  req.Title,
			Link:   req.Link,
			Desc:   req.Desc,
			Avatar: req.Avatar,
		})
}

func (uc *BlugUsecase) GetFriendLinkListReq(ctx context.Context, req *v1.GetFriendLinkListReq) (*v1.GetFriendLinkListResp, error) {
	return uc.repo.GetFriendLinkListInData(ctx)
}
