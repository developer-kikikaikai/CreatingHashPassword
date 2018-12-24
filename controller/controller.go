package controller

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo"
)

var templates map[string]*template.Template
type Template struct {
}

// 各HTMLテンプレートに共通レイアウトを適用した結果を保存します（初期化時に実行）。
func LoadTemplates() {
	var baseTemplate = "html/layout.html"
	templates = make(map[string]*template.Template)
	templates["indexs"] = template.Must(
	    template.ParseFiles(baseTemplate, "html/hello.html"))
}

func ExecTemplates(w io.Writer, name string, data interface{}) error {
    return templates[name].ExecuteTemplate(w, "layout.html", data)
}

func HandleIndex(c echo.Context) error{
    return c.Render(http.StatusOK, "indexs", "World")
}

func HandleUserInfo(c echo.Context) error{
    return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
}

func HandlePassword(c echo.Context) error{
    return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
}
