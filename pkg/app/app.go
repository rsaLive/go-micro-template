package app

import (
	configData "go-micro-template/config"
	"go-micro-template/internal/biz"
	"go-micro-template/internal/data"
	"go-micro-template/internal/service"
	"go-micro-template/internal/socialconfig"
	"go-micro-template/pkg/tracing"
	"go.uber.org/zap"
)

var ModuleClients *App

type App struct {
	AppConfig    *configData.AppConfig
	SocialConfig *socialconfig.Provider
	Lg           *zap.Logger
	Biz          *biz.UseCase
	Data         *data.Data
	Service      *service.CastService
	JaegerTracer *tracing.JaegerProvider
}
