//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	pb "github.com/Alphasxd/Micro4um/api/user/v1"
	"github.com/Alphasxd/Micro4um/app/forum-check/internal/biz"
	"github.com/Alphasxd/Micro4um/app/forum-check/internal/conf"
	"github.com/Alphasxd/Micro4um/app/forum-check/internal/data"
	"github.com/Alphasxd/Micro4um/app/forum-check/internal/server"
	"github.com/Alphasxd/Micro4um/app/forum-check/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Service, pb.UserClient, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
