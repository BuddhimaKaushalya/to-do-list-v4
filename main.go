package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	test1 "TO-DO-LIST-V4/functions"
	"TO-DO-LIST-V4/models"
)

func main() {
	file, err := os.Open("tasks.json") // open the file
	if err == nil {
		defer file.Close()
		fileInfo, _ := file.Stat()
		if fileInfo.Size() > 0 {
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&models.Tasks)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)
	var operation string

	for operation != "exit" {
		fmt.Println("What do you want to do? ")
		fmt.Println("Enter 1:AddTasks, 2:viewTasks, 3:deleteTask, 4:updateTask, 5:exit")
		operation, _ = reader.ReadString('\n')
		operation = strings.TrimSpace(operation)

		switch operation {
		case "1":
			var task string
			fmt.Print("What is the task? ")
			task, _ = reader.ReadString('\n')
			task = strings.TrimSpace(task)
			test1.AddTask(models.Task{Description: task})
			saveTasksToFile() // Save tasks immediately after adding
		case "2":
			test1.ViewTasks()
		case "3":
			var indexStr string
			fmt.Print("What is the index of the task to delete? ")
			indexStr, _ = reader.ReadString('\n')
			indexStr = strings.TrimSpace(indexStr)
			index, _ := strconv.Atoi(indexStr)
			test1.DeleteTask(index)
			saveTasksToFile() // Save tasks immediately after deleting
		case "4":
			var indexStr string
			fmt.Print("What is the index of the task to update? ")
			indexStr, _ = reader.ReadString('\n')
			indexStr = strings.TrimSpace(indexStr)
			index, _ := strconv.Atoi(indexStr)

			var task string
			fmt.Print("What is the new task? ")
			task, _ = reader.ReadString('\n')
			task = strings.TrimSpace(task)
			test1.UpdateTask(index, task)
			saveTasksToFile() // Save tasks immediately after updating
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown operation:", operation)
		}
	}
}

func saveTasksToFile() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(models.Tasks)
	if err != nil {
		fmt.Println("Error encoding tasks to JSON:", err)
		return
	}
}
