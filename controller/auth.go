package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/abbot/go-http-auth"
	"../db"
	"fmt"
)

func authenticate(user, realm string) string {
	fmt.Printf("user:%s\n", user);
	account, err := db.GetAccount(user)
	if err == nil {
		fmt.Printf("password:%s\n", account.Passphrase);
		return account.Passphrase
	}
	return ""
}

//define authenticator to use all
var authenticator * auth.DigestAuth = auth.NewDigestAuthenticator("example.com", authenticate)

//modify auth.Wrap function to use echo format function
func DigestAuthenticate(input func(c echo.Context, r *auth.AuthenticatedRequest) error) echo.HandlerFunc{
	return func(c echo.Context) error {
		r := c.Request()
		w := c.Response().Writer
		if username, authinfo := authenticator.CheckAuth(r); username == "" {
	fmt.Printf("if case\n");
			authenticator.RequireAuth(w, r)
			return echo.NewHTTPError(http.StatusUnauthorized, "Please write collect username and password")
		} else {
	fmt.Printf("else case\n");
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

