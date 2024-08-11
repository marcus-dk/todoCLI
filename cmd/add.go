package cmd

import (
    "fmt"
    "todoCLI/internal/todo"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add [task]",
    Short: "Add a new task to the todo list",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        task := args[0]
        todo.AddTask(task)
        fmt.Printf("Added task: %s\n", task)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
