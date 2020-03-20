package pubsub

import (
	"encoding/binary"
)

func uint16tobyte(input uint16) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, input)
	return bs
}

func uint32tobyte(input uint32) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, input)
	return bs
}

func bytetouint16(input []byte) uint16 {
	return binary.LittleEndian.Uint16(input)
}

func bytetouint32(input []byte) uint32 {
	return binary.LittleEndian.Uint32(input)
}
