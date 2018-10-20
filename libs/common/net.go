package common

import (
	"net"
	"strings"
)

func Connect(prototoAddr string) (net.Conn, error) {
	proto, address := ProtocolAndAddress(prototoAddr)
	conn, err := net.Dial(proto, address)
	return conn, err
}
func ProtocolAndAddress(listenAddr string) (string, string) {
	protocol, address := "tcp", listenAddr
	parts := strings.SplitN(address, "://", 2)
	if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	}
	return protocol, address
}
