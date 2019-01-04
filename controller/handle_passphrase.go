package controller

import (
	"errors"
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
)

type HandlePassphrase struct {
}

func (this *HandlePassphrase) GetHandlerFunc(method Method) AuthHandlerFunc {
	if method == METHOD_GET {
		return this.get
	} else {
		return nil
	}
}

func (this *HandlePassphrase) get(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandlePassphrase.get")
}

func (this *HandlePassphrase) DoesNeedAuthenticate(method Method) bool {
	return false
}
