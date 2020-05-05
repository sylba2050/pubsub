package pubsub

// DataType is Data in payload
type DataType uint16

const (
	ConnectToken DataType = 0x4001
	MessageID    DataType = 0x4002
	MessageBody  DataType = 0x4003
	TopicID      DataType = 0x4004
	SubscriberID DataType = 0x4005
)

const PayloadHeaderSize = 4

type Payload interface {
	SetType(MessageType) error
	GetType() (MessageType, error)
	SetLength(uint16) error
	GetLength() (uint16, error)
	SetValue([]byte) error
	GetValue() ([]byte, error)
	ToBytes() ([]byte, error)
}

type P struct {
	typ    MessageType
	length uint16
	value  []byte
}

func (p *P) SetType(m MessageType) error {
	return nil
}

func (p P) GetType() (MessageType, error) {
	return 0x0000, nil
}

func (p *P) SetLength(length uint16) error {
	return nil
}

func (p P) GetLength() (uint16, error) {
	return 0, nil
}

func (p *P) ToBytes() ([]byte, error) {
	return nil, nil
}
