//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	wire "github.com/google/wire"
	"github.com/spf13/viper"
	"nunu-project/internal/handler"
	"nunu-project/internal/middleware"
	"nunu-project/internal/repository"
	"nunu-project/internal/server"
	"nunu-project/internal/service"
	"nunu-project/pkg/helper/sid"
	"nunu-project/pkg/log"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var SidSet = wire.NewSet(sid.NewSid)

var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,

	handler.NewOrderHandler, // new
	handler.NewBlogHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,

	service.NewOrderService, // new
	service.NewBlogService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
	repository.NewOrderRepository, // new
	repository.NewBlogRepository,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		SidSet,
		JwtSet,
	))
}
