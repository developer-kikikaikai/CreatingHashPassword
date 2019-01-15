package main

import (
	"./controller"
	"io"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"sync"
)

type Template struct {
}

// Render is in echo.Context input parameter
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return controller.ExecTemplates(w, name, data)
}

//request dump
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
  fmt.Printf("Request Body: %v\n", string(reqBody))
  fmt.Printf("Response Body: %v\n", string(resBody))
}

var e *echo.Echo
// exit signal handler
func SignalHandler(doneCh chan struct{}) {
	//set signal channel
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGTERM,
 		syscall.SIGINT)
	defer signal.Stop(sigCh)
	for {
		select {
			case <-sigCh:
				//send exit by using 
				//fmt.Printf("receive signal\n")
				e.Close()
			case <-doneCh:
				//fmt.Printf("receive doneCh\n")
				close(sigCh)
				return;
		}
	}
}

func main() {
	e = echo.New()

	//create wait group to sync signal handler
	wg := sync.WaitGroup{}
	doneCh := make(chan struct{}, 1)

	//signal handler
	wg.Add(1)
	go func() {
		defer wg.Done()
		SignalHandler(doneCh)
	}()

	// set renderer to use html template file
	t := &Template{}
	e.Renderer = t

	// setup middleware(Todo, search spec of middleware
	//Set Logger if you want to show log
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(bodyDumpHandler))
	//Set Recover if you want to respawn server
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
	controller.SetStatic(e)
	e.Start(":60080")
	fmt.Printf("Exit server\n")
	//send message to done channel
	doneCh <- struct{}{}
	wg.Wait()
	close(doneCh)
}

//load template
func init() {
	controller.LoadTemplates()
}
