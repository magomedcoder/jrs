// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"json-rpc-skeleton/internal/provider"
	"json-rpc-skeleton/internal/transport/rpc/handler"
	"json-rpc-skeleton/pkg/rpc"
)

// Injectors from wire.go:

func Initialize() *Provider {
	example := handler.NewExampleHandler()
	handlerHandler := &handler.Handler{
		Example: example,
	}
	server := provider.NewRpcServer(handlerHandler)
	mainProvider := &Provider{
		Server: server,
	}
	return mainProvider
}

// wire.go:

type Provider struct {
	Server *rpc.Server
}

var newSet = wire.NewSet(wire.Struct(new(Provider), "*"), wire.Struct(new(handler.Handler), "*"), provider.NewRpcServer, handler.NewExampleHandler)
