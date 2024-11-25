package data

import (
	v1 "blug/api/blug/v1"
	"blug/internal/pkg"
	"context"
	"log"
)

type linkObj struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Desc   string `json:"desc"`
	Avatar string `json:"avatar"`
}

func (f *blugRepo) CreateNewFriendLinkInData(ctx context.Context, req *v1.CreateNewFriendLinkReq) error {
	err := f.data.DB.Friend.Create().
		SetTitle(req.Title).
		SetLink(req.Link).
		SetDesc(req.Desc).
		SetAvatar(req.Avatar).
		Exec(ctx)
	if err != nil {
		f.log.Error(err)
		return err
	}
	link := &linkObj{
		Title:  req.Title,
		Link:   req.Link,
		Desc:   req.Desc,
		Avatar: req.Avatar,
	}

	linkStr, err := pkg.AnyToJsonStr(link)
	if err != nil {
		f.log.Error(err)
	} else {
		f.data.FriendCache.RPush(ctx, pkg.LinkListKey, linkStr)
		log.Println("link str", linkStr, " to cache succeed")
	}

	return nil
}

func (f *blugRepo) GetFriendLinkListInData(ctx context.Context) (*v1.GetFriendLinkListResp, error) {
	result, err := f.data.FriendCache.LRange(ctx, pkg.LinkListKey, 0, -1).Result()
	if err != nil {
		f.log.Error(err)
	} else {
		links := make([]linkObj, 0)
		err = pkg.JsonStrSliceToAny(result, &links)
		if err != nil {
			f.log.Error(err)
		} else {
			f.log.Info("get link list shoot cache succeed!")
			return linkObjToV1LinkListResp(links), nil
		}
	}

	friendLinks, err := f.data.DB.Friend.Query().All(ctx)
	if err != nil {
		f.log.Error(err)
		return nil, pkg.InternalErr
	}
	resp := make([]*v1.Link, 0)
	for _, friendLink := range friendLinks {
		resp = append(resp, &v1.Link{
			Title:  friendLink.Title,
			Link:   friendLink.Link,
			Desc:   friendLink.Desc,
			Avatar: friendLink.Avatar,
		})
	}
	return &v1.GetFriendLinkListResp{Links: resp}, nil
}

func linkObjToV1Link(link linkObj) *v1.Link {
	return &v1.Link{
		Title:  link.Title,
		Link:   link.Link,
		Desc:   link.Desc,
		Avatar: link.Avatar,
	}
}
func linkObjToV1LinkListResp(links []linkObj) *v1.GetFriendLinkListResp {
	resp := make([]*v1.Link, 0)
	for _, link := range links {
		resp = append(resp, linkObjToV1Link(link))
	}
	return &v1.GetFriendLinkListResp{Links: resp}
}
