package amqp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/go-asyncapi/amqp"
)

func TestMessageWithInfo(t *testing.T) {
	m := amqp.MessageWithInfo(nil, amqp.Info{
		Exchanges: []string{"some-exchange", "another-exchange"},
		VHost:     "some-vhost",
	})

	assert.Equal(t, "AMQP VHost: some-vhost.\n\nAMQP Exchanges: some-exchange, another-exchange.", m.Description)
	assert.Equal(t, []string{"some-exchange", "another-exchange"}, m.MapOfAnything[amqp.Exchanges])
	assert.Equal(t, "some-vhost", m.MapOfAnything[amqp.VHost])
}
