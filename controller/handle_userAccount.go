package controller

import (
	"fmt"
	"net/http"

	auth "github.com/abbot/go-http-auth"
	"github.com/developer-kikikaikai/CreatingHashPassword/db"
	"github.com/labstack/echo"
)

type HandleUserAccount struct {
	MsgHandler
	Username   string
	Passphrase string
}

func (this *HandleUserAccount) Post(c echo.Context, r *auth.AuthenticatedRequest) error {
	//parse request body
	req_body := new(HandleUserAccount)
	if err := c.Bind(req_body); err != nil {
		fmt.Printf("HandleUserAccount.post Failed to read body data:%s\n", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	//update/insert DB information
	if _, err := db.GetAccount(req_body.Username); err == nil {
		fmt.Printf("HandleUserAccount.post requested user is exist\n")
		return c.NoContent(http.StatusBadRequest)
	}

	dbcolumn := db.Account{req_body.Username, GetInsertedPassphrase(req_body.Username, req_body.Passphrase)}
	if !db.SetAccount(dbcolumn) {
		fmt.Printf("HandleUserAccount.post Failed to insert\n")
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func (this *HandleUserAccount) Put(c echo.Context, r *auth.AuthenticatedRequest) error {
	//parse request body
	req_body := new(HandleUserAccount)
	if err := c.Bind(req_body); err != nil {
		fmt.Printf("HandleUserAccount.put Failed to read body data:%s\n", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	//update/insert DB information
	dbcolumn := db.Account{r.Username, GetInsertedPassphrase(r.Username, req_body.Passphrase)}
	if !db.SetAccount(dbcolumn) {
		fmt.Printf("HandleUserAccount.put Failed to insert\n")
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func (this *HandleUserAccount) Delete(c echo.Context, r *auth.AuthenticatedRequest) error {
	db.DeleteAccount(r.Username)
	return c.NoContent(http.StatusOK)
}

func (this *HandleUserAccount) DoesNeedAuthenticate(method Method) bool {
	if method == METHOD_POST {
		return false
	} else {
		return true
	}
}
