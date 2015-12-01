package interpoller

import (
	"fmt"
	"net"
	"testing"
)

func TestHostIfaces(t *testing.T) {
	//Hopefullying the tests are running on a VM/host with network interfaces
	ifaces, err := HostInterfaces()
	if err != nil {
		t.Errorf("Error gathering interfaces: %#v", err)
	}

	if len(*ifaces) < 1 {
		t.Errorf("No iterfaces returned!")
	}
	fmt.Printf("%#v\n", ifaces)
}

func TestAddrs(t *testing.T) {
	ifaces, _ := HostInterfaces()
	face := *ifaces
	addrs, _ := AddressesOnInterface(&face[0])
	fmt.Printf("%#v\n", addrs)
}

type NetMock struct{}

func (nm *NetMock) Interfaces() ([]net.Interface, error) {
	iface := &[]net.Interface{net.Interface{Index: 1, MTU: 65536, Name: "lo", HardwareAddr: net.HardwareAddr(nil), Flags: 0x5}, net.Interface{Index: 2, MTU: 1500, Name: "eth0", HardwareAddr: net.HardwareAddr{0x3c, 0x97, 0xe, 0xc1, 0xee, 0x15}, Flags: 0x13}}
	return *iface, nil
}

func (nm *NetMock) InterfaceAddrs() ([]net.Addr, error) {
	netAddrs := []net.Addr{AddrMock{network: "eth0s", address: "127.0.1.117"}}
	return netAddrs, nil
}

type AddrMock struct {
	network string
	address string
}

func (a AddrMock) Network() string {
	return a.network
}
func (a AddrMock) String() string {
	return a.address
}

func TestNetMock(t *testing.T) {
	nm := &NetMock{}
	addrs, _ := nm.InterfaceAddrs()

	if len(addrs) < 1 {
		t.Errorf("No addr entries returned from Mock")
	}
	addr := addrs[0]
	if addr.Network() != "eth0s" {
		t.Errorf("Network() call to Addr mock failed to return proper value!")
	}
}
