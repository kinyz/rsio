package raft

type Role uint8

const (
	Role_Creator  Role = 0
	Role_Follower Role = 1
	Role_Witness  Role = 2
	Role_Observer Role = 3
)
