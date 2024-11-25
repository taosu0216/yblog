package data

import (
	v1 "blug/api/blug/v1"
	"blug/internal/data/ent"
	"blug/internal/data/ent/article"
	"blug/internal/pkg"
	"context"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"log"
	"time"
)

type articleObj struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Desc        string `json:"desc"`
	Tags        string `json:"tags"`
	Category    string `json:"category"`
	CreatedTime string `json:"created_time"`
}

func (f *blugRepo) UploadArticleInData(title, desc, category, tags, url string, ctx transporthttp.Context) error {
	err := f.data.DB.Article.Create().
		SetTitle(title).
		SetDesc(desc).
		SetCategory(category).
		SetTags(tags).SetURL(url).SetCreateTime(time.Now().Format("2006-01-02")).
		Exec(ctx)
	if err != nil {
		f.log.Error(err)
		return pkg.InternalErr
	}

	content, err := pkg.GetArticleContent(url)
	if err != nil {
		f.log.Error(err)
	} else {
		a := &articleObj{
			Title:       title,
			Content:     content,
			Desc:        desc,
			Tags:        tags,
			Category:    category,
			CreatedTime: pkg.NowTimeStr(),
		}
		aStr, err := pkg.AnyToJsonStr(a)
		if err != nil {
			f.log.Error(err)
		} else {
			f.data.ArticleCache.RPush(ctx, pkg.ArticleListKey, aStr)
			f.data.ArticleCache.HSet(ctx, pkg.ArticleMapKey, title, aStr)
			log.Println("article str", aStr, " to cache succeed")
		}
	}
	return nil
}

func (f *blugRepo) GetArticleListInData(ctx context.Context, offset int) (*v1.GetArticleListResp, error) {

	result, err := f.data.ArticleCache.LRange(ctx, pkg.ArticleListKey, 0, -1).Result()
	if err != nil {
		f.log.Error(err)
	} else {
		articles := make([]articleObj, 0)
		err = pkg.JsonStrSliceToAny(result, &articles)
		if err != nil {
			f.log.Error(err)
		} else if len(articles) != 0 {
			f.log.Info("get article list shoot cache succeed!")
			return articleObjToV1ArticleListResp(articles), nil
		} else {
			f.log.Info("this req will get article list from db")
		}
	}

	articles, err := f.data.DB.Article.Query().Where(article.IsShowEQ(true)).All(ctx)
	if err != nil {
		f.log.Error(err)
		return &v1.GetArticleListResp{}, pkg.InternalErr
	}
	resp := make([]*v1.Article, 0)
	for _, aObj := range articles {
		resp = append(resp, &v1.Article{
			Title:      aObj.Title,
			Desc:       aObj.Desc,
			Tags:       aObj.Tags,
			Category:   aObj.Category,
			CreateTime: aObj.CreateTime,
		})
	}
	return &v1.GetArticleListResp{
		Articles: resp,
	}, nil
}

func (f *blugRepo) GetArticleInData(ctx context.Context, title string) (*v1.GetArticleByTitleResp, error) {

	a, err := f.data.ArticleCache.HGet(ctx, pkg.ArticleMapKey, title).Result()
	if err != nil {
		f.log.Error(err)
	} else {
		f.log.Info("get article content shoot cache succeed!")
		obj := articleObj{}
		err = pkg.JsonStrToAny(a, &obj)
		return &v1.GetArticleByTitleResp{
			Article: &v1.Article{
				Title:      obj.Title,
				Desc:       obj.Desc,
				Tags:       obj.Tags,
				Category:   obj.Category,
				CreateTime: obj.CreatedTime,
				Content:    obj.Content,
			},
		}, nil
	}

	aObj, err := f.data.DB.Article.Query().Where(article.TitleEQ(title)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			notFoundObj, err := f.data.DB.Article.Query().Where(article.TitleEQ("404")).First(ctx)
			if err != nil {
				f.log.Error(err)
				return &v1.GetArticleByTitleResp{}, pkg.InternalErr
			}
			notFoundContent, err := pkg.GetArticleContent(notFoundObj.URL)
			if err != nil {
				f.log.Error(err)
				return &v1.GetArticleByTitleResp{}, pkg.InternalErr
			}
			return &v1.GetArticleByTitleResp{
				Article: &v1.Article{
					Title:      notFoundObj.Title,
					Desc:       notFoundObj.Desc,
					Tags:       notFoundObj.Tags,
					Category:   notFoundObj.Category,
					CreateTime: notFoundObj.CreateTime,
					Content:    notFoundContent,
				},
			}, pkg.ArticleNotFoundErr
		} else {
			f.log.Error(err)
			internalErrObj, err := f.data.DB.Article.Query().Where(article.TitleEQ("500")).First(ctx)
			if err != nil {
				f.log.Error(err)
				return &v1.GetArticleByTitleResp{}, pkg.InternalErr
			}
			internalErrContent, err := pkg.GetArticleContent(internalErrObj.URL)
			if err != nil {
				f.log.Error(err)
				return &v1.GetArticleByTitleResp{}, pkg.InternalErr
			}
			return &v1.GetArticleByTitleResp{
				Article: &v1.Article{
					Title:      internalErrObj.Title,
					Desc:       internalErrObj.Desc,
					Tags:       internalErrObj.Tags,
					Category:   internalErrObj.Category,
					CreateTime: internalErrObj.CreateTime,
					Content:    internalErrContent,
				},
			}, pkg.InternalErr
		}
	}
	content, err := pkg.GetArticleContent(aObj.URL)
	if err != nil {
		f.log.Error(err)
		return &v1.GetArticleByTitleResp{}, pkg.InternalErr
	}
	return &v1.GetArticleByTitleResp{
		Article: &v1.Article{
			Title:      aObj.Title,
			Desc:       aObj.Desc,
			Tags:       aObj.Tags,
			Content:    content,
			Category:   aObj.Category,
			CreateTime: aObj.CreateTime,
		},
	}, nil
}

func articleObjToV1Link(article articleObj) *v1.Article {
	return &v1.Article{
		Title:      article.Title,
		Desc:       article.Desc,
		Tags:       article.Tags,
		Category:   article.Category,
		CreateTime: article.CreatedTime,
		Content:    article.Content,
	}
}
func articleObjToV1ArticleListResp(articles []articleObj) *v1.GetArticleListResp {
	resp := make([]*v1.Article, 0)
	for _, a := range articles {
		resp = append(resp, articleObjToV1Link(a))
	}
	return &v1.GetArticleListResp{Articles: resp}
}
