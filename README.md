# Todoist Tasks Scheduler

This is a Go application that integrates with the Todoist API and GPT 3.5 to generate a schedule for tasks due today.

## Installation

To use this application, you will need to set two environment variables: `SCHEDULER_GPT_API_TOKEN` and `SCHEDULER_TODOIST_API_TOKEN`. These should be set to your GPT token value and your Todoist token value, respectively.

You can then build the application by running the following command:

```
go build -o scheduler cli/main.go
```

## Usage

```
Usage:
   [flags]
   [command]

Available Commands:
  help        Help about any command
  prompt      Generate scheduler prompt for GPT

Flags:
      --dayEnd string            End day (required), format: 15:04 (default "00:00")
      --dayStart string          Start day (required), format: 15:04 (default "08:00")
  -h, --help                     help for this command
      --promptExtension string   Prompt extension
      --promptTemplate string    Prompt template (default "prompts/templates/default.txt")

Use " [command] --help" for more information about a command.
```

## Example

Generate a prompt to schedule tasks for today starting at 09:00 and asking to
sort tasks by descending their duration. 

`$ ./scheduler prompt --dayStart "09:00" --promptExtension "Sort tasks by duration descending."`

## Task Label Format

Tasks should have a label of the following format: `duration-N-minutes`, where `N` is a positive integer representing the duration of the task in minutes.

## Contributing

Contributions are welcome! Please open a pull request with your changes.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
