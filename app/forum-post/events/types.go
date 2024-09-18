package events

import (
	"context"
	"github.com/Alphasxd/Micro4um/app/forum-post/events/publish"
	"github.com/Alphasxd/Micro4um/app/forum-post/events/sync"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(sync.NewSyncConsumer, publish.NewSaramaSyncProducer, publish.NewPublishPostEventConsumer)

type Consumer interface {
	Start(ctx context.Context) error
}
