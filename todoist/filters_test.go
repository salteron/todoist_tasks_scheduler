package todoist_test

import (
	"testing"
	"time"

	"github.com/salteron/todoist_tasks_scheduler/todoist"
)

func TestFilterTasksByDate(t *testing.T) {
	tasks := []todoist.Task{
		{
			Due: todoist.Due{
				Date: todoist.TodoistDate{
					Time: time.Date(2020, time.July, 13, 20, 32, 0, 0, time.UTC),
				},
			},
		},
		{
			Due: todoist.Due{
				Date: todoist.TodoistDate{
					Time: time.Date(2020, time.July, 14, 20, 32, 0, 0, time.UTC),
				},
			},
		},
		{
			Due: todoist.Due{
				Date: todoist.TodoistDate{
					Time: time.Date(2021, time.July, 13, 20, 32, 0, 0, time.UTC),
				},
			},
		},
	}

	filterByDate := time.Date(2020, time.July, 13, 0, 0, 0, 0, time.UTC)
	got := todoist.FilterTasksByDate(tasks, filterByDate)

	if len(got) != 1 {
		t.Fatalf(
			"FilterTasksByDate(%v) returned %d tasks, want 1",
			tasks,
			len(got),
		)
	}

	if got[0].Due.Date.Time.Year() != filterByDate.Year() &&
		got[0].Due.Date.Time.Day() != filterByDate.Day() {
		t.Fatalf(
			"FilterTasksByDate(%v) returned a task with invalid date %v",
			tasks,
			got[0].Due.Date.Time,
		)
	}
}

func TestFilterTasksWithDurations(t *testing.T) {
	tasks := []todoist.Task{
		{
			Labels: []string{"duration-15-minutes", "duration-30-minutes"},
		},
		{
			Labels: []string{"15-minutes"},
		},
		{
			Labels: []string{"duration-15"},
		},
		{
			Labels: []string{"duration-a-minutes", "duration-minutes"},
		},
		{
			Labels: []string{"duration-minutes", "duration-30-minutes"},
		},
	}

	got := todoist.FilterTasksWithDurations(tasks)

	if len(got) != 2 {
		t.Fatalf("FilterTasksWithDurations(%v) returned slice of length %v, want %v", tasks, len(got), 2)
	}

	if got[0].Duration != 15*time.Minute {
		t.Fatalf(
			"FilterTasksWithDurations(%v)[%d].Duration returned %v, want %v",
			tasks, 0, got[0].Duration, 15*time.Minute,
		)
	}

	if got[1].Duration != 30*time.Minute {
		t.Fatalf(
			"FilterTasksWithDurations(%v)[%d].Duration returned %v, want %v",
			tasks, 1, got[1].Duration, 30*time.Minute,
		)
	}
}
