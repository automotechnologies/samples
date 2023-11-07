package dashboard

import (
	"path"
	"text/template"

	"github.com/labstack/echo/v4"
)

var (
	dashboardPath = "web/dashboard/views"
)

func HomepageHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles(path.Join(dashboardPath, "homepage.html")))
	return tmpl.Execute(c.Response().Writer, nil)
}
