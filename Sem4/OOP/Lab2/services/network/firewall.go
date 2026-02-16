package network

import "net"

type Firewall struct {
	rules         map[string]bool
	defaultPolicy string
	interfaces    map[string]string

	blockedIps    []net.IP
	activeSession []net.Conn

	natRules []string
}
