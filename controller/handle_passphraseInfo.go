package controller

import (
	"errors"
	"github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
)

type HandlePassphraseInfo struct {
}

func (this *HandlePassphraseInfo) GetHandlerFunc(method Method) AuthHandlerFunc {
	switch method {
	case METHOD_GET: return this.get
	case METHOD_POST: return this.post
	case METHOD_PUT: return this.put
	case METHOD_DELETE: return this.delete
	default: return nil
	}
}

func (this *HandlePassphraseInfo) get(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandlePassphraseInfo.get")
}

func (this *HandlePassphraseInfo) post(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandlePassphraseInfo.post")
}

func (this *HandlePassphraseInfo) put(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandlePassphraseInfo.put")
}

func (this *HandlePassphraseInfo) delete(c echo.Context, r *auth.AuthenticatedRequest) error {
	return errors.New("HandlePassphraseInfo.delete")
}

func (this *HandlePassphraseInfo) DoesNeedAuthenticate(method Method) bool {
	return true
}
