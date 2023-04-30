package prompts_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/salteron/todoist_tasks_scheduler/prompts"
	"github.com/salteron/todoist_tasks_scheduler/todoist"
)

func TestGeneratePrompt(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `[{"ID": "test-project-ID"}]`)
	}))
	defer srv.Close()

	todoist.DefaultClient = todoist.NewClient(srv.URL)

	task := todoist.Task{
		ProjectID: "test-project-ID",
		Content:   "test-content",
		Duration:  5 * time.Minute,
		Due: todoist.Due{
			Datetime: todoist.TodoistTime{
				Time: time.Date(2020, 1, 1, 13, 30, 0, 0, time.UTC),
			},
		},
	}

	got, err := prompts.GeneratePrompt(prompts.Params{
		DayStart:        "08:00",
		DayEnd:          "20:00",
		PromptPath:      "../prompts/templates/test.txt",
		PromptExtension: "extension",
		Tasks:           []todoist.Task{task},
	})

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	want := `08:00, 20:00, [{"duration":5,"project":"","start time":"13:30","title":"test-content"}], extension`

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
