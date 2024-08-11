package cmd

import (
    "fmt"
    "todoCLI/internal/todo"
    "github.com/spf13/cobra"
    "strconv"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [task number]",
    Short: "Delete a task from the todo list",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        index, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task number")
            return
        }

        if err := todo.DeleteTask(index - 1); err != nil {
            fmt.Println(err)
            return
        }

        fmt.Printf("Deleted task %d\n", index)
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
