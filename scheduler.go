/*
Todoist tasks scheduler powered by GPT.

Integrates with Todoist API to fetch tasks for today.

Then composes and executes a prompt for GPT 3.5 that asks to generate a schedule for these tasks.
Finally, outputs the GPT completion result.

Tasks should have a label of the following format: duration-N-minutes
where N is a positive integer.

Set SCHEDULER_GPT_API_TOKEN env variable to GPT token value.
Set SCHEDULER_TODOIST_API_TOKEN env variable to Todoist token value.
*/
package scheduler

import (
	"time"

	"github.com/salteron/todoist_tasks_scheduler/gpt"
	"github.com/salteron/todoist_tasks_scheduler/prompts"
	"github.com/salteron/todoist_tasks_scheduler/todoist"
)

// GenerateSchedule returns a GPT completion containing a schedule for today's
// Todoist tasks.
func GenerateSchedule(params Params) (schedule string, err error) {
	prompt, err := GeneratePrompt(params)

	if err != nil {
		return
	}

	return gpt.Complete(prompt)
}

// GeneratePrompt returns a prompt that asks to generate a schedule for Todoist
// tasks.
func GeneratePrompt(params Params) (prompt string, err error) {
	err = params.validate()
	if err != nil {
		return
	}

	tasks, err := todoist.GetTasks()
	if err != nil {
		return
	}

	tasks = todoist.FilterTasksByDate(
		todoist.FilterTasksWithDurations(tasks),
		time.Now(),
	)

	return prompts.GeneratePrompt(prompts.Params{
		DayStart:        params.DayStart,
		DayEnd:          params.DayEnd,
		PromptPath:      params.PromptPath,
		PromptExtension: params.PromptExtension,
		Tasks:           tasks,
	})
}

// Params represent the parameters for the scheduler
type Params struct {
	DayStart        string // when a day starts, format: 15:04
	DayEnd          string // when a day ends, 	 format: 15:04
	PromptPath      string // relative prompt template path
	PromptExtension string // arbitrary prompt extension
}

func (p Params) validate() error {
	_, err := time.Parse("15:04", p.DayStart)
	if err != nil {
		return err
	}

	_, err = time.Parse("15:04", p.DayEnd)
	if err != nil {
		return err
	}

	return nil
}
