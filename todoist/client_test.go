package todoist_test

import (
	"testing"

	"github.com/salteron/todoist_tasks_scheduler/todoist"
)

func TestGetTasks(t *testing.T) {
	srv := HelperStartHTTPServer(`[{"ID": "1"}]`)
	defer srv.Close()

	tasks, _ := todoist.NewClient(srv.URL).GetTasks()

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].ID != "1" {
		t.Errorf("Expected task ID 1, got %s", tasks[0].ID)
	}
}

func TestGetProjects(t *testing.T) {
	srv := HelperStartHTTPServer(`[{"ID": "1"}]`)
	defer srv.Close()

	projects, _ := todoist.NewClient(srv.URL).GetProjects()

	if len(projects) != 1 {
		t.Errorf("Expected 1 project, got %d", len(projects))
	}

	if projects[0].ID != "1" {
		t.Errorf("Expected project ID 1, got %s", projects[0].ID)
	}
}
