package cmd

import (
    "fmt"
    "todoCLI/internal/todo"
    "github.com/spf13/cobra"
    "strconv"
)

var completeCmd = &cobra.Command{
    Use:   "complete [task number]",
    Short: "Mark a task as complete",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        index, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task number")
            return
        }

        if err := todo.CompleteTask(index - 1); err != nil {
            fmt.Println(err)
            return
        }

        fmt.Printf("Marked task %d as complete\n", index)
    },
}

func init() {
    rootCmd.AddCommand(completeCmd)
}
