// Package provides the functionality for prompting the user for a schedule
// generation
package prompts

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/salteron/todoist_tasks_scheduler/todoist"
)

// GeneratePrompt returns a prompt that asks to generate a schedule for Todoist
// tasks.
func GeneratePrompt(params Params) (string, error) {
	tasksAsJSON, err := json.Marshal(params.Tasks)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	t, err := template.ParseFiles(params.PromptPath)
	if err != nil {
		return "", err
	}

	err = t.Execute(&tpl, map[string]interface{}{
		"Tasks":     template.HTML(tasksAsJSON),
		"DayStart":  params.DayStart,
		"DayEnd":    params.DayEnd,
		"Extension": params.PromptExtension,
	})
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

// Params represent the parameters for the prompts generator.
type Params struct {
	DayStart        string
	DayEnd          string
	PromptPath      string
	PromptExtension string
	Tasks           []todoist.Task
}
