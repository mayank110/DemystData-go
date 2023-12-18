package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"github.com/sirupsen/logrus"
	conf "DemystData/config"
)

// Todo represents the structure of a TODO item
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoInterface interface {
	FetchEvenTodos(n int, logger *logrus.Logger, config *conf.AppConfig) ([]Todo, error)
}

func NewTodoInterface() TodoInterface {
	return &Todo{}
}

func Handler(td TodoInterface, n int, logger *logrus.Logger, config *conf.AppConfig) ([]Todo, error) {
	return td.FetchEvenTodos(n, logger, config)
}


// FetchEvenTodos fetches the first n even TODOS
func (td *Todo) FetchEvenTodos(n int, logger *logrus.Logger, configs *conf.AppConfig) ([]Todo, error) {

	var wg sync.WaitGroup

	// Using channel for outout collection.
	result := make(chan Todo, n)

	for i := 2; i <= n; i += 2 {
		wg.Add(1)
		go td.fetchTodo(i, &wg, result, configs)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	var todos []Todo
	for todo := range result {
		todos = append(todos, todo)
	}

	return todos, nil
}


// Actual method to fetch the TODO Information
func (td *Todo) fetchTodo(todoID int, wg *sync.WaitGroup, result chan Todo, configs *conf.AppConfig) {
	
	defer wg.Done()
	
	url := fmt.Sprintf("%s%d", configs.TodoURL, todoID)
	
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching TODO %d: %v\n", todoID, err)
		return
	}
	
	defer response.Body.Close()

	var todo Todo
	
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Printf("Error decoding TODO %d: %v\n", todoID, err)
		return
	}

	result <- todo
}
