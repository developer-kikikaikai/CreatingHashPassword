package controller

import (
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
	"net/http"
)

type HandleLogout struct {
	MsgHandler
}

func (this *HandleLogout) Get(c echo.Context, r *auth.AuthenticatedRequest) error {
	return c.NoContent(http.StatusUnauthorized)
}

func (this *HandleLogout) DoesNeedAuthenticate(method Method) bool {
	return true
}
