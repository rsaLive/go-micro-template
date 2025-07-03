package service

import (
	"github.com/google/wire"
	"go-micro-template/internal/biz"
)

var ProviderSet = wire.NewSet(NewCastService)

type CastService struct {
	//cast.UnimplementedCastServer

	castCase *biz.CastUseCase
	//youtubeCase *biz.YouTubeUseCase
}
