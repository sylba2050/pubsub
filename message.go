package pubsub

func NewMessage(h Header, p ...Payload) Message {
	return &M{
		header:   h,
		payloads: p,
	}
}

type Message interface {
	ToBytes() ([]byte, error)
	AppendPayloads(...Payload) error
}

type M struct {
	header   Header
	payloads []Payload
}

func (m *M) ToBytes() ([]byte, error) {
	var payloadBytes []byte
	for _, payload := range m.payloads {
		p, err := payload.ToBytes()
		if err != nil {
			return nil, err
		}
		payloadBytes = append(payloadBytes, p...)
	}

	messageLength := uint16(len(payloadBytes))
	m.header.SetLength(messageLength)

	var message []byte

	header, err := m.header.ToBytes()
	if err != nil {
		return nil, err
	}
	message = append(message, header...)
	message = append(message, payloadBytes...)

	return message, nil
}

func (m *M) AppendPayloads(payloads ...Payload) error {
	for _, payload := range payloads {
		m.payloads = append(m.payloads, payload)
	}
	return nil
}
