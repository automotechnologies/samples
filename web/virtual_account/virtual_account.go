package transaction

import (
	"path"
	"text/template"

	"github.com/labstack/echo/v4"
)

var (
	viewPath = "web/virtual_account/views"
)

func VirtualAccountpageHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles(path.Join(viewPath, "list.html")))
	return tmpl.Execute(c.Response().Writer, nil)
}

func VirtualAccountpageDetailHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles(path.Join(viewPath, "detail.html")))
	return tmpl.Execute(c.Response().Writer, nil)
}
