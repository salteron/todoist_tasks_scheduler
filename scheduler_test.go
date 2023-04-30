package scheduler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	scheduler "github.com/salteron/todoist_tasks_scheduler"
	"github.com/salteron/todoist_tasks_scheduler/gpt"
	"github.com/salteron/todoist_tasks_scheduler/todoist"
)

func gptHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"choices":[{"message":{"content": "completion"}}]}`)
	}))
}

func todoistHTTPServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"ID": "test-project-ID", "Name": "test-project-Name"}]`)
	})

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"ID": "1"}]`)
	})

	return httptest.NewServer(mux)
}

func TestGenerateSchedule(t *testing.T) {
	todoistSrv := todoistHTTPServer()
	defer todoistSrv.Close()

	gptSrv := gptHTTPServer()
	defer gptSrv.Close()

	todoist.DefaultClient = todoist.NewClient(todoistSrv.URL)
	gpt.DefaultClient = gpt.NewClientWithURL("apiKey", gptSrv.URL)

	got, err := scheduler.GenerateSchedule(scheduler.Params{
		DayStart:        "08:00",
		DayEnd:          "00:00",
		PromptPath:      "prompts/templates/default.txt",
		PromptExtension: "extra",
	})

	if err != nil {
		t.Fatalf("got error: %s", err)
	}

	if got != "completion" {
		t.Fatalf("got %s, want %s", got, "completion")
	}
}
