package main

import (
	"log"

	"github.com/bernardolsp/gocton/orchestrator/internal/parser"
	"github.com/bernardolsp/gocton/orchestrator/pkg/communication"
)

func main() {
	m, err := communication.Initialize(&communication.Communicator{
		Type:             "rabbit",
		ConnectionString: "amqp://guest:guest@rabbitmq:5672/", // todo: grab via env
	})
	if err != nil {
		log.Fatalf("Failed to initialize communicator: %v", err)
	}

	messager := &communication.Messager{
		Channel: m.Channel,
		Context: m.Context,
		Queue:   m.Queue,
	}

	parser := &parser.Parser{
		Messager: *messager,
	}
	parser.Initialize()
	workflow, _ := parser.ParseTaskFile("./example_taskfile.yml")
	parser.PrintJobsAndSteps(workflow)
}
