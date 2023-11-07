package simulate

import (
	"path"
	"text/template"

	"github.com/labstack/echo/v4"
)

var (
	viewPath = "web/simulate/views"
)

func SimulatepageHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles(path.Join(viewPath, "simulate.html")))
	return tmpl.Execute(c.Response().Writer, nil)
}
