package main

import (
	"net"
)

func getIP(v uint8) string {
	addresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, a := range addresses {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			switch v {
			case 4:
				if ipNet.IP.To4() != nil {
					return ipNet.IP.String()
				}
			case 6:
				if ipNet.IP.To4() == nil {
					return ipNet.IP.String()
				}
			}
		}
	}

	return ""
}
