package controller

import (
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
	"../hash"
	"net/http"
	//"fmt"
)

type HandlePassphrase struct {
	Title string `json:"title" form:"title" query:"title"`
	Keyphrase string `json:"keyphrase" form:"keyphrase" query:"keyphrase"`
	Algorithm string `json:"algorithm" form:"algorithm" query:"algorithm"`
	Seed string `json:"seed" form:"seed" query:"seed"`
}

type responseBody struct {
	result string `json:"result" from:"result" query:"result"`
}

func (this *HandlePassphrase) GetHandlerFunc(method Method) AuthHandlerFunc {
	if method == METHOD_POST {
		return this.post
	} else {
		return nil
	}
}

func (this *HandlePassphrase) post(c echo.Context, r *auth.AuthenticatedRequest) error {
	res := new(responseBody)
	//parse json format
	req_body := new(HandlePassphrase)
	if err := c.Bind(req_body); err != nil {
		res.result = "HandlePassphrase.post Bind error"
		return c.JSON(http.StatusInternalServerError, res)
	}

	//create result
	res.result = hash.HashSum(req_body.Algorithm, req_body.Title + req_body.Keyphrase + req_body.Seed)
	if res.result != "" {
		return c.JSON(http.StatusOK, res)
	} else {
		return c.JSON(http.StatusBadRequest, res)
	}
}

func (this *HandlePassphrase) DoesNeedAuthenticate(method Method) bool {
	return false
}
