// Package classification Product Api.
//
// # Documentation for Product Api
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// License: BSD-3
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"github.com/abassGarane/microservices/grpc/protos"
	"log"
)

type ProductHandler struct {
	l  *log.Logger
	cc *protos.CurrencyClient
}

func NewProducts(l *log.Logger, c *protos.CurrencyClient) *ProductHandler {
	return &ProductHandler{l, c}
}

type ProductKey struct{}
