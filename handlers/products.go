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

import "log"

type ProductHandler struct{
  l *log.Logger
}


func NewProducts (l *log.Logger) *ProductHandler{
  return &ProductHandler{l}
}

type ProductKey struct{}






