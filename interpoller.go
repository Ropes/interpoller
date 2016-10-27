package interpoller

import "net"

func GetAddrs() (*[]net.Addr, error) {
	ifaces, err := HostInterfaces()
	allAddrs := make([]net.Addr, 0, 0)
	if err != nil {
		return nil, err
	}
	for _, i := range *ifaces {
		addrs, err := AddressesOnInterface(&i)
		if err != nil {
			return nil, err
		}
		for _, a := range *addrs {
			allAddrs = append(allAddrs, a)
		}
	}
	return &allAddrs, nil
}

type ActiveInterfaces struct {
	count     int
	Addresses []net.Addr
}

func NewActiveInterfaces(addrs *[]net.Addr) *ActiveInterfaces {
	return &ActiveInterfaces{count: len(*addrs), Addresses: *addrs}
}
