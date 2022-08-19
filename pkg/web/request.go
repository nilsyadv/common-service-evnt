package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	errs "github.com/nilsyadv/common-service-evnt/pkg/error"
)

func Request(out interface{}, req *http.Request) error {
	if req == nil {
		return errs.NewHTTPError("Request is invalid", http.StatusBadRequest)
	}
	reqbd := req.Body
	byt, err := ioutil.ReadAll(reqbd)
	if err != nil {
		return errs.NewInvalidRequestPayloadError(errs.ErrorCodeInternalError)
	}
	if len(byt) == 0 {
		return errs.NewHTTPError("Request Body is Empty", http.StatusBadRequest)
	}
	err = json.Unmarshal(byt, out)
	if err != nil {
		return errs.NewInvalidRequestPayloadError(errs.ErrorCodeInternalError)
	}
	return nil
}
