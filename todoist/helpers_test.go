package todoist_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func HelperStartHTTPServer(responseJSON string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, responseJSON)
	}))
}
