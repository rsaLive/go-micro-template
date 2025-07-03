package service

import (
	"go-micro-template/internal/biz"
)

func NewCastService(cast *biz.CastUseCase) *CastService {
	return &CastService{castCase: cast}
}
