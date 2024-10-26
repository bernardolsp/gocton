package main

import (
	"log"

	"github.com/bernardolsp/gocton/orchestrator/internal/distributor"
	"github.com/bernardolsp/gocton/orchestrator/internal/engine"
	"github.com/bernardolsp/gocton/orchestrator/internal/nodemanager"
	"github.com/bernardolsp/gocton/orchestrator/internal/statemanager"
	"github.com/bernardolsp/gocton/orchestrator/pkg/communication"
)

func main() {
	comm, err := communication.NewCommunicator()
	if err != nil {
		log.Fatalf("Failed to initialize communicator: %v", err)
	}

	engine := engine.NewWorkflowEngine()
	distributor := distributor.NewTaskDistributor()
	nodeManager := nodemanager.NewNodeManager()
	stateManager := statemanager.NewStateManager()

	// TODO: Implement main orchestrator logic
	log.Println("Orchestrator started")
}
