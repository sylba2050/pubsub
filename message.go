package pubsub

// MessageType is MessageType in header
type MessageType int16

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

type Payload struct {
	id                string
	t                 MessageType
	receiverTimestamp int32
	senderTimestamp   int32
}

// DataType is Data in header
type DataType int16

const (
	Message      DataType = 0x4003
	ConnectToken DataType = 0x4001
	MessageID    DataType = 0x4002
	TopicID      DataType = 0x4004
	SubscriberID DataType = 0x4005
)
