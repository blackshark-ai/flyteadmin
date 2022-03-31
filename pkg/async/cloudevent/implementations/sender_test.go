package implementations

import (
	"context"
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/stretchr/testify/assert"
)

type mockCloudEventClient struct{}

func (s mockCloudEventClient) Request(ctx context.Context, event event.Event) (*event.Event, protocol.Result) {
	return nil, nil
}

func (s mockCloudEventClient) StartReceiver(ctx context.Context, fn interface{}) error {
	return nil
}

func (s mockCloudEventClient) Send(ctx context.Context, event event.Event) protocol.Result {
	return nil
}

func TestPubSubSender(t *testing.T) {
	pubSubSender := PubSubSender{&testCloudEventPublisher}
	cloudEvent := cloudevents.NewEvent()
	err := pubSubSender.Send(context.Background(), "test", cloudEvent)
	assert.Nil(t, err)
}

func TestKafkaSender(t *testing.T) {
	pubSubSender := KafkaSender{mockCloudEventClient{}}
	cloudEvent := cloudevents.NewEvent()
	err := pubSubSender.Send(context.Background(), "test", cloudEvent)
	assert.Nil(t, err)
}
