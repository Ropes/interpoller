package interpoller

import "testing"

func TestHostIfaces(t *testing.T) {
	//Hopefullying the tests are running on a VM/host with network interfaces
	ifaces, err := HostInterfaces()
	if err != nil {
		t.Errorf("Error gathering interfaces: %#v", err)
	}

	if len(*ifaces) < 1 {
		t.Errorf("No iterfaces returned!")
	}
	t.Logf("%#v\n", ifaces)
}
