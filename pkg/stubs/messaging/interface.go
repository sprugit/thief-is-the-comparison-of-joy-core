package messaging

import "github.com/sprugit/thief-is-the-comparison-of-joy-core/pkg/model"

type Publisher interface {
	Publish(order model.Order) error
}

type Subscriber interface {
	GetOrderChannel() <-chan model.Order
}
