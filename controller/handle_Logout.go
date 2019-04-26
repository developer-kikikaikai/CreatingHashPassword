package controller

import (
	"net/http"

	auth "github.com/abbot/go-http-auth"
	"github.com/labstack/echo"
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
