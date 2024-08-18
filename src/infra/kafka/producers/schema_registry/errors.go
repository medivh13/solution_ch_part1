package kafka_schema_registry

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.ErrorCode, e.Message)
}

func newError(resp *http.Response) *Error {
	err := &Error{}
	parsingErr := json.NewDecoder(resp.Body).Decode(err)
	if parsingErr != nil {
		return &Error{resp.StatusCode, "Unrecognized error found"}
	}
	return err
}
