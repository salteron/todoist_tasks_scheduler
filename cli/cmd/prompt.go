package cmd

import (
	"fmt"

	scheduler "github.com/salteron/todoist_tasks_scheduler"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   `prompt`,
	Short: `Generate scheduler prompt for GPT`,
	Long: `Integrates with Todoist API to fetch tasks for today.

Then composes and outputs a prompt for GPT 3.5 that asks to generate a schedule for these tasks. 

Tasks should have a label of the following format: duration-N-minutes where N is a positive integer.

Set SCHEDULER_TODOIST_API_TOKEN env variable to Todoist token value.
`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Prompter started with the following parameters...\n")
		fmt.Printf("Day start: %s\n", dayStart)
		fmt.Printf("Day end: %s\n", dayEnd)
		fmt.Printf("Prompt extension: %s\n", promptExtension)
		fmt.Printf("Prompt template: %s\n", promptTemplate)

		prompt, err := scheduler.GeneratePrompt(scheduler.Params{
			DayStart:        dayStart,
			DayEnd:          dayEnd,
			PromptPath:      promptTemplate,
			PromptExtension: promptExtension,
		})

		if err != nil {
			return err
		}

		fmt.Printf("--- PROMPT ---\n%s--- PROMPT ---", prompt)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
	promptCmd.Flags().StringVar(&dayStart, "dayStart", "08:00", "Start day (required), format: 15:04")
	promptCmd.Flags().StringVar(&dayEnd, "dayEnd", "00:00", "End day (required), format: 15:04")
	promptCmd.Flags().StringVar(&promptExtension, "promptExtension", "", "Prompt extension")
	promptCmd.Flags().StringVar(&promptTemplate, "promptTemplate", "prompts/templates/default.txt", "Prompt template")

}
