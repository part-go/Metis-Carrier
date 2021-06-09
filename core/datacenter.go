package core

import (
	"fmt"
	"github.com/RosettaFlow/Carrier-Go/core/types"
	"github.com/RosettaFlow/Carrier-Go/grpclient"
	"github.com/RosettaFlow/Carrier-Go/params"
	"sync"
	"sync/atomic"
)

// DataCenter is mainly responsible for communicating with the data center service
type DataCenter struct {
	config 			*params.DataCenterConfig
	client 			*grpclient.GrpcClient
	mu     			sync.RWMutex 	// global mutex for locking data center operations.
	procmu 			sync.RWMutex 	// data processor lock
	processor 		Processor	 	// block processor interface
	running     	int32         	// running must be called atomically
	procInterrupt 	int32          	// interrupt signaler for block processing
	wg            	sync.WaitGroup 	// chain processing wait group for shutting down
}

// NewDataCenter returns a fully initialised data center using information available in the database.
func NewDataCenter(config *params.DataCenterConfig) (*DataCenter, error) {
	// todo: When to call Close??
	client, err := grpclient.Dial(fmt.Sprintf("%v:%v", config.GrpcUrl, config.Port))
	if err != nil {
		log.WithError(err).Error("dial grpc server failed")
		return nil, err
	}
	dc := &DataCenter{
		config: 		config,
		client: 		client,
	}
	return dc, nil
}

func (dc *DataCenter) getProcInterrupt() bool {
	return atomic.LoadInt32(&dc.procInterrupt) == 1
}

func (dc *DataCenter) SetProcessor(processor Processor) {
	dc.procmu.Lock()
	defer dc.procmu.Unlock()
	// do setting...
	dc.processor = processor
}

func (dc *DataCenter) Insert(metas []*types.Metadata, resources []*types.Resource, identities []*types.Identity, tasks []*types.Task) error {
	return nil
}

func (dc *DataCenter) Update(metas []*types.Metadata, resources []*types.Resource, identities []*types.Identity, tasks []*types.Task) error {
	return nil
}

func (dc *DataCenter) Delete(metas []*types.Metadata, resources []*types.Resource, identities []*types.Identity, tasks []*types.Task) error {
	return nil
}

func (dc *DataCenter) MetadataList(nodeId string) (types.MetadataArray, error) {
	return nil, nil
}

func (dc *DataCenter) ResourceList(nodeId string) (types.ResourceArray, error) {
	return nil, nil
}

func (dc *DataCenter) IdentityList(nodeId string) (types.IdentityArray, error) {
	return nil, nil
}

func (dc *DataCenter) TaskDataList(nodeId string) (types.TaskDataArray, error) {
	return nil, nil
}




