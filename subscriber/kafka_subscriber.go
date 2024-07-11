package subscriber

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

func NewKafkaSubscriber(brokers []string, consumerGroup string) *kafka.Subscriber {
	logger := watermill.NewStdLogger(true, false)
	sub, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:       brokers,
			Unmarshaler:   kafka.DefaultMarshaler{},
			ConsumerGroup: consumerGroup,
		},
		logger,
	)
	if err != nil {
		panic(err)
	}
	return sub
}
