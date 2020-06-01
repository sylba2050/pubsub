package pubsub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayloadBytes(t *testing.T) {
	m := []byte("Hello, World!")
	assert.Equal(t, len(m), 13)

	p, _ := NewPayload(MessageBody)
	p.SetLength(uint16(len(m)))
	p.SetValue(m)
	bytes, _ := p.ToBytes()

	assert.Equal(t, byte2uint16(bytes[0:2]), uint16(MessageBody))
	assert.Equal(t, byte2uint16(bytes[2:4]), uint16(13))
	assert.Equal(t, string(bytes[4:]), "Hello, World!")
}
