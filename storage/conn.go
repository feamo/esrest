package storage

import (
	"fmt"
	"sync"
	"time"
)

// Conn  представляет собой одно соединение с узлом в кластере.
type Conn struct {
	sync.RWMutex
	nodeID    string // node ID
	url       string
	failures  int
	dead      bool
	deadSince *time.Time
}

// newConn создает новое соединение с указанным URL.
func newConn(nodeID, url string) *Conn {
	c := &Conn{
		nodeID: nodeID,
		url:    url,
	}
	return c
}

// String returns a representation of the connection status.
func (c *Conn) String() string {
	c.RLock()
	defer c.RUnlock()
	return fmt.Sprintf("%s [dead=%v,failures=%d,deadSince=%v]", c.url, c.dead, c.failures, c.deadSince)
}
