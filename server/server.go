package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/automotechnologies/doitpay-go"
	invoice "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/invoice/handler"
	invoiceSVC "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/invoice/service"
	simulate "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/simulate/handler"
	simulateSVC "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/simulate/service"
	va "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/virtualaccount/handler"
	vaSVC "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/virtualaccount/service"
	webhook "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/webhook/handler"
	webhookSVC "github.com/automotechnologies/doitpay/sdk/golang/sample/internal/webhook/service"
	homeWeb "github.com/automotechnologies/doitpay/sdk/golang/sample/web/dashboard"
	simWeb "github.com/automotechnologies/doitpay/sdk/golang/sample/web/simulate"
	trxWeb "github.com/automotechnologies/doitpay/sdk/golang/sample/web/transaction"
	vaWeb "github.com/automotechnologies/doitpay/sdk/golang/sample/web/virtual_account"
	webhookWeb "github.com/automotechnologies/doitpay/sdk/golang/sample/web/webhook"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Server is list configuration to run Server
type Server struct {
	apikey            string
	checkoutLink      string
	client            *doitpay.APIClient
	invoiceSVC        invoiceSVC.InvoiceServiceMethod
	invoiceHdl        invoice.InvoiceHandler
	virtualaccountSVC vaSVC.VirtualAccountServiceMethod
	virtualaccountHdl va.VirtualAccountHandler
	simulateSVC       simulateSVC.SimulateServiceMethod
	simulateHdl       simulate.SimulateHandler
	webhookSVC        webhookSVC.WebhookServiceMethod
	webhookHdl        webhook.WebhookHandler
	handler           *echo.Echo
	port              string
}

// NewServer is func to create server with all configuration
func NewServer() (*Server, error) {
	s := &Server{}

	// ======== Init Dependencies Related ========
	// Load Env File
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		return s, err
	}

	// Load Env
	{
		s.apikey = os.Getenv("API_KEY")
		if len(s.apikey) <= 0 || s.apikey == "<API_KEY>" {
			fmt.Print("[Got Error]- Invalid API_KEY")
			return s, fmt.Errorf("[Got Error]-Invalid API_KEY")
		}

		s.checkoutLink = os.Getenv("CHECKOUT_LINK")
		if len(s.checkoutLink) <= 0 || s.apikey == "<CHECKOUT_LINK>" {
			fmt.Print("[Got Error]- Invalid CHECKOUT_LINK")
			return s, fmt.Errorf("[Got Error]-Invalid CHECKOUT_LINK")
		}
	}

	// Init client
	{
		s.client = doitpay.NewClient(s.apikey)
		if s.client == nil {
			fmt.Print("[Got Error]- Invalid Init Doitpay Client")
			return s, fmt.Errorf("[Got Error]-Invalid Init Doitpay Client")
		}
	}

	// Init Internal Service
	{
		// invoice service
		s.invoiceSVC = invoiceSVC.NewInvoiceService(s.client, s.checkoutLink)

		// virtual account service
		s.virtualaccountSVC = vaSVC.NewVirtualAccountService(s.client)

		// simulate service
		s.simulateSVC = simulateSVC.NewSimulateService(s.client)

		// webhook service
		s.webhookSVC = webhookSVC.NewWebhookService()
	}

	// Init Handler
	{
		// invoice handler
		s.invoiceHdl = invoice.NewInvoiceHandler(s.invoiceSVC)

		// virtual account handler
		s.virtualaccountHdl = va.NewVirtualAccountHandler(vaSVC.NewVirtualAccountService(s.client))

		// simulate handler
		s.simulateHdl = simulate.NewSimulateHandler(s.simulateSVC)

		// webhook handler
		s.webhookHdl = webhook.NewWebhookHandler(s.webhookSVC)
	}

	// Init Router
	{
		e := echo.New()

		// Web Page handler
		e.GET("/", homeWeb.HomepageHandler)
		e.GET("/transaction", trxWeb.TransactionpageHandler)
		e.GET("/transaction/:id", trxWeb.TransactionpageDetailHandler)
		e.GET("/virtual-account", vaWeb.VirtualAccountpageHandler)
		e.GET("/virtual-account/:id", vaWeb.VirtualAccountpageDetailHandler)
		e.GET("/simulate", simWeb.SimulatepageHandler)
		e.GET("/webhook", webhookWeb.WebhookpageHandler)

		// API handler
		e.POST("/api/v1/purchase", s.invoiceHdl.CreateInvoiceHandler)
		e.GET("/api/v1/transaction", s.invoiceHdl.GetInvoiceListHandler)
		e.GET("/api/v1/transaction/:id", s.invoiceHdl.GetInvoiceDetailHandler)
		e.GET("/api/v1/virtual-account", s.virtualaccountHdl.GetVirtualAccountListHandler)
		e.POST("/api/v1/virtual-account", s.virtualaccountHdl.CreateVirtualAccountListHandler)
		e.POST("/api/v1/simulate", s.simulateHdl.DoSimulateHandler)
		e.POST("/api/v1/webhook", s.webhookHdl.WebhookHandler)
		e.GET("/api/v1/webhook", s.webhookHdl.GetWebhookHandler)

		s.handler = e
		s.port = ":8000"
	}
	return s, nil
}

func (s *Server) Start() int {
	go func() {
		if err := s.handler.Start(s.port); err != nil {
			fmt.Println(err)
		}
	}()

	// Wait for a signal to shut down the application
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a context with a timeout to allow the server to cleanly shut down
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	s.handler.Shutdown(ctx)
	return 0
}

// Run is func to create server and invoke Start()
func Run() int {
	s, err := NewServer()
	if err != nil {
		return 1
	}

	return s.Start()
}
