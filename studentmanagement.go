package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func main() {
	tasks := make([]Task, 0)

	for {
		fmt.Println("\nTask Management System")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Quit")

		fmt.Print("Enter your choice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskName := strings.TrimSpace(scanner.Text())
			if taskName != "" {
				task := Task{Name: taskName, Completed: false}
				tasks = append(tasks, task)
				fmt.Printf("Task '%s' added!\n", taskName)
			} else {
				fmt.Println("Task name cannot be empty.")
			}

		case "2":
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
			} else {
				fmt.Println("Tasks:")
				for i, task := range tasks {
					status := "Incomplete"
					if task.Completed {
						status = "Completed"
					}
					fmt.Printf("%d. %s (%s)\n", i+1, task.Name, status)
				}
			}

		case "3":
			fmt.Print("Enter the task number to mark as completed: ")
			scanner.Scan()
			taskNumber := strings.TrimSpace(scanner.Text())
			if taskIndex, err := strconv.Atoi(taskNumber); err == nil && taskIndex >= 1 && taskIndex <= len(tasks) {
				taskIndex--
				tasks[taskIndex].Completed = true
				fmt.Printf("Task '%s' marked as completed!\n", tasks[taskIndex].Name)
			} else {
				fmt.Println("Invalid task number.")
			}

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
