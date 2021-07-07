package witness

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"rsio/pb"
	"rsio/raft"
	"time"
)



type RequestHandle interface {
	Oauth(req *pb.RequestJoinCluster) (*pb.ResponseJoinResult, error)
}

/*
	观察者节点，用来通过集群验证
 */
type Witness interface {

}
func New()Witness{
	return &witness{}
}
type witness struct {

	raft raft.Raft

	reqHandle RequestHandle
}

func (w *witness) RequestJoin(ctx context.Context, req *pb.RequestJoinCluster) (*pb.ResponseJoinResult, error) {
	resp, err := w.reqHandle.Oauth(req)
	if err != nil {
		return nil, err
	}
	switch req.GetJoinRole() {
	case pb.Role_Follower:
		err = w.addNode(resp.GetClusterId(), req.GetNodeId(), req.GetRaftAddresses())
		if err != nil {
			return nil, err
		}
		break
	case pb.Role_Witness:
		err = w.addWitnessNode(resp.GetClusterId(), req.GetNodeId(), req.GetRaftAddresses())
		if err != nil {
			return nil, err
		}
		break
	case pb.Role_Observer:
		err = w.addObserverNode(resp.GetClusterId(), req.GetNodeId(), req.GetRaftAddresses())
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, errors.New("role error")
	}

	return resp, nil
}

func (w*witness)Start(){


}

func (w*witness) StartOauthService(addr string, srv RequestHandle) error {
	if w.reqHandle != nil {
		return errors.New("do not turn on repeatedly ！")
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	w.reqHandle = srv
	s := grpc.NewServer()
	pb.RegisterRequestServer(s, w)
	log.Println("oauth service start: ", addr)
	return s.Serve(ln)
}



func (w*witness) addNode(clusterId uint64, nodeId uint64, addr string) error {
	_, err := w.raft.GetNodeHost().RequestAddNode(clusterId, nodeId, addr, 0, 5*time.Second)
	if err != nil {
		return err
	}

	return nil
}
func (w*witness) addWitnessNode(clusterId uint64, nodeId uint64, addr string) error {
	_, err := w.raft.GetNodeHost().RequestAddWitness(clusterId, nodeId, addr, 0, 5*time.Second)
	if err != nil {
		return err
	}

	return nil
}

func (w*witness) addObserverNode(clusterId uint64, nodeId uint64, addr string) error {
	_, err := w.raft.GetNodeHost().RequestAddObserver(clusterId, nodeId, addr, 0, 5*time.Second)
	if err != nil {
		return err
	}
	return nil
}