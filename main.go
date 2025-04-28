package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jamuna112/todo-app.git/internal"
)

func main() {
	var choice int
	var response string
	fmt.Println("===================================")
	fmt.Println("Welcome to TODO APP")
	fmt.Println("===================================")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please select from below menu 1 - 6:")
		fmt.Println("1. Add task")
		fmt.Println("2. Update task")
		fmt.Println("3. Delete task")
		fmt.Println("4. Search task")
		fmt.Println("5. List all the tasks")
		fmt.Println("6. Exit")
		fmt.Scanf("%v", &choice)

		filename := "tasks.json"
		var t []internal.Tasks
		data := internal.ReadFile(filename) // read from the file

		switch choice {
		case 1:
			fmt.Println("Adding task..")
			data = internal.AddTask(data)
			internal.WriteJsonRecord(filename, data)
		case 2:
			var taskId int

			fmt.Print("Please enter id you want to update: ")
			fmt.Scanf("%d", &taskId)
			fmt.Print("Update topic: ")
			topic, _ := reader.ReadString('\n')
			topic = strings.TrimSpace(topic)
			fmt.Print("Update description:")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			t = internal.UpdateJsonRecord(data, taskId, topic, description)
			internal.WriteJsonRecord(filename, t)
			fmt.Printf("update data %v\n", t)

		case 3:
			// Delete task
			var deleteId int
			fmt.Print("Enter Id you want to delete?")
			fmt.Scanf("%d", &deleteId)
			t = internal.DeleteTask(filename, deleteId)
			internal.WriteJsonRecord(filename, t)
			fmt.Printf("After deletion remaining data %v\n", t)

		case 4:
			//Search task
			var searchId int
			fmt.Print("Enter Id you want to search?")
			fmt.Scanf("%d", &searchId)
			data = internal.ReadFile(filename)
			fmt.Println("====================================")
			internal.SearchByID(data, searchId)
			fmt.Println("====================================")

		case 5:
			fmt.Println("====================================")
			fmt.Println("List of all tasks")
			fmt.Println("====================================")
			internal.ListAllTask(data)
			fmt.Println("====================================")
		case 6:
			fmt.Println("Thank you choosing TODO APP!")
			return
		default:
			fmt.Println("please select valid input")

		}

		for {
			fmt.Print("Do you want to continue? ")
			fmt.Scanf("%s", &response)
			if response == "yes" {
				break
			} else if response == "no" {
				fmt.Println("Good bye..")
				return
			} else {
				fmt.Println("Please enter valid input..")
			}

		}
	}

}
