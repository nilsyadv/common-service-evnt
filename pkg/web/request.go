package web

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// func Request(out interface{}, req *http.Request) error {
// 	if req == nil {
// 		return errs.NewHTTPError("Request is invalid", http.StatusBadRequest)
// 	}
// 	reqbd := req.Body
// 	byt, err := ioutil.ReadAll(reqbd)
// 	if err != nil {
// 		return errs.NewInvalidRequestPayloadError(errs.ErrorCodeInternalError)
// 	}
// 	if len(byt) == 0 {
// 		return errs.NewHTTPError("Request Body is Empty", http.StatusBadRequest)
// 	}
// 	err = json.Unmarshal(byt, out)
// 	if err != nil {
// 		return errs.NewInvalidRequestPayloadError(errs.ErrorCodeInternalError)
// 	}
// 	return nil
// }

func Request(out interface{}, req *http.Request) error {
	if req == nil {
		return errors.New("request is invalid")
	}
	reqbd := req.Body
	byt, err := ioutil.ReadAll(reqbd)
	if err != nil {
		return errors.New("unable to read request")
	}
	if len(byt) == 0 {
		return errors.New("request body is empty")
	}
	err = json.Unmarshal(byt, out)
	if err != nil {
		return errors.New("failed to unmarshal the request")
	}
	return nil
}
