package azuresb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type AzureSBClient struct { //nolint:revive
	s *azservicebus.Sender
}

func NewAzureSBClient(conn, topic string) (*AzureSBClient, error) {
	client, err := azservicebus.NewClientFromConnectionString(conn, &azservicebus.ClientOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not create client: %w", err)
	}
	sender, err := client.NewSender(topic, &azservicebus.NewSenderOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not create sender: %w", err)
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
		return fmt.Errorf("could not create message batch: %w", err)
	}
	for _, m := range payloads {
		payload, err := json.Marshal(m)
		if err != nil {
			return fmt.Errorf("could not marshal payload: %w", err)
		}
		err = batch.AddMessage(&azservicebus.Message{Body: payload}, nil)
		if err != nil {
			return fmt.Errorf("could not add message: %w", err)
		}
	}
	err = c.s.SendMessageBatch(ctx, batch, nil)
	if err != nil {
		return fmt.Errorf("could not send message batch: %w", err)
	}
	return nil
}
