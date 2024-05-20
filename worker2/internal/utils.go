package internal

import (
	"net"

	"github.com/sirupsen/logrus"
)

func GetIPAddr() net.IP {
	nConn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logrus.Errorf("cannot get ip address : %v", err)
	}
	lA := nConn.LocalAddr().(*net.UDPAddr)
	return lA.IP
}
