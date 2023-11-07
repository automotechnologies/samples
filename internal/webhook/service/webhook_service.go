package service

import (
	"context"
	"sync"
)

type WebhookServiceMethod interface {
	CreateWebhook(ctx context.Context, request WebhookDetail) error
	GetWebhook(ctx context.Context) []WebhookDetail
}

type WebhookService struct {
}

func NewWebhookService() WebhookServiceMethod {
	return WebhookService{}
}

var listWebhook []WebhookDetail

type WebhookDetail struct {
	ReferenceID          string
	VirtualAccountLinkID string
	VirtualAccountNumber string
	Amount               int
	Status               string
}

func (w WebhookService) CreateWebhook(ctx context.Context, request WebhookDetail) error {

	var mtx sync.Mutex
	mtx.Lock()
	listWebhook = append(listWebhook, request)
	mtx.Unlock()

	return nil
}

func (w WebhookService) GetWebhook(ctx context.Context) []WebhookDetail {
	return listWebhook
}
