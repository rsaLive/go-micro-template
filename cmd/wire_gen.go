// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/oa-meeting/pkg/cache"
	"github.com/oa-meeting/pkg/db"
	"github.com/oa-meeting/pkg/logger"
	"github.com/oa-meeting/pkg/snowf"
	"github.com/oa-meeting/pkg/tracing"
)

import (
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps/strategy"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/oa-meeting/internal/handler"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	zapLogger := logger.ZapInit()
	client := cache.NewRedis()
	jaegerProvider := tracing.NewTracing()
	gormDB := db.NewDb()
	node := snowf.NewSf()
	app := NewApp(zapLogger, client, jaegerProvider, gormDB, node)
	return app, nil
}