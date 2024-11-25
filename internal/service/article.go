package service

import (
	v1 "blug/api/blug/v1"
	"blug/internal/pkg"
	"blug/internal/pkg/auth"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func (s *BlugService) Upload(ctx transporthttp.Context) error {
	if header, ok := transport.FromServerContext(ctx); ok {
		jwtToken := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)[1]
		token, err := auth.CheckJWT(jwtToken)
		if err != nil {
			s.uc.Log.Errorf("token解析失败：%v", err)
			return pkg.PermissionDeniedErr
		}
		username := token["username"].(string)
		if username != "root" {
			s.uc.Log.Info("非root")
			return pkg.PermissionDeniedErr
		}
	} else {
		return pkg.PermissionDeniedErr
	}
	file, header, err := ctx.Request().FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// 修改文件名并创建保存文件
	file2, err := os.Create(pkg.GetPostsLocation() + header.Filename)
	if err != nil {
		return err
	}
	defer file2.Close()

	_, err = io.Copy(file2, file)
	if err != nil {
		return err
	}

	filepath := pkg.GetArticleLocation(pkg.GetPostsLocation() + header.Filename)

	content, err := ioutil.ReadFile(pkg.GetPostsLocation() + header.Filename)
	if err != nil {
		log.Errorf("读取文件失败：%v", err)
		return pkg.InternalErr
	}

	err = s.uc.UploadArticle(filepath, string(content), ctx)
	return err
}

func (s *BlugService) GetArticleList(ctx context.Context, req *v1.GetArticleListReq) (*v1.GetArticleListResp, error) {
	return s.uc.GetArticleList(ctx, 0)
}

func (s *BlugService) GetArticleByTitle(ctx context.Context, req *v1.GetArticleByTitleReq) (*v1.GetArticleByTitleResp, error) {
	return s.uc.GetArticle(ctx, req.Title)
}
