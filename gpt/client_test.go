package gpt_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/salteron/todoist_tasks_scheduler/gpt"
)

func TestComplete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"choices":[{"message":{"content": "completion"}}]}`)
	}))
	defer srv.Close()

	completion, err := gpt.NewClientWithURL("123", srv.URL).Complete("complete-me")

	if err != nil {
		t.Fatalf("got error: %s", err)
	}

	if completion != "completion" {
		t.Fatalf("got: %s, want: %s", completion, "completion")
	}
}
