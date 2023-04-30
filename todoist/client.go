// Package provides a Client for Todoist API.
package todoist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	defaultURL      = "https://api.todoist.com/rest/v2"
	getProjectsPath = "/projects"
	getTasksPath    = "/tasks"
)

// Represents a client for Todoist API.
type Client struct {
	url      string
	apiToken string
}

// Returns a new client for Todoist API
func NewClient(url string) Client {
	return Client{url: url, apiToken: os.Getenv("SCHEDULER_TODOIST_API_TOKEN")}
}

// A default client for Todoist API
var DefaultClient = NewClient(defaultURL)

// Returns all tasks.
func GetTasks() ([]Task, error) {
	return DefaultClient.GetTasks()
}

// Returns all projects.
func GetProjects() ([]Project, error) {
	return DefaultClient.GetProjects()
}

// Returns all tasks.
func (c Client) GetTasks() ([]Task, error) {
	var tasks []Task
	url, _ := url.JoinPath(c.url, getTasksPath)
	tasksJSON, err := c.doTodoistGetReq(url)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(tasksJSON, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// Returns all projects.
func (c Client) GetProjects() ([]Project, error) {
	var projects []Project
	url, _ := url.JoinPath(c.url, getProjectsPath)

	if projectsJSON, err := c.doTodoistGetReq(url); err != nil {
		return nil, err
	} else if err = json.Unmarshal(projectsJSON, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (c Client) doTodoistGetReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
