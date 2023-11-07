package service

import (
	"context"
	"fmt"

	"github.com/automotechnologies/doitpay-go"
	"github.com/automotechnologies/doitpay-go/simulate"
	"github.com/automotechnologies/doitpay/sdk/golang/sample/helper"
)

type SimulateServiceMethod interface {
	DoSimulate(ctx context.Context, req SimulateRequest) error
}

func NewSimulateService(client *doitpay.APIClient) SimulateServiceMethod {
	return SimulateService{
		client: client,
	}
}

type SimulateService struct {
	client *doitpay.APIClient
}

type SimulateRequest struct {
	VirtualAccountNumber string
}

func (p SimulateService) DoSimulate(ctx context.Context, req SimulateRequest) error {
	cli := p.client.SimulateAPI.SimulatePayment(ctx)
	http, err := cli.Request(simulate.InternalWebControllersMerchantApiv1SimulateSimulateRequest{
		AccountNumber: helper.StringPointer(req.VirtualAccountNumber),
	}).Execute()
	if err != nil {
		return fmt.Errorf("Got Error while execute Simulate : %v", err)
	}
	if http != nil && http.StatusCode > 300 {
		return fmt.Errorf("Got Status code %v while execute Simulate", http.StatusCode)
	}

	return nil
}
