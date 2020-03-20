package pubsub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayloadBytes(t *testing.T) {
	m := []byte("Hello, World!")
	assert.Equal(t, len(m), 13)

	p := NewPublishPayload(m)
	bytes := p.Bytes()

	assert.Equal(t, bytetouint16(bytes[0:2]), uint16(Publish))
	assert.Equal(t, bytetouint16(bytes[2:4]), uint16(17))
	assert.Equal(t, string(bytes[4:]), "Hello, World!")
}
