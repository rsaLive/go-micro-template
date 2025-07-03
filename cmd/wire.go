// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	configData "go-micro-template/config"
	"go-micro-template/internal/biz"
	"go-micro-template/internal/data"
	"go-micro-template/internal/service"
	"go-micro-template/internal/socialconfig"
	"go-micro-template/pkg/app"
	"go-micro-template/pkg/logger"
	"go-micro-template/pkg/tracing"
	"github.com/google/wire"
)

func InitApp() (*app.App, error) {
	wire.Build(configData.ProviderSet, socialconfig.ProviderSet, logger.ProviderSet, biz.ProviderSet, data.ProviderSet, service.ProviderSet, tracing.Provider, NewApp)
	return &app.App{}, nil
}
