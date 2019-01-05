package controller

import (
	"html/template"
	"io"
	"net/http"
	"../hash"
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
)

var templates map[string]*template.Template
type Template struct {
}

// store template file in templates map
func LoadTemplates() {
	templates = make(map[string]*template.Template)
	templates["indexs"] = template.Must(
	    template.ParseFiles("templates/index.html.tpl"))
}

func ExecTemplates(w io.Writer, name string, data interface{}) error {
    return templates[name].ExecuteTemplate(w, "index.html.tpl", data)
}

func HandleIndex(c echo.Context, r *auth.AuthenticatedRequest) error{
	var data struct {
		Algorithms  []string
	}
    data.Algorithms = hash.AlgorithmList()
    return c.Render(http.StatusOK, "indexs", data)
}


