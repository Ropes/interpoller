package interpoller

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
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

func HostInterfaces() (*[]net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Errorf("localAddresses: %+v\n", err.Error())
		return nil, &IfaceError{ErrStr: err.Error(), Name: "hostInterfaces()"}
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

func localInterfaces() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Errorf("localAddresses: %+v\n", err.Error())
		return
	}
	for _, i := range ifaces {
		_, err := i.Addrs()
		if err != nil {
			log.Errorf("localAddresses: %+v\n", err.Error())
			continue
		}
		log.Debugf("IFace: %#v %#v", i.Name, i.Flags)
	}
}
