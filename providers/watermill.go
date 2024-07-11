package providers

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/sony/gobreaker"
)

func NewWatermill() *message.Router {
	router, err := message.NewRouter(
		message.RouterConfig{},
		watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}
	retryMiddleware := middleware.Retry{
		MaxRetries:      1,
		InitialInterval: time.Millisecond * 10,
	}
	router.AddPlugin(plugin.SignalsHandler)

	middlewares := []message.HandlerMiddleware{
		middleware.Recoverer,
		retryMiddleware.Middleware,
		middleware.NewThrottle(10, time.Second).Middleware,
		middleware.CorrelationID,
		middleware.NewCircuitBreaker(gobreaker.Settings{
			MaxRequests:   100,
			Interval:      10,
			Timeout:       1 * time.Second,
			ReadyToTrip:   nil,
			OnStateChange: nil,
			IsSuccessful:  nil,
		}).Middleware,
	}

	router.AddMiddleware(middlewares...)
	return router
}
