package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"payment-service/config"
	"payment-service/models"

	"cloud.google.com/go/pubsub"
)

func PublishPaymentCreated(payment models.Payment) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.GcpPojectId)
	if err != nil {
		log.Printf(`{"message":"PubSub client error to publish payment event: %v", "service":"payment", "severity":"ERROR"}`, err)
		return
	}

	log.Println(`{"message":"Client created in PubSub request", "service":"payment", "severity":"INFO"}`)

	topic := client.Topic(config.PaymentTopicId)
	data, _ := json.Marshal(payment)
	result := topic.Publish(ctx, &pubsub.Message{Data: data})

	_, err = result.Get(ctx)
	if err != nil {
		log.Printf(`{"message":"Failed to publish payment event: %v", "service":"payment", "severity":"ERROR"}`, err)
	}

	log.Printf(`{"message":"Successfully published payment event: %v", "service":"payment", "severity":"INFO"}`, payment)
}
