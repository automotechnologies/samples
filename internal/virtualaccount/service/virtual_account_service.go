package service

import (
	"context"
	"fmt"
	"time"

	"github.com/automotechnologies/doitpay-go"
	"github.com/automotechnologies/doitpay-go/virtualaccount"
	"github.com/automotechnologies/doitpay/sdk/golang/sample/helper"
	"github.com/brianvoe/gofakeit/v6"
)

type VirtualAccountServiceMethod interface {
	CreateVirtualAccount(ctx context.Context, req CreateVirtualAccountRequest) (string, error)
	GetVirtualAccountList(ctx context.Context) ([]VirtualAccountDetail, error)
}

func NewVirtualAccountService(client *doitpay.APIClient) VirtualAccountServiceMethod {
	return VirtualAccountService{
		client: client,
	}
}

type VirtualAccountService struct {
	client *doitpay.APIClient
}

type CreateVirtualAccountRequest struct {
	Suffix string
	Amount float32
}

func (p VirtualAccountService) CreateVirtualAccount(ctx context.Context, req CreateVirtualAccountRequest) (string, error) {
	va := p.client.VirtualAccountAPI.CreateVirtualAccount(ctx)
	res, http, err := va.Request(virtualaccount.InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest{
		VirtualAccountSuffix: helper.StringPointer(req.Suffix),
		Amount:               helper.Float32Pointer(req.Amount),
		IsClosed:             helper.BoolPointer(true),
		IsReusable:           helper.BoolPointer(false),
		Currency:             helper.StringPointer("IDR"),
		Customer: &virtualaccount.InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountCustomer{
			Email: helper.StringPointer("doit-sample@mailinator.com"),
			Name:  helper.StringPointer(gofakeit.Name()),
		},
		ExpirationDate:    helper.StringPointer(getExpiredTimeFormatted()),
		ReferenceId:       helper.StringPointer(gofakeit.UUID()),
		PaymentMethodCode: helper.StringPointer("sampoerna-test"),
	}).Execute()
	if err != nil {
		return "", fmt.Errorf("Got Error while execute VirtualAccount : %v", err)
	}
	if http != nil && http.StatusCode > 300 {
		return "", fmt.Errorf("Got Status code %v while execute VirtualAccount", http.StatusCode)
	}

	data, ok := res.GetDataOk()
	if !ok {
		return "", fmt.Errorf("Got Invalid body response")
	}

	return data.GetVirtualAccountLinkId(), nil
}

func getCurrentTimeFormatted() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")
	return formattedTime
}

func getExpiredTimeFormatted() string {
	currentTime := time.Now().Add(24 * time.Hour)
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")
	return formattedTime
}

type VirtualAccountDetail struct {
	ID                   string
	ReferenceID          string
	VirtualAccountNumber string
	Amount               float64
	Status               string
}

func convertDataToVirtualAccountDetail(data *virtualaccount.InternalWebControllersMerchantApiv1VirtualaccountVirtualAccountResponse) VirtualAccountDetail {
	return VirtualAccountDetail{
		ID:                   data.GetVirtualAccountLinkId(),
		ReferenceID:          data.GetReferenceId(),
		VirtualAccountNumber: data.GetVirtualAccountNumber(),
		Amount:               float64(data.GetAmount()),
		Status:               data.GetStatus(),
	}
}

func (p VirtualAccountService) GetVirtualAccountList(ctx context.Context) ([]VirtualAccountDetail, error) {
	res, http, err := p.client.VirtualAccountAPI.GetVirtualAccounts(ctx).Execute()

	if err != nil {
		return nil, fmt.Errorf("Got Error while execute VirtualAccount list : %v", err)
	}

	if http != nil && http.StatusCode > 300 {
		return nil, fmt.Errorf("Got Status code %v while execute VirtualAccount list", http.StatusCode)
	}

	data := res.GetData()

	var list []VirtualAccountDetail

	for _, x := range data {
		list = append(list, convertDataToVirtualAccountDetail(&x))
	}

	return list, nil
}
