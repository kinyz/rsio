package broker

import (
	"github.com/lni/dragonboat/v3/config"
	"rsio/global"
	"rsio/raft"
)

type Broker interface {

}
type broker struct {
	raft raft.Raft

}


func (b*broker)Start(){

	nh:=b.raft.GetNodeHost()

	rc := config.Config{
		NodeID:             b.raft.GetNodeId(),
		ClusterID:          global.KvClusterId,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    10,
		CompactionOverhead: 5,
	}

	nh.StartCluster(make(map[uint64]string),true,newState,rc)
}