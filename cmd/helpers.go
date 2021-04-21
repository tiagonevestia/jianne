package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func printErrorAndExit(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func readFromFile(filepath string) ([]*Todo, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var todos []*Todo
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, fmt.Errorf("não foi possível decodificar o arquivo JSON: %v", err)
	}

	return todos, nil
}

func writeToFile(todos []*Todo, filepath string) error {
	bytes, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return fmt.Errorf("não foi possível codificar as tarefas para JSON: %v", err)
	}
	if err := ioutil.WriteFile(filepath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func sortTodos(todos []*Todo) {
	sort.Slice(todos, func(i, j int) bool {
		x, y := todos[i], todos[j]

		if x.Priority != y.Priority {
			return x.Priority < y.Priority
		}
		return x.CreatedAt.After(y.CreatedAt)
	})
}

func filter(todos []*Todo, callback func(*Todo) bool) []*Todo {
	var res []*Todo
	for _, t := range todos {
		if callback(t) {
			res = append(res, t)
		}
	}
	return res
}

func titles(todos []*Todo) []string {
	var res []string
	for _, t := range todos {
		res = append(res, t.Title)
	}
	return res
}

func findByID(todos []*Todo, id string) *Todo {
	var t *Todo
	for _, todo := range todos {
		if todo.ID == id {
			t = todo
			break
		}
	}
	return t
}

func findByTitle(todos []*Todo, title string) *Todo {
	var t *Todo

	for _, todo := range todos {
		if todo.Title == title {
			t = todo
			break
		}
	}

	return t
}

func contains(elements []string, target string) bool {

	for _, elem := range elements {
		if elem == target {
			return true
		}
	}

	return false
}
