// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"prometheus-manager/app/demo/internal/biz"
	"prometheus-manager/app/demo/internal/data"
	"prometheus-manager/app/demo/internal/server"
	"prometheus-manager/app/demo/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(string2 *string) (*kratos.App, func(), error) {
	bootstrap, err := loadConfig(string2)
	if err != nil {
		return nil, nil, err
	}
	log := bootstrap.Log
	logger := newLogger(log)
	confServer := bootstrap.Server
	confData := bootstrap.Data
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	pingRepo := data.NewPingRepo(dataData, logger)
	pingUseCase := biz.NewPingUseCase(pingRepo, logger)
	pingService := service.NewPingService(pingUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, pingService, logger)
	httpServer := server.NewHTTPServer(confServer, pingService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
