package network

import (
	"Lab1/interfaces"
	"fmt"
	"hash"
)

type NetworkManager struct {
	devices       []string
	topology      string
	errorCoounter uint64
	latencyHash   hash.Hash
	activeTunnels []string
	Requests      map[Requesrt]hash.Hash
	sessions      interfaces.Connection
}

type Requesrt struct {
	Sender  string
	Command string
}

func NewNetworkManager(topology string) *NetworkManager {
	return &NetworkManager{
		devices:       make([]string, 0),
		topology:      topology,
		errorCoounter: 0,
		latencyHash:   nil,
		activeTunnels: make([]string, 0),
	}
}

func (n *NetworkManager) AddDevice(device string) {
	n.devices = append(n.devices, device)
}
func (n *NetworkManager) GetAllDevices() []string {
	return n.devices
}

func (n *NetworkManager) GetTopology() string {
	return n.topology
}
func (n *NetworkManager) SendRequest(r Requesrt) error {
	var h hash.Hash
	_, err := h.Write([]byte(r.Command))
	if err != nil {
		return err
	}
	n.proceedRequest(r)
	n.latencyHash = h
	return nil

}

func (n *NetworkManager) proceedRequest(request Requesrt) {
	fmt.Println("Some systemp calcualtiosn")
}
