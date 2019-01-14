package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/abbot/go-http-auth"
	"../db"
	_ "fmt"
)

func authenticate(user, realm string) string {
	account, err := db.GetAccount(user)
	if err == nil {
		//fmt.Printf("%v\n", account)
		return account.Passphrase
	}
	return ""
}

//define authenticator to use all
var authenticator * auth.DigestAuth = authwrapNewDigestAuthenticator()

func authwrapNewDigestAuthenticator() *auth.DigestAuth {
	res := auth.NewDigestAuthenticator("CreatingHashPassword", authenticate)
	return res
}

//modify auth.Wrap function to use echo format function
func DigestAuthenticate(input func(c echo.Context, r *auth.AuthenticatedRequest) error) echo.HandlerFunc{
	return func(c echo.Context) error {
		r := c.Request()
		w := c.Response().Writer
		if username, authinfo := authenticator.CheckAuth(r); username == "" {
			authenticator.RequireAuth(w, r)
			return echo.NewHTTPError(http.StatusUnauthorized, "Please write collect username and password")
		} else {
			ar := &auth.AuthenticatedRequest{Request: *r, Username: username}
			if authinfo != nil {
				w.Header().Set(authenticator.Headers.V().AuthInfo, *authinfo)
			}
			return input(c, ar)
		}
	}
}

func NoAuthenticate(input func(c echo.Context, r *auth.AuthenticatedRequest) error) echo.HandlerFunc{
	return func(c echo.Context) error {
		return input(c, nil);
	}
}

