package services

import (
	"fmt"
	"my-todolist/models"
)

var tasks []models.Task
var counter int = 1

func AddTask(title string) {
    task := models.Task{ID: counter, Title: title, Done: false}
    tasks = append(tasks, task)
    counter++
    fmt.Println("Task berhasil ditambahkan:", title)
}

func ListTasks() {
    fmt.Println("\nDaftar Tugas:")
    for _, task := range tasks {
        status := "Belum"
        if task.Done {
            status = "Selesai"
        }
        fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, status)
    }
}

func CompleteTask(id int) {
    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Done = true
            fmt.Println("Task selesai:", tasks[i].Title)
            return
        }
    }
    fmt.Println("Task tidak ditemukan!")
}
