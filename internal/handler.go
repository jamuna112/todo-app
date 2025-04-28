package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Tasks struct {
	Id          int
	Topic       string
	Description string
}

var tasks []Tasks
var count int = 0

/*
 add
 ask user for topic and description as soon as user adds it display it
 id topic description
 1  Golang learn file reading writing project
 Todo added successfully
*/

func AddTask(taskList []Tasks) []Tasks {

	var task Tasks
	id, _ := ListAllTask(taskList)

	reader := bufio.NewReader(os.Stdin)

	max := slices.Max(id)
	task.Id = max + 1
	fmt.Printf("Id: %v, max: %d\n", id, max)
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

// func AddTask(taskList []Tasks) []Tasks {

// 	var task Tasks
// 	reader := bufio.NewReader(os.Stdin)
// 	count++

// 	task.Id = count
// 	fmt.Print("Topic: ")
// 	task.Topic, _ = reader.ReadString('\n')
// 	task.Topic = strings.TrimSpace(task.Topic)
// 	fmt.Print("Description: ")
// 	task.Description, _ = reader.ReadString('\n')
// 	task.Description = strings.TrimSpace(task.Description)
// 	taskList = append(taskList, task)
// 	fmt.Printf("Successfully added task id: %d - %v\n", task.Id, task)

// 	return taskList
// }

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

func SearchByID(tasks []Tasks, searchString string) {

	found := false
	for _, v := range tasks {
		if strings.ToLower(v.Topic) == strings.ToLower(searchString) || strings.Contains(strings.ToLower((v.Description)), strings.ToLower(searchString)) {
			fmt.Printf("Id: %v, topic: %s, description: %s\n", v.Id, v.Topic, v.Description)
			found = true
			break
		}

	}
	if !found {
		fmt.Printf("%s string not found in the tasks\n", searchString)

	}

}

func ReadFile(filename string) []Tasks {
	data, err := os.ReadFile(filename)
	var todo []Tasks

	if err != nil {
		fmt.Printf("error reading file %v", err)

		//creating a file
		file, err := os.Create(filename)

		if err != nil {
			fmt.Println("error creating file", err)
			return nil
		}
		defer file.Close()
		fmt.Print("File created ", filename)
		//new file created

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

func ListAllTask(task []Tasks) ([]int, []string) {
	var ids []int
	var displayList []string
	var displayMsg string
	for _, v := range task {
		displayMsg = fmt.Sprintf("s.no: %v. Topic: %s, Description: %s", v.Id, v.Topic, v.Description)
		displayList = append(displayList, displayMsg)

		ids = append(ids, v.Id)
	}
	return ids, displayList
}

func DisplayAllListOfTask(task []Tasks) {
	fmt.Println("====================================")
	fmt.Println("List of all tasks")
	fmt.Println("====================================")
	_, msg := ListAllTask(task)
	for _, v := range msg {
		fmt.Println(v)
	}
	fmt.Println("====================================")
}
