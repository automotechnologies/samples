package handler

import (
	"context"
	"net/http"

	"github.com/automotechnologies/doitpay/sdk/golang/sample/internal/invoice/service"
	"github.com/labstack/echo/v4"
)

type InvoiceHandler struct {
	svc service.InvoiceServiceMethod
}

func NewInvoiceHandler(svc service.InvoiceServiceMethod) InvoiceHandler {
	return InvoiceHandler{
		svc: svc,
	}
}

type InvoiceRequest struct {
	ProductID string `json:"product_id"`
}

func (p InvoiceHandler) CreateInvoiceHandler(c echo.Context) error {
	requestData := new(InvoiceRequest)

	// Bind the JSON request to the RequestData struct
	if err := c.Bind(requestData); err != nil {
		return err
	}

	trxID, err := p.svc.CreateInvoice(context.Background(), service.CreateInvoiceRequest{
		ProductID: requestData.ProductID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":    "Success",
		"Invoice_id": trxID,
	})
}

func (p InvoiceHandler) GetInvoiceListHandler(c echo.Context) error {
	list, err := p.svc.GetInvoiceList(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    list,
	})
}

func (p InvoiceHandler) GetInvoiceDetailHandler(c echo.Context) error {
	id := c.Param("id")
	Invoice, err := p.svc.GetInvoiceDetail(context.Background(), id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    Invoice,
	})
}
