package cmd

import (
	"fmt"
	"os"

	scheduler "github.com/salteron/todoist_tasks_scheduler"
	"github.com/spf13/cobra"
)

var dayStart, dayEnd, promptExtension, promptTemplate string

var rootCmd = &cobra.Command{
	Short: `Todoist tasks scheduler powered by GPT`,
	Long: `Integrates with Todoist API to fetch tasks for today.
	
Then composes and executes a prompt for GPT 3.5 that asks to generate a schedule for these tasks. 
Finally, outputs the GPT completion result.

Tasks should have a label of the following format: duration-N-minutes
where N is a positive integer.

Set SCHEDULER_GPT_API_TOKEN env variable to GPT token value.
Set SCHEDULER_TODOIST_API_TOKEN env variable to Todoist token value.
`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Scheduler started with the following parameters...\n")
		fmt.Printf("Day start: %s\n", dayStart)
		fmt.Printf("Day end: %s\n", dayEnd)
		fmt.Printf("Prompt extension: %s\n", promptExtension)
		fmt.Printf("Prompt template: %s\n", promptTemplate)

		schedule, err := scheduler.GenerateSchedule(scheduler.Params{
			DayStart:        dayStart,
			DayEnd:          dayEnd,
			PromptPath:      promptTemplate,
			PromptExtension: promptExtension,
		})

		if err != nil {
			return err
		}

		fmt.Printf("--- SCHEDULE ---\n%s--- SCHEDULE ---", schedule)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&dayStart, "dayStart", "08:00", "Start day (required), format: 15:04")
	rootCmd.Flags().StringVar(&dayEnd, "dayEnd", "00:00", "End day (required), format: 15:04")
	rootCmd.Flags().StringVar(&promptExtension, "promptExtension", "", "Prompt extension")
	rootCmd.Flags().StringVar(&promptTemplate, "promptTemplate", "prompts/templates/default.txt", "Prompt template")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
