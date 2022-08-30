package protocol

type ICMP struct {
	Type       uint8
	Code       uint8
	Checksum   uint16
	Identifier uint16
	Sequence   uint16
	Data       []byte
}