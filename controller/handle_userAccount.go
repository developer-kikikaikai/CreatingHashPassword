package controller

import (
	"errors"
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
)

type HandleUserAccount struct {
}

func (this *HandleUserAccount) GetHandlerFunc(method Method) AuthHandlerFunc {
	switch method {
	case METHOD_POST: return this.post
	case METHOD_PUT: return this.put
	case METHOD_DELETE: return this.delete
	default: return nil
	}
}

func (this *HandleUserAccount) post(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandleUserAccount.post")
}

func (this *HandleUserAccount) put(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandleUserAccount.put")
}

func (this *HandleUserAccount) delete(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandleUserAccount.delete")
}

func (this *HandleUserAccount) DoesNeedAuthenticate(method Method) bool {
	return true
}
