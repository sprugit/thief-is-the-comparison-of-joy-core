package model

import "time"

type OrderType int

const (
	UnknownType OrderType = iota
	BUY
	SELL
)

type OrderStatus int

const (
	UnknownStatus OrderStatus = iota
	PENDING
	PROCESSED
	EXPIRED
)

type Order struct {
	UserID         string    `json:"user_id"`
	Ticker         string    `json:"ticker"`
	OrderType      OrderType `json:"order_type"`
	StrikePrice    float64   `json:"strike_price"`
	Quantity       int64     `json:"quantity"`
	Timestamp      time.Time `json:"timestamp"`
	ExpirationDate time.Time `json:"expiration_date"`
}
