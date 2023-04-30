package todoist

import (
	"regexp"
	"strconv"
	"time"
)

// Returns tasks that are due today.
func FilterTasksByDate(tasks []Task, date time.Time) []Task {
	var filteredTasks []Task

	for _, task := range tasks {
		if task.Due.Date.Time.Year() == date.Year() && task.Due.Date.Time.YearDay() == date.YearDay() {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}

/*
Returns tasks with a duration label, fills the Duration field with the
corresponding value.
*/
func FilterTasksWithDurations(tasks []Task) []Task {
	tasksWithDurations := []Task{}

	durationRegexp := regexp.MustCompile(`^duration-(\d+)-minutes$`)

	for _, task := range tasks {
		for _, label := range task.Labels {
			matches := durationRegexp.FindStringSubmatch(label)

			if matches != nil {
				minutes, _ := strconv.Atoi(matches[1])
				task.Duration = time.Duration(minutes) * time.Minute
				tasksWithDurations = append(tasksWithDurations, task)
				break
			}
		}
	}

	return tasksWithDurations
}
