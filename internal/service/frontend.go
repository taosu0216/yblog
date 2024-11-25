package service

import (
	v1 "blug/api/blug/v1"
	"blug/internal/pkg"
	"errors"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"net/http"
)

func (s *BlugService) GetIndex(ctx transporthttp.Context) error {
	listsResp, _ := s.GetArticleList(ctx, &v1.GetArticleListReq{})

	// 解析模板文件
	tmpl, err := template.ParseFiles(pkg.GetRootLocation() + "dist/index.html")
	if err != nil {
		http.Error(ctx.Response(), err.Error(), http.StatusInternalServerError)
		return nil
	}

	// 渲染模板并写入 HTTP 响应
	err = tmpl.Execute(ctx.Response(), map[string]interface{}{
		"title":   "index",
		"infos":   listsResp.Articles,
		"avatar":  "https://s3.bmp.ovh/imgs/2024/11/22/6c668980b1e2aa2f.jpg",
		"pageNum": len(listsResp.Articles) / 5,
	})
	if err != nil {
		http.Error(ctx.Response(), err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func (s *BlugService) GetArticle(ctx transporthttp.Context) error {

	articleResp, err := s.GetArticleByTitle(ctx, &v1.GetArticleByTitleReq{Title: ctx.Vars()["article"][0]})
	if err != nil {
		if errors.Is(err, pkg.ArticleNotFoundErr) {
			notFoundMdContent := blackfriday.Run([]byte(articleResp.Article.Content), blackfriday.WithExtensions(
				blackfriday.CommonExtensions,
			))
			tmpl, err3 := template.ParseFiles(pkg.GetRootLocation() + "dist/article.html")
			if err3 != nil {
				http.Error(ctx.Response(), err.Error(), http.StatusInternalServerError)
				return nil
			}
			err = tmpl.Execute(ctx.Response(), map[string]interface{}{
				"title":   "my blug",
				"info":    articleResp.Article,
				"content": template.HTML(notFoundMdContent),
			})
			if err != nil {
				return err
			}
			return nil
		} else if errors.Is(err, pkg.InternalErr) {
			internalErrContent := blackfriday.Run([]byte(articleResp.Article.Content), blackfriday.WithExtensions(
				blackfriday.CommonExtensions,
			))
			tmpl, err3 := template.ParseFiles(pkg.GetRootLocation() + "dist/article.html")
			if err3 != nil {
				http.Error(ctx.Response(), err.Error(), http.StatusInternalServerError)
				return nil
			}
			err = tmpl.Execute(ctx.Response(), map[string]interface{}{
				"title":   "my blug",
				"info":    articleResp.Article,
				"content": template.HTML(internalErrContent),
			})
			if err != nil {
				return err
			}
			return nil
		}
	}
	tmpl, err2 := template.ParseFiles(pkg.GetRootLocation() + "dist/article.html")
	if err2 != nil {
		http.Error(ctx.Response(), err2.Error(), http.StatusInternalServerError)
		return nil
	}

	markdownContent := blackfriday.Run([]byte(articleResp.Article.Content), blackfriday.WithExtensions(
		blackfriday.CommonExtensions,
	))

	// 渲染模板并写入 HTTP 响应
	err = tmpl.Execute(ctx.Response(), map[string]interface{}{
		"title":   "my blug",
		"info":    articleResp.Article,
		"content": template.HTML(markdownContent),
	})
	if err != nil {
		http.Error(ctx.Response(), err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func (s *BlugService) GetAvatar(ctx transporthttp.Context) error {
	// 获取图片文件路径
	avatarPath := pkg.GetRootLocation() + "dist/avatar.jpg"

	fileBytes, err := ioutil.ReadFile(avatarPath)
	if err != nil {
		panic(err)
	}
	ctx.Response().WriteHeader(http.StatusOK)
	ctx.Response().Header().Set("Content-Type", "application/octet-stream")
	ctx.Response().Write(fileBytes)
	return nil
}
