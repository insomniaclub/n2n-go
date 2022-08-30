package protocol

import (
	"reflect"
	"unsafe"
)

type Protocol interface {
	ICMP
}

func checksum[P Protocol](data P) uint16 {
	len := (unsafe.Sizeof(data) + 1) / 2
	bytesHeader := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&data)),
		Len:  int(len),
		Cap:  int(len),
	}
	bytePairs := *(*[]uint16)(unsafe.Pointer(bytesHeader))

	var chksum uint32
	for _, bytePair := range bytePairs {
		chksum += uint32(bytePair)
	}
	chksum = chksum>>16 + chksum&0xFFFF
	chksum += chksum >> 16

	return ^uint16(chksum)
}

func bytes[P Protocol](data P) []byte {
	len := unsafe.Sizeof(data)
	bytesHeader := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&data)),
		Len:  int(len),
		Cap:  int(len),
	}
	return *(*[]byte)(unsafe.Pointer(bytesHeader))
}
