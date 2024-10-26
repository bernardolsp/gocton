package nodemanager

import "github.com/bernardolsp/gocton/orchestrator/pkg/models"

type NodeManager struct {
}

func NewNodeManager() *NodeManager {
	return &NodeManager{}
}

func (nm *NodeManager) RegisterNode(node *models.Node) error {
	return nil
}
