package pubsub

import (
	"errors"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

// MessageType is MessageType in header
type MessageType uint16

func NewMessageType(input uint16) MessageType {
	return MessageType(input)
}

func NewMessageTypeFromByte(input []byte) (MessageType, error) {
	if len(input) != 2 {
		return NoneMessageType, errors.New("Can't parse byte to header")
	}
	return NewMessageType(byte2uint16(input)), nil
}

const HeaderSize = 12

const (
	NoneMessageType                 MessageType = 0x0000
	Connect                         MessageType = 0x0001
	ConnectAckSuccess               MessageType = 0x0002
	ConnectAckFailure               MessageType = 0x0003
	ReconnectRequest                MessageType = 0x0004
	Disconnect                      MessageType = 0x0005
	CreateNewTopicRequest           MessageType = 0x0010
	Subscribe                       MessageType = 0x0020
	UnSubscribe                     MessageType = 0x0021
	SubscribeAckSuccess             MessageType = 0x0022
	SubscribePermisionError         MessageType = 0x0023
	Publish                         MessageType = 0x0040
	PublishAckSuccess               MessageType = 0x0041
	PublishPermissionError          MessageType = 0x0042
	Ping                            MessageType = 0x0060
	Pong                            MessageType = 0x0061
	AddSubscriber                   MessageType = 0x0100
	AddSubscriberPermissionError    MessageType = 0x0101
	RemoveSubscriber                MessageType = 0x0102
	RemoveSubscriberPermissionError MessageType = 0x0103
	GetSubscribersRequest           MessageType = 0x1000
	GetSubscribersPermissionError   MessageType = 0x1001
	SubscribersList                 MessageType = 0x1002
	CloseRequest                    MessageType = 0x2000
)

type Header interface {
	SetType(MessageType) error
	GetType() (MessageType, error)
	SetLength(uint16) error
	GetLength() (uint16, error)
	SetSenderTimestamp(uint32) error
	GetSenderTimestamp() (uint32, error)
	SetReceiverTimestamp(uint32) error
	GetReceiverTimestamp() (uint32, error)
	ToBytes() ([]byte, error)
}

type header struct {
	typ               MessageType
	length            uint16
	senderTimestamp   uint32
	receiverTimestamp uint32
}

func (h *header) SetType(m MessageType) error {
	h.typ = m
	return nil
}

func (h header) GetType() (MessageType, error) {
	return h.typ, nil
}

func (h *header) SetLength(length uint16) error {
	h.length = length
	return nil
}

func (h header) GetLength() (uint16, error) {
	return h.length, nil
}

func (h *header) SetSenderTimestamp(timestamp uint32) error {
	h.senderTimestamp = timestamp
	return nil
}

func (h header) GetSenderTimestamp() (uint32, error) {
	return h.senderTimestamp, nil
}

func (h *header) SetReceiverTimestamp(timestamp uint32) error {
	h.receiverTimestamp = timestamp
	return nil
}

func (h header) GetReceiverTimestamp() (uint32, error) {
	return h.receiverTimestamp, nil
}

func (h *header) ToBytes() ([]byte, error) {
	var bytes []byte

	bytes = append(bytes, uint16tobyte(uint16(h.typ))...)
	bytes = append(bytes, uint16tobyte(h.length)...)
	bytes = append(bytes, uint32tobyte(h.senderTimestamp)...)
	bytes = append(bytes, uint32tobyte(h.receiverTimestamp)...)

	return bytes, nil
}

func NewHeader(m MessageType) (Header, error) {
	return &header{typ: m}, nil
}

func ReadHeader(conn net.Conn) (Header, error) {
	buf := make([]byte, HeaderSize)
	n, err := conn.Read(buf)
	if n < HeaderSize {
		err := errors.New("Can't read header")
		log.Error().Err(err).Send()
		return &header{}, err
	}
	if err != nil {
		fmt.Printf("Header read error: %s\n", err)
	}

	messageType, err := NewMessageTypeFromByte(buf[:2])

	h, err := NewHeader(messageType)
	h.SetLength(byte2uint16(buf[2:4]))
	h.SetSenderTimestamp(byte2uint32(buf[4:8]))
	h.SetReceiverTimestamp(byte2uint32(buf[8:12]))

	return h, nil
}
