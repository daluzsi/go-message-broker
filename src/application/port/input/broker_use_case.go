package input

import "context"

type BrokerUseCase interface {
	Listen(ctx context.Context) error
}
