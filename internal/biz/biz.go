package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUseCase, NewCastUseCase)

type UseCase struct {
	Cast *CastUseCase
}

func NewUseCase(cast *CastUseCase) *UseCase {
	return &UseCase{Cast: cast}
}

const (
	WorkLogSourceCreateImage int = iota + 1
	WorkLogSourceCreateVideo
	WorkLogSourceCreateApproval
	WorkLogSourceApprovalPass
	WorkLogSourceApprovalReject
	WorkLogSourceApprovalDelete
	WorkLogSourceArtistPass
	WorkLogSourceArtistReject
	WorkLogSourcePublish
	WorkLogSourcePublishSuccess
	WorkLogSourcePublishFailed
)

type WorkLogCreateIn struct {
	WorkLogSource int
	WorkLogRepo   WorkLogRepo
}
