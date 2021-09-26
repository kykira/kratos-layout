//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/kykira/kratos-layout/internal/biz"
	"github.com/kykira/kratos-layout/internal/conf"
	"github.com/kykira/kratos-layout/internal/data"
	"github.com/kykira/kratos-layout/internal/server"
	"github.com/kykira/kratos-layout/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
