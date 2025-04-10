package azuresb

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type AzureSBClient struct {
	s *azservicebus.Sender
}

func NewAzureSBClient(conn, topic string) (*AzureSBClient, error) {
	client, err := azservicebus.NewClientFromConnectionString(conn, &azservicebus.ClientOptions{})
	if err != nil {
		return nil, err
	}
	sender, err := client.NewSender(topic, &azservicebus.NewSenderOptions{})
	if err != nil {
		return nil, err
	}
	return &AzureSBClient{
		s: sender,
	}, nil
}

func (c *AzureSBClient) SendBatch(ctx context.Context, payloads ...interface{}) error {
	if (len(payloads)) == 0 {
		return nil
	}
	batch, err := c.s.NewMessageBatch(ctx, nil)
	if err != nil {
		return err
	}
	for _, m := range payloads {
		payload, err := json.Marshal(m)
		if err != nil {
			return err
		}
		err = batch.AddMessage(&azservicebus.Message{Body: payload}, nil)
		if err != nil {
			return err
		}
	}
	return c.s.SendMessageBatch(ctx, batch, nil)
}
