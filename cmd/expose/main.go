package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ropes/interpoller"
)

func main() {
	ifaces, err := interpoller.HostInterfaces()
	if err != nil {
		log.Errorf(err.Error())
	}
	for _, i := range *ifaces {
		addrs, err := interpoller.AddressesOnInterface(&i)
		if err != nil {
			log.Errorf(err.Error())
		}
		for _, a := range *addrs {
			log.WithFields(log.Fields{"IFace": i.Name, "NetworkType": a.Network(), "Address": a.String()}).Infof("Interface Address")
		}
	}
}
