package process

import "context"

type ProcessStarter interface {
	Start(ctx context.Context) error
}

type ProcessShutdowner interface {
	Shutdown() error
}
