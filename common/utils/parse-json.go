package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReqParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func ResParseJSON(w *http.Response, payload any) error {
	if w.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(w.Body).Decode(payload)
}
