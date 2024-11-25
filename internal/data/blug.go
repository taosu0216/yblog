package data

import (
	"blug/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type blugRepo struct {
	data *Data
	log  *log.Helper
}

func NewBlugRepo(data *Data, logger log.Logger) biz.BlugRepo {
	return &blugRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
