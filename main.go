package main

import (
	"./controller"
    "io"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

type Template struct {
}

// Render はHTMLテンプレートにデータを埋め込んだ結果をWriterに書き込みます。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return controller.ExecTemplates(w, name, data)
}

func main() {
    // Echoのインスタンスを生成
    e := echo.New()

    // テンプレートを利用するためのRendererの設定
    t := &Template{}
    e.Renderer = t

    // ミドルウェアを設定
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // 静的ファイルのパスを設定
	e.Static("/public/css/", "./public/css")
	e.Static("/public/js/", "./public/js/")
	e.Static("/public/img/", "./public/img/")

	// 各ルーティングに対するハンドラを設定
	e.GET("/", controller.HandleIndex)
	e.GET("/api/user", controller.HandleUserInfo)
	e.GET("/api/password", controller.HandlePassword)

    // サーバーを開始
    e.Logger.Fatal(e.Start(":3000"))
}

// 初期化を行います。
func init() {
	controller.LoadTemplates()
}
