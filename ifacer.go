package interpoller

import (
	"fmt"
	"net"
)

type IfaceError struct {
	Name   string
	Addrs  []net.Addr
	ErrStr string
}

func (ife IfaceError) Error() string {
	if ife.ErrStr == "" {
		return fmt.Sprintf("Error gathering host interface details: %#v %#v\n", ife.Name, ife.Addrs)
	} else {
		return fmt.Sprintf("Error in iface.go: %s %#v %#v\n", ife.ErrStr, ife.Name, ife.Addrs)
	}
}

type NetworkInfo interface {
	Interfaces() ([]net.Interface, error)
	InterfaceAddrs() ([]net.Addr, error)
}

func HostInterfaces() (*[]net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
		return nil, err
	}
	return &ifaces, nil
}

func AddressesOnInterface(iface *net.Interface) (*[]net.Addr, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}
	return &addrs, nil
}
