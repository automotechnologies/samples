package handler

import (
	"context"
	"net/http"

	"github.com/automotechnologies/doitpay/sdk/golang/sample/internal/simulate/service"
	"github.com/labstack/echo/v4"
)

type SimulateHandler struct {
	svc service.SimulateServiceMethod
}

func NewSimulateHandler(svc service.SimulateServiceMethod) SimulateHandler {
	return SimulateHandler{
		svc: svc,
	}
}

type RequestData struct {
	AccountNumber string `json:"account_number"`
}

func (p SimulateHandler) DoSimulateHandler(c echo.Context) error {
	var requestData RequestData
	if err := c.Bind(&requestData); err != nil {
		return err
	}

	if len(requestData.AccountNumber) <= 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "account number cannot be null",
		})
	}

	err := p.svc.DoSimulate(context.Background(), service.SimulateRequest{
		VirtualAccountNumber: requestData.AccountNumber,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}
