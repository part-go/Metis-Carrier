package scheduler

import (
	"github.com/RosettaFlow/Carrier-Go/consensus"
	"github.com/RosettaFlow/Carrier-Go/types"
)

type SchedulerStarveFIFO struct {
	engine consensus.Consensus
}

func (sche *SchedulerStarveFIFO) SetTaskEngine(engine consensus.Consensus) error {return nil}

func (sche *SchedulerStarveFIFO) OnSchedule() error {
	return nil
}
func (sche *SchedulerStarveFIFO) OnError () error {
	return nil
}
func (sche *SchedulerStarveFIFO) SchedulerName() string {
	return ""
}
func (sche *SchedulerStarveFIFO) PushTasks(tasks types.TaskMsgs) error {

	return nil
}