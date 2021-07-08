package witness

import (
	"github.com/lni/dragonboat/v3/statemachine"
	"io"
)

func newState(clusterID uint64,
	nodeID uint64) statemachine.IStateMachine {
	return &state{clusterID: clusterID, nodeID: nodeID}

}

type state struct {
	clusterID uint64
	nodeID    uint64
}

func (s *state) Update(bytes []byte) (statemachine.Result, error) {
	panic("implement me")
}

func (s *state) Lookup(i interface{}) (interface{}, error) {
	panic("implement me")
}

func (s *state) SaveSnapshot(writer io.Writer, collection statemachine.ISnapshotFileCollection, i <-chan struct{}) error {
	return nil
}

func (s *state) RecoverFromSnapshot(reader io.Reader, files []statemachine.SnapshotFile, i <-chan struct{}) error {
	return nil
}

func (s state) Close() error {
	return nil
}
