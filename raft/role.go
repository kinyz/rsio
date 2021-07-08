package raft

type Role uint8

const (
	RoleCreator  Role = 0
	RoleFollower Role = 1
	RoleWitness  Role = 2
	RoleObserver Role = 3
)
