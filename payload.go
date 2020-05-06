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
	typ    DataType
	length uint16
	value  []byte
}

func (p *P) SetType(m DataType) error {
	p.typ = m
	return nil
}

func (p P) GetType() (DataType, error) {
	return p.typ, nil
}

func (p *P) SetLength(length uint16) error {
	p.length = length
	return nil
}

func (p P) GetLength() (uint16, error) {
	return p.length, nil
}

func (p *P) ToBytes() ([]byte, error) {
	var bytes []byte

	bytes = append(bytes, uint16tobyte(uint16(p.typ))...)
	bytes = append(bytes, uint16tobyte(p.length)...)
	bytes = append(bytes, p.value...)

	return bytes, nil
}

func NewPayload(d DataType) (P, error) {
	return P{typ: d}, nil
}
