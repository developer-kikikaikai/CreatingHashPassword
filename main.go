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
    e := echo.New()

    // set renderer to use html template file
    t := &Template{}
    e.Renderer = t

    // setup middleware(Todo, search spec of middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

	// regist handler
	group := e.Group("/api/")
	msg_handler := controller.MsgHandlerFactory()
	for uri, instance := range msg_handler {
		//Set method GET
		if handler_fnc := controller.GetEchoHandler(instance, controller.METHOD_GET); handler_fnc != nil {
			//store URI
			group.GET(uri, handler_fnc)
		}

		if handler_fnc := controller.GetEchoHandler(instance, controller.METHOD_POST); handler_fnc != nil {
			//store URI
			group.POST(uri, handler_fnc)
		}

		if handler_fnc := controller.GetEchoHandler(instance, controller.METHOD_PUT); handler_fnc != nil {
			//store URI
			group.PUT(uri, handler_fnc)
		}

		if handler_fnc := controller.GetEchoHandler(instance, controller.METHOD_DELETE); handler_fnc != nil {
			//store URI
			group.DELETE(uri, handler_fnc)
		}
	}
	e.GET("/", controller.NoAuthenticate(controller.HandleIndex))

    // サーバーを開始
    e.Logger.Fatal(e.Start(":3000"))
}

// 初期化を行います。
func init() {
	controller.LoadTemplates()
}
