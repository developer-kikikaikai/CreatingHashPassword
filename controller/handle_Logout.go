package controller

import (
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
	"net/http"
)

type HandleLogout struct {
}

func (this *HandleLogout) GetHandlerFunc(method Method) AuthHandlerFunc {
	switch method {
	case METHOD_GET: return this.get
	default: return nil
	}
}

func (this *HandleLogout) get(c echo.Context, r *auth.AuthenticatedRequest) error {
	return c.NoContent(http.StatusUnauthorized)
}

func (this *HandleLogout) DoesNeedAuthenticate(method Method) bool {
	return true
}
