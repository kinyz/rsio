package raft

import (
	"errors"
	"fmt"
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/statemachine"
	"path/filepath"
	"strconv"
)

type Raft interface {
	GetRaftId() string
	GetNodeId() uint64
	GetNodeHost() *dragonboat.NodeHost
	StartCluster(clusterId uint64, role Role, create statemachine.CreateOnDiskStateMachineFunc) error
}

func New(addr, port string, nodeId uint64) Raft {
	datadir := filepath.Join(
		"sync-data",
		"node-data",
		fmt.Sprintf("node%d", nodeId))

	nhc := config.NodeHostConfig{
		WALDir:         datadir,
		NodeHostDir:    datadir,
		RTTMillisecond: 200,
		RaftAddress:    addr + ":" + port,
		//	AddressByNodeHostID:true,
		ListenAddress: "0.0.0.0:" + port,
	}
	nh, err := dragonboat.NewNodeHost(nhc)
	if err != nil {
		panic(err)
	}
	intNum, err := strconv.Atoi(nh.ID())

	if err != nil {
		panic(err)
	}

	return &raft{nh: nh, nodeId: uint64(intNum)}
}

type raft struct {
	nh     *dragonboat.NodeHost
	nodeId uint64
}

func (r *raft) GetRaftId() string {
	return r.nh.ID()
}
func (r *raft) GetNodeId() uint64 {
	return r.nodeId
}
func (r *raft) GetNodeHost() *dragonboat.NodeHost {

	return r.nh
}

func (r *raft) StartCluster(clusterId uint64, role Role, create statemachine.CreateOnDiskStateMachineFunc) error {
	join := false
	initialMembers := make(map[uint64]string)

	rc := config.Config{
		NodeID:             r.nodeId,
		ClusterID:          clusterId,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    10,
		CompactionOverhead: 5,
	}
	switch role {
	case RoleCreator:
		initialMembers[1] = r.nh.RaftAddress()
		break
	case RoleFollower:
		join = true
		break
	case RoleWitness:
		rc.SnapshotEntries = 0
		rc.IsWitness = true
		join = true
	case RoleObserver:
		rc.IsObserver = true
		join = true
	default:
		return errors.New("role is error")
	}
	return r.nh.StartOnDiskCluster(initialMembers, join, create, rc)
}
