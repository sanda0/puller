package test

import (
	"fmt"
	"net"
	"testing"
)

func TestGetIp(t *testing.T) {

	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)

	}
	for _, addr := range address {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Print(`


	USE THIS URL AS YOUR GITLAB WEBHOOK URL: http://` + ipNet.IP.String() + `:8080/puller


				`)
			}
		}
	}

}
