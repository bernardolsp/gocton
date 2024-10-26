package engine

import "github.com/bernardolsp/gocton/orchestrator/internal/parser"

type WorkflowEngine struct {
}

func NewWorkflowEngine() *WorkflowEngine {
	return &WorkflowEngine{}
}

func (we *WorkflowEngine) ExecuteWorkflow(taskFile *parser.TaskFile) error {
	return nil
}
