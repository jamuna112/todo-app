package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Tasks struct {
	Id          int
	Topic       string
	Description string
}

var tasks []Tasks
var count int = 0

func AddTask(taskList []Tasks) []Tasks {

	var task Tasks
	reader := bufio.NewReader(os.Stdin)
	count++

	task.Id = count
	fmt.Print("Topic: ")
	task.Topic, _ = reader.ReadString('\n')
	task.Topic = strings.TrimSpace(task.Topic)
	fmt.Print("Description: ")
	task.Description, _ = reader.ReadString('\n')
	task.Description = strings.TrimSpace(task.Description)
	taskList = append(taskList, task)
	fmt.Printf("Successfully added task id: %d - %v\n", task.Id, task)

	return taskList
}

func UpdateJsonRecord(updateTask []Tasks, id int, topic, description string) []Tasks {

	for i, v := range updateTask {
		if v.Id == id {
			updateTask[i].Topic = topic
			updateTask[i].Description = description
			break
		}
	}

	return updateTask
}

func DeleteTask(filename string, id int) []Tasks {

	tasks := ReadFile(filename)

	var newTasks []Tasks

	for _, v := range tasks {
		if v.Id != id {
			newTasks = append(newTasks, v)
		}
	}
	fmt.Println("Deleted successfully!")
	return newTasks
}

func SearchByID(tasks []Tasks, id int) {

	found := false
	for _, v := range tasks {
		if v.Id == id {
			fmt.Printf("Id: %v, topic: %s, description: %s\n", v.Id, v.Topic, v.Description)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("%d id not found in the tasks\n", id)

	}

}

func ReadFile(filename string) []Tasks {
	data, err := os.ReadFile(filename)
	var todo []Tasks

	if err != nil {
		fmt.Printf("error reading file %v", err)
		return nil
	}

	err = json.Unmarshal(data, &todo)
	if err != nil {
		fmt.Printf("error parsing json %v\n", err)
		return nil
	}
	return todo
}

func WriteJsonRecord(filename string, todo []Tasks) []Tasks {

	data, err := json.MarshalIndent(todo, "", " ")
	if err != nil {
		fmt.Printf("error formating data: %v\n", err)
		return nil
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Printf("Error writing json data: %v\n", err)
		return nil
	}
	return todo
}

func ListAllTask(task []Tasks) {
	for i, v := range task {
		fmt.Printf("s.no: %v. Topic: %s, Description: %s\n", (i + 1), v.Topic, v.Description)
	}
}
