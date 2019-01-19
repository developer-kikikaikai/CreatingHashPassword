package controller

import (
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
	"net/http"
)

type Method int
const (
	METHOD_GET Method = iota
	METHOD_POST
	METHOD_PUT
	METHOD_DELETE
)

type AuthHandlerFunc func(c echo.Context, r *auth.AuthenticatedRequest) error

var method_names map[Method]string = map[Method]string{METHOD_GET:"Get", METHOD_POST:"Post", METHOD_PUT:"Put", METHOD_DELETE:"Delete"}

type MsgHandlerIf interface {
	Get(c echo.Context, r *auth.AuthenticatedRequest) error
	Post(c echo.Context, r *auth.AuthenticatedRequest) error
	Put(c echo.Context, r *auth.AuthenticatedRequest) error
	Delete(c echo.Context, r *auth.AuthenticatedRequest) error
	DoesNeedAuthenticate(method Method) bool
}

type MsgHandler struct {
}
func (this MsgHandler) Get(c echo.Context, r *auth.AuthenticatedRequest) error {return c.NoContent(http.StatusMethodNotAllowed)}
func (this MsgHandler) Put(c echo.Context, r *auth.AuthenticatedRequest) error {return c.NoContent(http.StatusMethodNotAllowed)}
func (this MsgHandler) Post(c echo.Context, r *auth.AuthenticatedRequest) error {return c.NoContent(http.StatusMethodNotAllowed)}
func (this MsgHandler) Delete(c echo.Context, r *auth.AuthenticatedRequest) error {return c.NoContent(http.StatusMethodNotAllowed)}

//return: map[uri]MsgHandler
func MsgHandlerFactory() map[string]MsgHandlerIf {
	var results map[string]MsgHandlerIf = map[string]MsgHandlerIf{"passphrase":&HandlePassphrase{}, "passphraseInfo":&HandlePassphraseInfo{}, "userAccount":&HandleUserAccount{}, "logout":&HandleLogout{}}
	return results
}

func GetEchoHandler(instance MsgHandlerIf, handler_fnc AuthHandlerFunc, method Method) echo.HandlerFunc{
	if ok := instance.DoesNeedAuthenticate(method); ok {
		return DigestAuthenticate(handler_fnc)
	} else {
		return NoAuthenticate(handler_fnc)
	}
}
