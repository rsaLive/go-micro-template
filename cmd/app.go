package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps/strategy"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	configData "go-micro-template/config"
	"go-micro-template/internal/biz"
	"go-micro-template/internal/data"
	_ "go-micro-template/internal/handler"
	"go-micro-template/internal/service"
	"go-micro-template/internal/socialconfig"
	"go-micro-template/pkg/app"
	common "go-micro-template/pkg/init"
	"go-micro-template/pkg/tracing"
	//"golang.org/x/oauth2"
	"go.uber.org/zap"
)

//
func NewApp(confData *configData.AppConfig, socialConfig *socialconfig.Provider, lg *zap.Logger, bizUseCase *biz.UseCase, data *data.Data, service *service.CastService, jaegerTracer *tracing.JaegerProvider) *app.App {
	return &app.App{
		AppConfig:    confData,
		SocialConfig: socialConfig,
		Lg:           lg,
		Biz:          bizUseCase,
		Data:         data,
		Service:      service,
		JaegerTracer: jaegerTracer,
	}
}

func main() {
	var err error
	app.ModuleClients, err = InitApp()
	if err != nil {
		panic(err)
	}
	//注册服务
	config.SetProviderService(app.ModuleClients.Service)
	common.Init()
	if err = config.Load(); err != nil {
		panic(err)
	}
	// 启动审批检查器
	//approvalChecker := task.NewApprovalChecker(app.ModuleClients.Biz, app.ModuleClients.Data)
	//go approvalChecker.Start(context.Background())
	select {}
}
