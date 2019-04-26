package controller

import (
	"net/http"

	auth "github.com/abbot/go-http-auth"
	"github.com/developer-kikikaikai/CreatingHashPassword/hash"
	"github.com/labstack/echo"
	//"fmt"
)

type HandlePassphrase struct {
	MsgHandler
	//It's ok not to use `json:"title" form:"title" query:"title"` information if member name is same
	Title         string
	Keyphrase     string
	Algorithm     string
	Seed          string
	Length        int
	DisableSymbol bool `json:"disable_symbol" form:"disable_symbol" query:"disable_symbol"`
}

type PassphraseResponseBody struct {
	Result string `json:"result" from:"result" query:"result"`
}

func (this *HandlePassphrase) Post(c echo.Context, r *auth.AuthenticatedRequest) error {
	res := new(PassphraseResponseBody)
	//parse json format
	req_body := new(HandlePassphrase)
	if err := c.Bind(req_body); err != nil {
		res.Result = "HandlePassphrase.post Bind error"
		return c.JSON(http.StatusInternalServerError, res)
	}

	//create result
	res.Result = hash.HashSum(req_body.Algorithm, req_body.Title+req_body.Keyphrase+req_body.Seed)
	if res.Result != "" {
		scaleDownDenominator := len(res.Result) / req_body.Length
		res.Result = hash.CompressHash(res.Result, scaleDownDenominator, !req_body.DisableSymbol)
		return c.JSON(http.StatusOK, res)
	} else {
		return c.JSON(http.StatusBadRequest, res)
	}
}

func (this *HandlePassphrase) DoesNeedAuthenticate(method Method) bool {
	return false
}
