//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Alphasxd/Micro4um/app/forum-interactive/internal/biz"
	"github.com/Alphasxd/Micro4um/app/forum-interactive/internal/conf"
	"github.com/Alphasxd/Micro4um/app/forum-interactive/internal/data"
	"github.com/Alphasxd/Micro4um/app/forum-interactive/internal/server"
	"github.com/Alphasxd/Micro4um/app/forum-interactive/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
