package process

import (
	"context"
	"errors"
	"sync"

	"go.uber.org/zap"
)

type ProcessManager struct {
	runningLock          sync.RWMutex
	running              []Process
	processes            []Process
	logger               *zap.Logger
	globalShutdownerFunc ShutdownFunc
}

type Process struct {
	Name   string
	Object ProcessStarter
}

func NewProcessManager(
	logger *zap.Logger,
	processes ...Process,
) *ProcessManager {
	return &ProcessManager{
		runningLock: sync.RWMutex{},
		processes:   processes,
		logger:      logger,
	}
}

func (p *ProcessManager) WithGlobalShutdowner(f ShutdownFunc) *ProcessManager {
	p.globalShutdownerFunc = f

	return p
}

func (p *ProcessManager) Shutdown() error {
	p.logger.Info("Shutdown process manager")

	p.runningLock.RLock()
	defer p.runningLock.RUnlock()

	var resErr error
	for _, process := range p.running {
		if pr, ok := process.Object.(ProcessShutdowner); ok {
			p.logger.Info("Shutting down process start", zap.String("name", process.Name))

			if err := pr.Shutdown(); err != nil {
				p.logger.Error("failed to shutdown Object", zap.Error(err), zap.String("name", process.Name))

				resErr = errors.Join(resErr, err)
			}

			p.logger.Info("Shutting down process finish", zap.String("name", process.Name))
		}
	}

	if p.globalShutdownerFunc != nil {
		if err := p.globalShutdownerFunc(); err != nil {
			resErr = errors.Join(resErr, err)
		}
	}

	return resErr
}

func (p *ProcessManager) Start(ctx context.Context) error {
	p.logger.Info("Process manager starting...")
	p.runningLock.Lock()
	defer p.Shutdown()
	defer p.runningLock.Unlock()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, process := range p.processes {
		p.logger.Info("Starting process", zap.String("name", process.Name))

		p.running = append(p.running, process)

		go func() {
			if err := process.Object.Start(ctx); err != nil {
				p.logger.Error("Process error", zap.Error(err))
			}
			p.logger.Info("Process done", zap.String("name", process.Name))

			cancel()
		}()
	}

	<-ctx.Done()

	return nil
}
