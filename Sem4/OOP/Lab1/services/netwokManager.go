package services

import "hash"

type NetwokManager struct {
	devices       []string
	topology      string
	errorCoounter uint64
	latencyHash   hash.Hash
	activeTunnels []string
}
