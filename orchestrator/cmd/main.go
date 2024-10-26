package main

import (
	"log"

	"github.com/bernardolsp/gocton/orchestrator/internal/parser"
)

func main() {
	// comm, err := communication.NewCommunicator()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize communicator: %v", err)
	// }

	// engine := engine.NewWorkflowEngine()
	// distributor := distributor.NewTaskDistributor()
	// nodeManager := nodemanager.NewNodeManager()
	// stateManager := statemanager.NewStateManager()

	log.Println("Orchestrator started")

	p, err := parser.ParseTaskFile("./example_taskfile.yml")
	if err != nil {
		log.Fatalf("Error, %v", err)
	}
	parser.PrintJobsAndSteps(p)
}
