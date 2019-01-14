package controller

import (
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
	"net/http"
	"fmt"
	"../db"
)

type HandleUserAccount struct {
	Username string
	Passphrase string
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
	//parse request body
	req_body := new(HandleUserAccount)
	if err := c.Bind(req_body); err != nil {
		fmt.Printf("HandleUserAccount.post Failed to read body data:%s\n", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	//update/insert DB information
	if _,err := db.GetAccount(req_body.Username); err == nil {
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

func (this *HandleUserAccount) put(c echo.Context, r *auth.AuthenticatedRequest) error {
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

func (this *HandleUserAccount) delete(c echo.Context, r *auth.AuthenticatedRequest) error {
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
