package pubsub

func NewMessage(h Header, p ...Payload) Message {
	return &message{
		header:   h,
		payloads: p,
	}
}

type Message interface {
	ToBytes() ([]byte, error)
	AppendPayloads(...Payload) error
}

type message struct {
	header   Header
	payloads []Payload
}

func (m *message) ToBytes() ([]byte, error) {
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

func (m *message) AppendPayloads(payloads ...Payload) error {
	for _, payload := range payloads {
		m.payloads = append(m.payloads, payload)
	}
	return nil
}
