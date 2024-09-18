//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Alphasxd/Micro4um/app/forum-user/internal/biz"
	"github.com/Alphasxd/Micro4um/app/forum-user/internal/conf"
	"github.com/Alphasxd/Micro4um/app/forum-user/internal/data"
	"github.com/Alphasxd/Micro4um/app/forum-user/internal/middleware"
	"github.com/Alphasxd/Micro4um/app/forum-user/internal/server"
	"github.com/Alphasxd/Micro4um/app/forum-user/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, middleware.ProviderSet, newApp))
}
