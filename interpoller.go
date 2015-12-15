package interpoller

import (
	"net"

	log "github.com/Sirupsen/logrus"
)

func GetAddrs() (*[]net.Addr, error) {
	ifaces, err := HostInterfaces()
	var addrs *[]net.Addr
	if err != nil {
		return nil, err
	}
	for _, i := range *ifaces {
		addrs, err := AddressesOnInterface(&i)
		if err != nil {
			return nil, err
		}
		for _, a := range *addrs {
			log.WithFields(log.Fields{"IFace": i.Name, "NetworkType": a.Network(), "Address": a.String()}).Infof("Interface Address")
		}
	}
	return addrs, nil
}

type ActiveInterfaces struct {
	count     int
	Addresses []net.Addr
}

func NewActiveInterfaces(addrs *[]net.Addr) *ActiveInterfaces {
	return &ActiveInterfaces{count: len(*addrs), Addresses: *addrs}
}
