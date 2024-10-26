package distributor

import "github.com/bernardolsp/gocton/orchestrator/pkg/models"

type TaskDistributor struct {
}

func NewTaskDistributor() *TaskDistributor {
	return &TaskDistributor{}
}

func (td *TaskDistributor) DistributeTask(task *models.Task) error {
	return nil
}
