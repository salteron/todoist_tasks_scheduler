package todoist_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/salteron/todoist_tasks_scheduler/todoist"
)

func TestMarshalJSON(t *testing.T) {
	srv := HelperStartHTTPServer(`[{"ID": "test-project-ID", "Name": "test-project-Name"}]`)
	defer srv.Close()

	todoist.DefaultClient = todoist.NewClient(srv.URL)

	task := todoist.Task{
		ProjectID: "test-project-ID",
		Content:   "test-content",
		Duration:  5 * time.Minute,
		Due: todoist.Due{
			Datetime: todoist.TodoistTime{time.Date(2020, 1, 1, 13, 30, 0, 0, time.UTC)},
		},
	}

	got, _ := json.Marshal(task)
	want := `{"duration":5,"project":"test-project-Name","start time":"13:30","title":"test-content"}`

	if want != string(got) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestProjectsRepoGetByID(t *testing.T) {
	srv := HelperStartHTTPServer(`[{"ID": "test-project-ID"}]`)
	defer srv.Close()

	todoist.DefaultClient = todoist.NewClient(srv.URL)

	project, _ := todoist.ProjectsRepo.GetById("test-project-ID")

	if project.ID != "test-project-ID" {
		t.Errorf("got %s, want %s", project.ID, "test-project-ID")
	}

	project, _ = todoist.ProjectsRepo.GetById("unknown-project-ID")

	if project != nil {
		t.Errorf("got %v, want nil", project)
	}
}
