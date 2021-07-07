package pb

import (
	"github.com/golang/protobuf/proto"
	"sync"
)

// ProposePool 对象池
var ProposePool = sync.Pool{
	New: func() interface{} {
		return &Propose{}
	},
}

func GetObjPropose() *Propose {
	return ProposePool.Get().(*Propose)
}

func PutObjPropose(x *Propose) {
	ProposePool.Put(x)
}

func (x *Propose) Encode() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *Propose) Decode(b []byte) error {

	return proto.Unmarshal(b, x)
}
