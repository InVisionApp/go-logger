package testlog

import "sync"

/*************
 Safe Counter
*************/

type counter struct {
	num int
	mu  *sync.Mutex
}

//New counter stating at 0
func newCounter() *counter {
	return &counter{num: 0, mu: &sync.Mutex{}}
}

func (c *counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += 1
}

func (c *counter) dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num -= 1
}

func (c *counter) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num = 0
}

func (c *counter) val() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.num
}
