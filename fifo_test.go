package fifo

import (
	"context"
	"testing"

	"github.com/grafana/sobek"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/lib"
)

type mockVU struct {
	name string
}

// instantiate the interface
func (mockVU) Context() context.Context {
	return context.TODO()
}
func (mockVU) InitEnv() *common.InitEnvironment {
	return nil
}
func (mockVU) State() *lib.State {
	return nil
}
func (mockVU) Runtime() *sobek.Runtime {
	return &sobek.Runtime{}
}
func (mockVU) RegisterCallback() (enqueueCallback func(func() error)) {
	return enqueueCallback
}

func TestFIFO(t *testing.T) {
	Arguments := []sobek.Value{}

	client1 := newClient(Arguments, mockVU{})
	client1.Push("first value")
	client1.Push("2nd value")
	client2 := newClient(Arguments, mockVU{})
	out1, _ := client1.Pop()
	out2, _ := client2.Pop()
	if out1 != "first value" || out2 != "2nd value" {
		t.Errorf("Single fifo : out1=%s out2=%s", out1, out2)
	}
}

func TestNamedFIFO(t *testing.T) {
	vm := sobek.New()

	Arguments := []sobek.Value{
		vm.ToValue("liste A"),
	}
	client1 := newClient(Arguments, mockVU{})
	client1.Push("first value A")
	client1.Push("2nd value A")

	Arguments = []sobek.Value{
		vm.ToValue("liste B"),
	}
	client2 := newClient(Arguments, mockVU{})
	client2.Push("first value B")
	client2.Push("2nd value B")

	client3 := newClient(Arguments, mockVU{})

	out1, _ := client1.Pop()
	out2, _ := client2.Pop()
	out3, _ := client3.Pop()

	if out1 != "first value A" || out2 != "first value B" || out3 != "2nd value B" {
		t.Errorf("Named fifos : out1=%s out2=%s out3=%s", out1, out2, out3)
	}

}
