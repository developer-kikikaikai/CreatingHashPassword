package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	auth "github.com/abbot/go-http-auth"
	"github.com/developer-kikikaikai/CreatingHashPassword/db"
	"github.com/labstack/echo"
)

type HandlePassphraseInfo struct {
	MsgHandler
}

type PartOfPassphraseInfo struct {
	Title         string `json:"title" form:"title" query:"title"`
	Algorithm     string `json:"algorithm" form:"algorithm" query:"algorithm"`
	Seed          string `json:"seed" form:"seed" query:"seed"`
	Length        int    `json:"length" form:"length" query:"length"`
	DisableSymbol bool   `json:"disable_symbol" form:"disable_symbol" query:"disable_symbol"`
}

func (this *HandlePassphraseInfo) Get(c echo.Context, r *auth.AuthenticatedRequest) error {
	var res []PartOfPassphraseInfo
	//get request
	passphrases, err := db.GetAllPassphrase(r.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return errors.New("Failed to get PassphraseInfo from DB")
	}

	//create response data
	var part PartOfPassphraseInfo
	var _ error
	for i := 0; i < len(passphrases); i++ {
		part.Title = passphrases[i].Title
		part.Algorithm = passphrases[i].Algorithm
		part.Seed = passphrases[i].Seed
		part.Length, _ = strconv.Atoi(passphrases[i].Length)
		part.DisableSymbol, _ = strconv.ParseBool(passphrases[i].DisableSymbol)
		res = append(res, part)
	}
	return c.JSON(http.StatusOK, res)
}

func (this *HandlePassphraseInfo) Put(c echo.Context, r *auth.AuthenticatedRequest) error {
	//parse request body
	req_body := new(PartOfPassphraseInfo)
	if err := c.Bind(req_body); err != nil {
		fmt.Printf("HandlePassphraseInfo.put Failed to read body data:%s\n", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	//update/insert DB information
	var _ error
	dbcolumn := db.PassphraseInfo{r.Username, req_body.Title, req_body.Algorithm, req_body.Seed, strconv.Itoa(req_body.Length), strconv.FormatBool(req_body.DisableSymbol)}
	if result := db.SetPassphraseInfo(dbcolumn); !result {
		fmt.Printf("HandlePassphraseInfo.put Failed to insert\n")
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (this *HandlePassphraseInfo) Delete(c echo.Context, r *auth.AuthenticatedRequest) error {
	//parse request body
	//var req_body []string = make([]string, 0)
	var req_body []string
	if err := c.Bind(&req_body); err != nil {
		fmt.Printf("HandlePassphraseInfo.delete Failed to read body data:%s\n", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	//delete DB information
	for i := 0; i < len(req_body); i++ {
		db.DeletePassphraseInfo(r.Username, req_body[i])
	}

	return c.NoContent(http.StatusOK)
}

//always need authorization
func (this *HandlePassphraseInfo) DoesNeedAuthenticate(method Method) bool {
	if method == METHOD_POST {
		return false
	} else {
		return true
	}
}
