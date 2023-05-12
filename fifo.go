package fifo

import (
	"github.com/dop251/goja"
	"github.com/enriquebris/goconcurrentqueue"
	"go.k6.io/k6/js/modules"
)

// TODO : manage named queues

type (
	// FIFO is the global module instance that will create Client
	// instances for each VU.
	FIFO struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		vu modules.VU
		*Client
	}
)

// Ensure the interfaces are implemented correctly
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &FIFO{}
)

type Client struct {
	vu    modules.VU
	queue *goconcurrentqueue.FIFO
}

var check = false
var client *Client

func init() {
	modules.Register("k6/x/fifo", new(FIFO))
}

// New returns a pointer to a new KV instance
func New() *FIFO {
	return &FIFO{}
}

// NewModuleInstance implements the modules.Module interface and returns
// a new instance for each VU.
func (*FIFO) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{vu: vu, Client: &Client{vu: vu}}
}

// Exports implements the modules.Instance interface and returns
// the exports of the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{Named: map[string]interface{}{
		"Client": mi.NewClient,
	}}
}

// NewClient is the JS constructor for the Client
func (mi *ModuleInstance) NewClient(call goja.ConstructorCall) *goja.Object {
	rt := mi.vu.Runtime()

	if check != true {
		var q *goconcurrentqueue.FIFO
		q = goconcurrentqueue.NewFIFO()
		client = &Client{vu: mi.vu, queue: q}
		check = true
	}

	return rt.ToValue(client).ToObject(rt)
}

// Push the given value.
func (c *Client) Push(value string) error {
	err := c.queue.Enqueue(value)
	return err
}

// Pop returns the oldest value.
func (c *Client) Pop() (string, error) {
	value, err := c.queue.Dequeue()
	if value == nil {
		return "", err
	}
	return value.(string), err
}
