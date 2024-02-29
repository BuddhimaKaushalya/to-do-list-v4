package test1

import (
	"fmt"
	"time"

	"TO-DO-LIST-V4/models"
)

func AddTask(task models.Task) {
	task.CreatedAt = time.Now()
	models.Tasks = append(models.Tasks, task)
	fmt.Println("Task added:", task.Description)
}

func ViewTasks() {
	fmt.Println("Tasks:-")
	for index, task := range models.Tasks {
		fmt.Printf("%d. %s\n", index+1, task.Description)
	}
}

func DeleteTask(index int) {
	if index >= 1 && index <= len(models.Tasks) {
		deletedTask := models.Tasks[index-1]
		models.Tasks = append(models.Tasks[:index-1], models.Tasks[index:]...)
		fmt.Println("Task deleted:", deletedTask.Description)
	} else if index == 1 {
		deletedTask := models.Tasks[0]
		models.Tasks = models.Tasks[1:]
		fmt.Println("Task deleted:", deletedTask.Description)
	} else {
		fmt.Println("Invalid index", index)
	}
}

func UpdateTask(index int, newTask string) {
	if index >= 1 && index <= len(models.Tasks) {
		models.Tasks[index-1].Description = newTask
		models.Tasks[index-1].CreatedAt = time.Now()
		fmt.Println("Task updated:", newTask)
	} else {
		fmt.Println("Invalid index", index)
	}
}
