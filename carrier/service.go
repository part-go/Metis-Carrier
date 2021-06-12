package carrier

import (
	"context"
	"github.com/RosettaFlow/Carrier-Go/core"
	"github.com/RosettaFlow/Carrier-Go/core/resource"
	"github.com/RosettaFlow/Carrier-Go/params"
	"sync"
)

type Service struct {
	isRunning       bool
	processingLock  sync.RWMutex
	config          *params.CarrierConfig
	proxy           *core.DataCenter
	datachain       *core.DataChain
	ctx             context.Context
	cancel          context.CancelFunc
	mempool         *core.Mempool

	resourceManager *resource.Manager
	runError        error
}

func NewService(ctx context.Context, config *params.CarrierConfig, dataCenterConfig *params.DataCenterConfig) (*Service, error) {
	ctx, cancel := context.WithCancel(ctx)
	_ = cancel // govet fix for lost cancel. Cancel is handled in service.Stop()
	proxy, err := core.NewDataCenter(dataCenterConfig)
	if err != nil {
		cancel()
		return nil, err
	}
	s := &Service{
		ctx:             ctx,
		cancel:          cancel,
		config:          config,
		proxy:           proxy,
		mempool:         core.NewMempool(nil), // todo need  set mempool cfg
		resourceManager: resource.NewResourceManager(),
	}
	// todo: some logic could be added...

	// todo: set datachain....
	return s, nil
}

func (s *Service) Start() {

}

func (s *Service) Stop() error {
	if s.cancel != nil {
		defer s.cancel()
	}
	// todo: could add some logic for here
	return nil
}

// Status is service health checks. Return nil or error.
func (s *Service) Status() error {
	// Service don't start
	if !s.isRunning {
		return nil
	}
	// get error from run function
	if s.runError != nil {
		return s.runError
	}
	return nil
}

