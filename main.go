package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"

	"kafka2mysql/configs"
	"kafka2mysql/providers"
	"kafka2mysql/publisher"
	"kafka2mysql/subscriber"
)

var (
	apiTopicPublisher = "api2kafka"
)

func main() {
	config := configs.LoadConfig()

	// config database
	db := providers.NewMySql(config.Database)
	err := db.AutoMigrate(&ReceivedMessage{})
	if err != nil {
		panic(err)
	}

	// config kafka publisher
	kafkaPublisher, err := publisher.NewPublisher(config.Kafka.Brokers)
	if err != nil {
		panic(err)
	}

	// configuration kafka router handler
	router := providers.NewWatermill()
	router.AddHandler("kafka2mysql",
		apiTopicPublisher,
		subscriber.NewKafkaSubscriber(config.Kafka.Brokers, apiTopicPublisher),
		"",
		nil,
		func(msg *message.Message) ([]*message.Message, error) {
			var msgPayload ReceivedMessage
			err = json.Unmarshal(msg.Payload, &msgPayload)
			if err != nil {
				return nil, err
			}
			db.Create(&ReceivedMessage{
				Message:   msgPayload.Message,
				CreatedAt: msgPayload.CreatedAt,
			})
			return nil, nil
		},
	)

	// configuration gin
	engine := providers.NewGinServer()
	engine.Handle(http.MethodGet, "/api/v1/kafka2mysql", func(context *gin.Context) {
		msgPayload := ReceivedMessage{
			CreatedAt: time.Now(),
			Message:   context.Query("message"),
		}
		payload, err := json.Marshal(msgPayload)
		if err != nil {
			context.Status(http.StatusBadRequest)
			return
		}
		msg := message.NewMessage(watermill.NewUUID(), payload)
		// Use a middleware to set the correlation ID, it's useful for debugging
		middleware.SetCorrelationID(watermill.NewShortUUID(), msg)
		err = kafkaPublisher.Publish(apiTopicPublisher, msg)
		if err != nil {
			fmt.Println("cannot publish message:", err)
			context.Status(http.StatusBadRequest)
			return
		}
	})

	go func() {
		err = engine.Run(fmt.Sprintf(":%d", config.Server.Port))
		if err != nil {
			panic(err)
		}
	}()
	if err = router.Run(context.Background()); err != nil {
		panic(err)
	}
}

type ReceivedMessage struct {
	Id        uint `gorm:"primaryKey"`
	Message   string
	CreatedAt time.Time
}
