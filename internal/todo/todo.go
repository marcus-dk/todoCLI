package todo

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Task struct {
    Description string
    Completed   bool
}

var tasks []Task
var taskFile = "tasks.json"

// AddTask adds a new task to the list and saves the updated list to the file.
func AddTask(description string) {
    loadTasks() // Load existing tasks before adding a new one
    tasks = append(tasks, Task{Description: description, Completed: false})
    saveTasks()
}

// GetTasks returns all the tasks.
func GetTasks() []string {
    loadTasks()
    var taskList []string
    for i, task := range tasks {
        status := " "
        if task.Completed {
            status = "x"
        }
        taskList = append(taskList, fmt.Sprintf("%d. [%s] %s", i+1, status, task.Description))
    }
    return taskList
}

// CompleteTask marks a task as complete by its index and saves the updated list to the file.
func CompleteTask(index int) error {
    loadTasks()
    if index < 0 || index >= len(tasks) {
        return fmt.Errorf("task number %d does not exist", index)
    }
    tasks[index].Completed = true
    saveTasks()
    return nil
}

// DeleteTask deletes a task by its index and saves the updated list to the file.
func DeleteTask(index int) error {
    loadTasks()
    if index < 0 || index >= len(tasks) {
        return fmt.Errorf("task number %d does not exist", index)
    }
    tasks = append(tasks[:index], tasks[index+1:]...)
    saveTasks()
    return nil
}

// saveTasks saves the current list of tasks to a file.
func saveTasks() {
    file, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }
    err = ioutil.WriteFile(taskFile, file, 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
    }
}

// loadTasks loads the tasks from the file into memory.
func loadTasks() {
    file, err := ioutil.ReadFile(taskFile)
    if err != nil {
        if os.IsNotExist(err) {
            return // No file yet, that's okay
        }
        fmt.Println("Error reading file:", err)
        return
    }
    err = json.Unmarshal(file, &tasks)
    if err != nil {
        fmt.Println("Error parsing tasks:", err)
    }
}
