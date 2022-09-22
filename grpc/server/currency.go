package server

import (
	"context"
	"github.com/abassGarane/currency/protos"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{
		log: l,
	}
}
func (c *Currency) GetRate(ctx context.Context, request *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate :: ", "Base:: ", request.GetBase(), "Destinaltion:: ", request.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
