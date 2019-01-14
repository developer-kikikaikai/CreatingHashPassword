package controller

import (
	"errors"
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
	"net/http"
	"../db"
	"fmt"
)

type HandlePassphraseInfo struct {
}

type PartOfPassphraseInfo struct {
	Title string `json:"title" form:"title" query:"title"`
	Algorithm string `json:"algorithm" form:"algorithm" query:"algorithm"`
	Seed string `json:"seed" form:"seed" query:"seed"`
}

func (this *HandlePassphraseInfo) GetHandlerFunc(method Method) AuthHandlerFunc {
	switch method {
	case METHOD_GET: return this.get
	case METHOD_PUT: return this.put
	case METHOD_DELETE: return this.deleteMethod
	default: return nil
	}
}

func (this *HandlePassphraseInfo) get(c echo.Context, r *auth.AuthenticatedRequest) error {
	var res []PartOfPassphraseInfo
	//get request
	passphrases, err := db.GetAllPassphrase(r.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res)
		return errors.New("Failed to get PassphraseInfo from DB")
	}

	//create response data
	var part PartOfPassphraseInfo
	for i := 0; i < len(passphrases); i++ {
		part.Title = passphrases[i].Title
		part.Algorithm = passphrases[i].Algorithm
		part.Seed = passphrases[i].Seed
		res = append(res, part)
	}
	return c.JSON(http.StatusOK, res)
}

func (this *HandlePassphraseInfo) put(c echo.Context, r *auth.AuthenticatedRequest) error {
	//parse request body
	req_body := new(PartOfPassphraseInfo)
	if err := c.Bind(req_body); err != nil {
		fmt.Printf("HandlePassphraseInfo.put Failed to read body data:%s\n", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	//update/insert DB information
	dbculumn := db.PassphraseInfo{r.Username, req_body.Title, req_body.Algorithm, req_body.Seed}
	if result := db.SetPassphraseInfo(dbculumn); !result {
		fmt.Printf("HandlePassphraseInfo.put Failed to insert\n")
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (this *HandlePassphraseInfo) deleteMethod(c echo.Context, r *auth.AuthenticatedRequest) error {
	//parse request body
	//var req_body []string = make([]string, 0)
	var req_body []string
	if err := c.Bind(&req_body); err != nil {
		fmt.Printf("HandlePassphraseInfo.delete Failed to read body data:%s\n", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	//delete DB information
	for i:=0; i<len(req_body) ; i++ {
		db.DeletePassphraseInfo(r.Username, req_body[i])
	}

	return c.NoContent(http.StatusOK)
}

//always need authorization
func (this *HandlePassphraseInfo) DoesNeedAuthenticate(method Method) bool {
	return true
}
