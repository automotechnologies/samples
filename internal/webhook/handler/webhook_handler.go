package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/automotechnologies/doitpay/sdk/golang/sample/internal/webhook/service"
	"github.com/labstack/echo/v4"
)

type WebhookHandler struct {
	service service.WebhookServiceMethod
}

func NewWebhookHandler(service service.WebhookServiceMethod) WebhookHandler {
	return WebhookHandler{
		service: service,
	}
}

type Request struct {
	Event       string    `json:"event"`
	BusinessID  int       `json:"business_id"`
	CreatedTime time.Time `json:"created_time"`
	Data        struct {
		VirtualAccountLinkID string    `json:"virtual_account_link_id"`
		VirtualAccountNumber string    `json:"virtual_account_number"`
		VirtualAccountRefID  string    `json:"virtual_account_ref_id"`
		CustomerName         string    `json:"customer_name"`
		CustomerEmail        string    `json:"customer_email"`
		CustomerPhone        string    `json:"customer_phone"`
		Currency             string    `json:"currency"`
		Amount               int       `json:"amount"`
		ExpirationDate       time.Time `json:"expiration_date"`
		Status               string    `json:"status"`
	} `json:"data"`
}

func (p WebhookHandler) WebhookHandler(c echo.Context) error {
	var requestData Request
	if err := c.Bind(&requestData); err != nil {
		return err
	}

	p.service.CreateWebhook(context.Background(), service.WebhookDetail{
		ReferenceID:          requestData.Data.VirtualAccountRefID,
		VirtualAccountLinkID: requestData.Data.VirtualAccountLinkID,
		VirtualAccountNumber: requestData.Data.VirtualAccountNumber,
		Amount:               requestData.Data.Amount,
		Status:               requestData.Data.Status,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (p WebhookHandler) GetWebhookHandler(c echo.Context) error {
	var requestData Request
	if err := c.Bind(&requestData); err != nil {
		return err
	}

	data := p.service.GetWebhook(context.Background())

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    data,
	})
}
