package service

import (
	pb "blug/api/blug/v1"
	"blug/internal/biz"
)

type BlugService struct {
	pb.UnimplementedBlugServer
	uc *biz.BlugUsecase
}

func NewBlugService(uc *biz.BlugUsecase) *BlugService {
	return &BlugService{uc: uc}
}
