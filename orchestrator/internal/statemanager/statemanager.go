package statemanager

import "github.com/bernardolsp/gocton/orchestrator/pkg/models"

type StateManager struct {
}

func NewStateManager() *StateManager {
	return &StateManager{}
}

func (sm *StateManager) UpdateState(pipelineID string, state *models.PipelineState) error {
	return nil
}
