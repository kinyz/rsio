package kv

import "rsio/raft"

type Kv interface {
	SetMemoryData(key string, value []byte) error
	GetMemoryData(key string) (value []byte, err error)
	SetDiskData(key string, value []byte) error
	GetDiskData(key string) (value []byte, err error)
}

func New() Kv {

	return &kv{}
}

type kv struct {
	raft raft.Raft
}

func (k *kv) SetMemoryData(key string, value []byte) error {

	return nil

}

func (k *kv) GetMemoryData(key string) (value []byte, err error) {
	value = nil
	err = nil
	return
}

func (k *kv) SetDiskData(key string, value []byte) error {

	return nil
}

func (k *kv) GetDiskData(key string) (value []byte, err error) {

	value = nil
	err = nil
	return
}
