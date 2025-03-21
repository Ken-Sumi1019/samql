package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

func Init() *Template {
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/**.html")),
	}
	return t
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
