package biz

import (
	v1 "blug/api/blug/v1"
	"blug/internal/pkg/aiService"
	"context"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

func (uc *BlugUsecase) UploadArticle(url, fileContent string, ctx transporthttp.Context) error {
	title := ctx.Request().FormValue("title")
	tags := ctx.Request().FormValue("tags")
	category := ctx.Request().FormValue("category")

	desc, err := aiService.GenDesc(fileContent)
	if err != nil {
		return err
	}

	return uc.repo.UploadArticleInData(title, desc, category, tags, url, ctx)
}

func (uc *BlugUsecase) GetArticleList(ctx context.Context, offset int) (*v1.GetArticleListResp, error) {
	return uc.repo.GetArticleListInData(ctx, offset)
}

func (uc *BlugUsecase) GetArticle(ctx context.Context, title string) (*v1.GetArticleByTitleResp, error) {
	return uc.repo.GetArticleInData(ctx, title)
}
