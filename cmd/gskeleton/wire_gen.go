// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/cli"
	handler3 "github.com/magomedcoder/gskeleton/internal/cli/handler"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc"
	handler2 "github.com/magomedcoder/gskeleton/internal/delivery/grpc/handler"
	middleware2 "github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v1"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v2"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/router"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	repository2 "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

// Injectors from wire.go:

func NewHttpInjector(conf *config.Config) *http.AppProvider {
	db := provider.NewPostgresDB(conf)
	userRepository := repository.NewUserRepository(db)
	client := provider.NewRedisClient(conf)
	userCacheRepository := repository2.NewUserCacheRepository(client)
	userUseCase := &usecase.UserUseCase{
		PostgresUserRepo:         userRepository,
		RedisUserCacheRepository: userCacheRepository,
	}
	user := v1.NewUserHandler(userUseCase)
	v1V1 := &v1.V1{
		User: user,
	}
	v2User := v2.NewUserHandler(userUseCase)
	v2V2 := &v2.V2{
		User: v2User,
	}
	handlerHandler := &handler.Handler{
		V1: v1V1,
		V2: v2V2,
	}
	authMiddleware := middleware.NewAuthMiddleware()
	middlewareMiddleware := &middleware.Middleware{
		AuthMiddleware: authMiddleware,
	}
	engine := router.NewRouter(handlerHandler, middlewareMiddleware)
	appProvider := &http.AppProvider{
		Conf:   conf,
		Engine: engine,
	}
	return appProvider
}

func NewGrpcInjector(conf *config.Config) *grpc.AppProvider {
	tokenMiddleware := middleware2.NewTokenMiddleware(conf)
	grpcMethodService := middleware2.NewGrpMethodsService()
	unimplementedAuthServiceServer := pb.UnimplementedAuthServiceServer{}
	db := provider.NewPostgresDB(conf)
	userRepository := repository.NewUserRepository(db)
	client := provider.NewRedisClient(conf)
	userCacheRepository := repository2.NewUserCacheRepository(client)
	userUseCase := &usecase.UserUseCase{
		PostgresUserRepo:         userRepository,
		RedisUserCacheRepository: userCacheRepository,
	}
	authHandler := &handler2.AuthHandler{
		UnimplementedAuthServiceServer: unimplementedAuthServiceServer,
		TokenMiddleware:                tokenMiddleware,
		UserUseCase:                    userUseCase,
	}
	unimplementedUserServiceServer := pb.UnimplementedUserServiceServer{}
	userHandler := &handler2.UserHandler{
		UnimplementedUserServiceServer: unimplementedUserServiceServer,
		UserUseCase:                    userUseCase,
		TokenMiddleware:                tokenMiddleware,
	}
	appProvider := &grpc.AppProvider{
		Conf:            conf,
		TokenMiddleware: tokenMiddleware,
		RoutesServices:  grpcMethodService,
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
	}
	return appProvider
}

func NewCliInjector(conf *config.Config) *cli.AppProvider {
	db := provider.NewPostgresDB(conf)
	migrate := &handler3.Migrate{
		Conf: conf,
		Db:   db,
	}
	appProvider := &cli.AppProvider{
		Conf:    conf,
		Migrate: migrate,
	}
	return appProvider
}
