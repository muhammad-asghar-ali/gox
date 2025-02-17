package helpers

import (
	"io"
	"net/http"
)

func ReadBody(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	HandleError(err)

	return body
}
