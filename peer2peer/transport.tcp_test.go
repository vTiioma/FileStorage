package peer2peer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	port := ":4000"
	tr := CreateTcpTransport(port)

	assert.Equal(t, tr.listenAddress, port)
	assert.Nil(t, tr.Listen())
}
