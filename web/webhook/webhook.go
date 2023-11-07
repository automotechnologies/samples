package transaction

import (
	"path"
	"text/template"

	"github.com/labstack/echo/v4"
)

var (
	viewPath = "web/webhook/views"
)

func WebhookpageHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles(path.Join(viewPath, "list.html")))
	return tmpl.Execute(c.Response().Writer, nil)
}
