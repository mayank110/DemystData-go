package main

import (
	"flag"
	"os"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to the console (stderr)
	logger.SetOutput(os.Stderr)
}

// type PaymentProcessorService struct {
// 	cfg *ms.AppConfig
// }

// func main() {
	// cfg := ms.GetConfig()
	// l := logging.MustGetLogger("microservice.PaymentProcessorService", cfg.Logs)
	// environment.Init(l)
	// svc := &PaymentProcessorService{DefaultService: microservice.DefaultService{Log: l}, cfg: cfg}
	// s := microservice.NewServer(l, microservice.WithService(svc))
	// s.Start()
// }

func main() {
	numTodos := flag.Int("numTodos", 20, "Number of even TODOS to fetch")
	flag.Parse()
	
	if *numTodos <= 0 {
		logger.Fatal("Please provide a positive value for the number of TODOS.")
		return
	}

	logger.Info(*numTodos)
	// todos, err := FetchEvenTodos(*numTodos)
	// if err != nil {
	// 	fmt.Printf("Error fetching TODOS: %v\n", err)
	// 	return
	// }

	// for _, todo := range todos {
	// 	fmt.Printf("Title: %s, Completed: %v\n", todo.Title, todo.Completed)
	// }
}
