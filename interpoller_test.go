package interpoller

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

type MockAddr struct {
	Addr string
	Net  string
}

func (ma MockAddr) Network() string {
	return ma.Net
}
func (ma MockAddr) String() string {
	return fmt.Sprintf("%s %s", ma.Addr, ma.Net)
}

func TestInterfaceAddresses(t *testing.T) {
	addr1 := &MockAddr{Addr: "a.b.c.d", Net: "127.0.0.1/32"}
	addr2 := &MockAddr{Addr: "a.b.c.d", Net: "127.0.0.1/32"}
	addr3 := &MockAddr{Addr: "c.b.a.d", Net: "169.0.0.0/8"}

	if reflect.DeepEqual(addr1, addr2) != true {
		t.Errorf("DeepEqual failed to assert equivalently")
	}

	addrs := &[]net.Addr{*addr1, *addr2}
	ai1 := NewActiveInterfaces(addrs)
	ai2 := NewActiveInterfaces(addrs)

	addrs2 := &[]net.Addr{*addr1, *addr3}
	ai3 := NewActiveInterfaces(addrs2)

	if reflect.DeepEqual(ai1, ai2) != true {
		t.Errorf("ActiveInterfaces failed equivalence test:\n%#v\n%#v\n", ai1, ai2)
	}

	if reflect.DeepEqual(ai1, ai3) == true {
		t.Errorf("Differeing ActiveInterfaces tested as equeivalent:\n%#v\n%#v\n", ai1, ai3)
	}
}
