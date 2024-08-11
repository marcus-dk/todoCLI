package cmd

import (
    "fmt"
    "todoCLI/internal/todo"
    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks := todo.GetTasks()
        for i, task := range tasks {
            fmt.Printf("%d: %s\n", i+1, task)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
