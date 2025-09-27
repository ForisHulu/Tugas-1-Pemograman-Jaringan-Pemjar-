package main

import (
    "fmt"
    "my-todolist/services"
)

func main() {
    for {
        fmt.Println("\nMENU TO-DO LIST")
        fmt.Println("1. Tambah Task")
        fmt.Println("2. Lihat Task")
        fmt.Println("3. Selesaikan Task")
        fmt.Println("4. Keluar")

        var choice int
        fmt.Print("Pilih menu: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var title string
            fmt.Print("Masukkan nama task: ")
            fmt.Scan(&title)
            services.AddTask(title)
        case 2:
            services.ListTasks()
        case 3:
            var id int
            fmt.Print("Masukkan ID task: ")
            fmt.Scan(&id)
            services.CompleteTask(id)
        case 4:
            fmt.Println("Keluar dari aplikasi")
            return
        default:
            fmt.Println("Pilihan tidak valid")
        }
    }
}
