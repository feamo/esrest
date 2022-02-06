package storage

import (
	"fmt"
	"github.com/feamo/esrest/models"
	"sync"
	"time"
)

// Conn represents a single connection to a node in a cluster
type Conn struct {
	sync.RWMutex
	nodeID    string // node ID
	url       string
	failures  int
	dead      bool
	deadSince *time.Time
}

// newConn creates a new connection to the given URL.
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

func (e *Conn) UserInsert(user *models.User) error {
	query := `insert into "users"("email", "password") values($1, $2)`
	_, err := e.db.Exec(query, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
