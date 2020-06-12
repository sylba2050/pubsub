package pubsub

import (
	"errors"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

// DataType is Data in payload
type DataType uint16

const (
	NoneDataType DataType = 0x0000
	ConnectToken DataType = 0x4001
	MessageID    DataType = 0x4002
	MessageBody  DataType = 0x4003
	TopicID      DataType = 0x4004
	SubscriberID DataType = 0x4005
)

func NewDataType(input uint16) DataType {
	return DataType(input)
}

func NewDataTypeFromByte(input []byte) (DataType, error) {
	if len(input) != 2 {
		return NoneDataType, errors.New("Can't parse byte to payload header")
	}
	return NewDataType(byte2uint16(input)), nil
}

const PayloadHeaderSize = 4

type Payload interface {
	SetType(DataType) error
	GetType() (DataType, error)
	SetLength(uint16) error
	GetLength() (uint16, error)
	SetValue([]byte) error
	GetValue() ([]byte, error)
	ToBytes() ([]byte, error)
}

type payload struct {
	typ    DataType
	length uint16
	value  []byte
}

func (p *payload) SetType(m DataType) error {
	p.typ = m
	return nil
}

func (p payload) GetType() (DataType, error) {
	return p.typ, nil
}

func (p *payload) SetLength(length uint16) error {
	p.length = length
	return nil
}

func (p payload) GetLength() (uint16, error) {
	return p.length, nil
}

func (p *payload) SetValue(value []byte) error {
	p.value = value
	return nil
}

func (p payload) GetValue() ([]byte, error) {
	return p.value, nil
}

func (p *payload) ToBytes() ([]byte, error) {
	var bytes []byte

	bytes = append(bytes, uint16tobyte(uint16(p.typ))...)
	bytes = append(bytes, uint16tobyte(p.length)...)
	bytes = append(bytes, p.value...)

	return bytes, nil
}

func NewPayload(d DataType) (Payload, error) {
	return &payload{typ: d}, nil
}

func readPayloadHeaderFromConn(conn net.Conn) (Payload, error) {
	buf := make([]byte, PayloadHeaderSize)
	n, err := conn.Read(buf)
	if n < PayloadHeaderSize {
		err := errors.New("Can't read payload header")
		log.Error().Err(err).Send()
		return &payload{}, err
	}
	if err != nil {
		fmt.Printf("Payload header read error: %s\n", err)
	}

	dataType, err := NewDataTypeFromByte(buf[:2])

	p, err := NewPayload(dataType)
	p.SetLength(byte2uint16(buf[2:4]))

	return p, nil
}

func readPayloadDataFromConn(conn net.Conn, p Payload) (Payload, error) {
	dataLength, err := p.GetLength()
	if dataLength == 0 {
		return &payload{}, errors.New("payload data length must not be 0")
	}

	buf := make([]byte, dataLength)
	n, err := conn.Read(buf)
	if n < int(dataLength) {
		err := errors.New("Can't read payload data")
		log.Error().Err(err).Send()
		return &payload{}, err
	}
	if err != nil {
		fmt.Printf("Payload data read error: %s\n", err)
	}

	p.SetValue(buf)

	return p, nil
}

func ReadPayload(conn net.Conn) (Payload, error) {
	p, err := readPayloadHeaderFromConn(conn)
	if err != nil {
		return &payload{}, err
	}

	p, err = readPayloadDataFromConn(conn, p)
	if err != nil {
		return &payload{}, err
	}
	return p, nil
}
