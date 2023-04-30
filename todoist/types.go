package todoist

import (
	"encoding/json"
	"fmt"
	"time"
)

// Represent a Todoist task
type Task struct {
	CreatorID    string   `json:"creator_id"`
	CreatedAt    string   `json:"created_at"`
	AssigneeID   string   `json:"assignee_id"`
	AssignerID   string   `json:"assigner_id"`
	CommentCount int64    `json:"comment_count"`
	IsCompleted  bool     `json:"is_completed"`
	Content      string   `json:"content"`
	Description  string   `json:"description"`
	Due          Due      `json:"due"`
	ID           string   `json:"id"`
	Labels       []string `json:"labels"`
	Order        int64    `json:"order"`
	Priority     int64    `json:"priority"`
	ProjectID    string   `json:"project_id"`
	SectionID    string   `json:"section_id"`
	ParentID     string   `json:"parent_id"`
	URL          string   `json:"url"`
	Duration     time.Duration
}

func (t Task) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"title":    t.Content,
		"duration": t.Duration.Minutes(),
	}

	if t.ProjectID != "" {
		if project, err := ProjectsRepo.GetById(t.ProjectID); err != nil {
			return nil, err
		} else {
			data["project"] = project.Name
		}
	}

	if !t.Due.Datetime.Time.IsZero() {
		data["start time"] = t.Due.Datetime.Format("15:04")
	}
	return json.Marshal(data)
}

type Due struct {
	Date        TodoistDate `json:"date"`
	IsRecurring bool        `json:"is_recurring"`
	Datetime    TodoistTime `json:"datetime"`
	String      string      `json:"string"`
	Timezone    string      `json:"timezone"`
}

type TodoistDate struct {
	time.Time
}

func (t *TodoistDate) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type TodoistTime struct {
	time.Time
}

func (t *TodoistTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

// Represents a Todoist project
type Project struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CommentCount   int64  `json:"comment_count"`
	Order          int64  `json:"order"`
	Color          string `json:"color"`
	IsShared       bool   `json:"is_shared"`
	IsFavorite     bool   `json:"is_favorite"`
	ParentID       string `json:"parent_id"`
	IsInboxProject bool   `json:"is_inbox_project"`
	IsTeamInbox    bool   `json:"is_team_inbox"`
	ViewStyle      string `json:"view_style"`
	URL            string `json:"url"`
}

/*
Represents projects repository with caching functionality.
*/
var ProjectsRepo = projectsRepo{}

type projectsRepo struct {
	instances map[string]Project
	isLoaded  bool
}

/*
Returns a Project with the specified ID.
Returns error if a Project with such ID is not found.

Uses Client.GetProjects() to get all projects and caches them so that
subsequent calls don't make network calls.
*/
func (r *projectsRepo) GetById(projectId string) (*Project, error) {
	if !r.isLoaded {
		r.instances = make(map[string]Project)
		projects, err := GetProjects()

		if err != nil {
			return nil, err
		}

		for _, p := range projects {
			r.instances[p.ID] = p
		}

		r.isLoaded = true
	}

	if project, found := r.instances[projectId]; !found {
		return nil, fmt.Errorf("Project with ID=%s not found", projectId)
	} else {
		return &project, nil
	}
}
