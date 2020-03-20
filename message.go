package pubsub

// MessageType is MessageType in header
type MessageType uint16

const (
	Connect                 MessageType = 0x0001
	ConnectAckSuccess       MessageType = 0x0002
	ConnectAckFailure       MessageType = 0x0003
	ReconnectRequest        MessageType = 0x0004
	Disconnect              MessageType = 0x0005
	Subscribe               MessageType = 0x0020
	UnSubscribe             MessageType = 0x0021
	SubscribeAckSuccess     MessageType = 0x0022
	SubscribePermisionError MessageType = 0x0023
	Publish                 MessageType = 0x0040
	PublishAckSuccess       MessageType = 0x0041
	PublishPermissionError  MessageType = 0x0042
	Ping                    MessageType = 0x0060
	Pong                    MessageType = 0x0061
	AddSubscriber           MessageType = 0x0100
	RemoveSubscriber        MessageType = 0x0101
	GetSubscribersRequest   MessageType = 0x1000
	SubscribersList         MessageType = 0x1001
	CloseRequest            MessageType = 0x2000
)

// DataType is Data in header
type DataType uint16

const (
	MessageBody  DataType = 0x4003
	ConnectToken DataType = 0x4001
	MessageID    DataType = 0x4002
	TopicID      DataType = 0x4004
	SubscriberID DataType = 0x4005
)

func NewData(h Header, p ...Payload) Data {
	return Data{
		header:  h,
		payload: p,
	}
}

type Data struct {
	header  Header
	payload []Payload
}

func (d *Data) BuildData() []byte {
	headertype := uint16tobyte(uint16(d.header.typ))
	senderTimestamp := uint32tobyte(d.header.senderTimestamp)
	receiverTimestamp := uint32tobyte(d.header.receiverTimestamp)

	var payload []byte
	for _, p := range d.payload {
		payload = append(payload, p.Bytes()...)
	}

	size := uint16tobyte(uint16(len(payload)) + HeaderSize)

	data := make([]byte, 0, len(headertype)+len(size)+len(senderTimestamp)+len(receiverTimestamp)+len(payload))
	data = append(data, headertype...)
	data = append(data, size...)
	data = append(data, senderTimestamp...)
	data = append(data, receiverTimestamp...)
	data = append(data, payload...)

	return data
}

const HeaderSize = 8

type Header struct {
	typ               MessageType
	size              uint16
	receiverTimestamp uint32
	senderTimestamp   uint32
}

type Payload struct {
	typ   MessageType
	size  uint16
	value []byte
}

func (p *Payload) Bytes() []byte {
	return append(uint16tobyte(uint16(p.typ)), append(uint16tobyte(p.size), p.value...)...)
}

func NewPublishPayload(message []byte) Payload {
	// FIXME: uint16のサイズを越えていたらエラー
	return Payload{
		typ:   Publish,
		size:  uint16(len(message)),
		value: message,
	}
}
