package service

import (
	"context"
	"fmt"
	"time"

	"github.com/automotechnologies/doitpay-go"
	"github.com/automotechnologies/doitpay-go/invoice"
	"github.com/automotechnologies/doitpay/sdk/golang/sample/helper"
	"github.com/brianvoe/gofakeit/v6"
)

type InvoiceServiceMethod interface {
	CreateInvoice(ctx context.Context, req CreateInvoiceRequest) (string, error)
	GetInvoiceList(ctx context.Context) ([]InvoiceDetail, error)
	GetInvoiceDetail(ctx context.Context, id string) (InvoiceDetail, error)
}

func NewInvoiceService(client *doitpay.APIClient, checkoutLink string) InvoiceServiceMethod {
	return InvoiceService{
		client:       client,
		checkoutLink: checkoutLink,
	}
}

type InvoiceService struct {
	client       *doitpay.APIClient
	checkoutLink string
}

type CreateInvoiceRequest struct {
	ProductID string
}

const checkoutUrl = "https://dev.doit-frontend.pages.dev/checkout/"

func (p InvoiceService) CreateInvoice(ctx context.Context, req CreateInvoiceRequest) (string, error) {
	val, ok := mapProductID[req.ProductID]
	if !ok {
		return "", fmt.Errorf("invalid product id")
	}

	Invoice := p.client.InvoiceAPI.CreateInvoice(ctx)
	adr := gofakeit.Address()
	res, http, err := Invoice.Request(invoice.InternalWebControllersMerchantApiv1InvoiceInvoiceRequest{
		Amount:         &val.Amount,
		AmountCurrency: helper.StringPointer("IDR"),
		Country:        helper.StringPointer("IDN"),
		Customer: &invoice.InternalWebControllersMerchantApiv1InvoiceCustomer{
			Addresses: helper.StringPointer(adr.Address),
			City:      helper.StringPointer(adr.City),
			Country:   helper.StringPointer(adr.Country),
			Email:     helper.StringPointer("doit-sample@mailinator.com"),
			Name:      helper.StringPointer(gofakeit.Name()),
			Notes:     &val.Notes,
		},
		Items: []invoice.InternalWebControllersMerchantApiv1InvoiceItem{
			{
				SKU:      helper.StringPointer("001"),
				Name:     &val.Name,
				Notes:    &val.Notes,
				Price:    &val.Amount,
				Quantity: helper.Int32Pointer(1),
			},
		},
		Notes:     &val.Notes,
		TimeBegin: helper.StringPointer(getCurrentTimeFormatted()),
		TimeEnd:   helper.StringPointer(getExpiredTimeFormatted()),
		RedirectUrl: &invoice.InternalWebControllersMerchantApiv1InvoiceRedirectURL{
			Success: helper.StringPointer("http://localhost:8000/"),
			Cancel:  helper.StringPointer("http://localhost:8000/"),
		},
	}).Execute()

	if err != nil {
		return "", fmt.Errorf("Got Error while execute Invoice : %v", err)
	}

	if http != nil && http.StatusCode > 300 {
		return "", fmt.Errorf("Got Status code %v while execute Invoice", http.StatusCode)
	}

	data, ok := res.GetDataOk()
	if !ok {
		return "", fmt.Errorf("Got Invalid body response")
	}

	return data.GetInvoiceLinkId(), nil
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

type InvoiceDetail struct {
	InvoiceID   string
	ReferenceID string
	Name        string
	Amount      float64
	Status      string
	PaymentLink string
}

func (p InvoiceService) GetInvoiceDetail(ctx context.Context, id string) (InvoiceDetail, error) {
	res, http, err := p.client.InvoiceAPI.GetInvoiceById(ctx, id).Execute()

	if err != nil {
		return InvoiceDetail{}, fmt.Errorf("Got Error while execute Invoice detail : %v", err)
	}

	if http != nil && http.StatusCode > 300 {
		return InvoiceDetail{}, fmt.Errorf("Got Status code %v while execute Invoice detail", http.StatusCode)
	}

	data, ok := res.GetDataOk()
	if !ok {
		return InvoiceDetail{}, fmt.Errorf("Got Invalid body response")
	}

	return convertDataToInvoiceDetail(data), nil
}

func convertDataToInvoiceDetail(data *invoice.InternalWebControllersMerchantApiv1InvoiceInvoiceDetailResponse) InvoiceDetail {
	items := data.GetItems()
	paymentLink := checkoutUrl + data.GetInvoiceLinkId()

	return InvoiceDetail{
		InvoiceID:   data.GetInvoiceLinkId(),
		ReferenceID: data.GetExternalId(),
		Name:        items[0].GetName(),
		Amount:      float64(data.GetAmount()),
		Status:      data.GetStatus(),
		PaymentLink: paymentLink,
	}
}

func (p InvoiceService) GetInvoiceList(ctx context.Context) ([]InvoiceDetail, error) {
	res, http, err := p.client.InvoiceAPI.GetInvoices(ctx).Execute()

	if err != nil {
		return nil, fmt.Errorf("Got Error while execute Invoice list : %v", err)
	}

	if http != nil && http.StatusCode > 300 {
		return nil, fmt.Errorf("Got Status code %v while execute Invoice list", http.StatusCode)
	}

	data := res.GetData()

	var list []InvoiceDetail

	for _, x := range data {
		list = append(list, convertDataToInvoiceDetail(&x))
	}

	return list, nil
}
