package protocol

import (
	"net"
	"testing"
)

func TestICMP(t *testing.T) {
	icmp := ICMP{
		Type:       8,
		Code:       0,
		Checksum:   0,
		Identifier: 13,
		Sequence:   37,
	}

	icmp.Checksum = checksum(icmp)

	raddr := &net.IPAddr{IP: net.ParseIP("114.114.114.114")}
	conn, err := net.DialIP("ip:icmp", nil, raddr)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = conn.Write(bytes(icmp))
	if err != nil {
		t.Error(err)
		return
	}
}
