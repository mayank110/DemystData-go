package main

import (
	"flag"
	"os"
	"github.com/sirupsen/logrus"
	todo "DemystData/todo"
	conf "DemystData/config"
)

var logger = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to the console (stderr)
	logger.SetOutput(os.Stderr)
	
	// Setting the level to Debug because we are printing information as output.
	logger.SetLevel(logrus.DebugLevel)

}

func main() {
	
	config := conf.GetConfig()

	numTodos := flag.Int("numTodos", config.DefaultTodo, "Number of even TODOS to fetch")
	flag.Parse()
	
	if *numTodos <= 0 {
		logger.Fatal("Please provide a positive value for the number of TODOS.")
		return
	}

	// Logging the number of Todos
	logger.Info(*numTodos)
	TodoInstance := todo.NewTodoInterface()
	todos, err := todo.Handler(TodoInstance, *numTodos, logger, config)
	if err != nil {
		logger.Fatal("Error fetching TODOS:", err)
		return
	}

	// Iterate over the response
	for _, todo := range todos {
		logger.WithFields(logrus.Fields{
			"Title":     todo.Title,
			"Completed": todo.Completed,}).Info("TODO Information")
	}
}
