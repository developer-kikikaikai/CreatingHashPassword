package controller

import (
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
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

type MsgHandler interface {
	GetHandlerFunc(method Method) AuthHandlerFunc
	DoesNeedAuthenticate(method Method) bool
}

//return: map[uri]MsgHandler
func MsgHandlerFactory() map[string]MsgHandler {
	var results map[string]MsgHandler = map[string]MsgHandler{"passphrase":&HandlePassphrase{}, "passphraseInfo":&HandlePassphraseInfo{}, "userAccount":&HandleUserAccount{}}
	return results
}

func GetEchoHandler(instance MsgHandler, method Method) echo.HandlerFunc{

	handler_fnc := instance.GetHandlerFunc(method)
	if handler_fnc == nil {
		return nil
	}

	if ok := instance.DoesNeedAuthenticate(method); ok {
		return DigestAuthenticate(handler_fnc)
	} else {
		return NoAuthenticate(handler_fnc)
	}
}
