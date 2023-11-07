package handler

import (
	"context"
	"net/http"

	"github.com/automotechnologies/doitpay/sdk/golang/sample/internal/virtualaccount/service"
	"github.com/labstack/echo/v4"
)

type VirtualAccountHandler struct {
	svc service.VirtualAccountServiceMethod
}

func NewVirtualAccountHandler(svc service.VirtualAccountServiceMethod) VirtualAccountHandler {
	return VirtualAccountHandler{
		svc: svc,
	}
}

func (p VirtualAccountHandler) GetVirtualAccountListHandler(c echo.Context) error {
	list, err := p.svc.GetVirtualAccountList(context.Background())

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

type RequestData struct {
	Suffix string `json:"suffix"`
	Amount int    `json:"amount"`
}

func (p VirtualAccountHandler) CreateVirtualAccountListHandler(c echo.Context) error {
	var requestData RequestData
	if err := c.Bind(&requestData); err != nil {
		return err
	}

	if len(requestData.Suffix) > 6 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "suffix length cannot more than 6",
		})
	}

	if requestData.Amount < 10000 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "amount cannot less than 10000",
		})
	}

	list, err := p.svc.CreateVirtualAccount(context.Background(), service.CreateVirtualAccountRequest{
		Suffix: requestData.Suffix,
		Amount: float32(requestData.Amount),
	})

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
