//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"lehu-data-center/app/agility_data/service/internal/biz"
	"lehu-data-center/app/agility_data/service/internal/conf"
	"lehu-data-center/app/agility_data/service/internal/data"
	"lehu-data-center/app/agility_data/service/internal/server"
	"lehu-data-center/app/agility_data/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
