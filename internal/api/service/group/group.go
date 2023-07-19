package group

import "github.com/chepaqq99/halo-lab-test-task/internal/models"

type groupRepository interface {
	GetAverageTransparency(groupName string) (float64, error)
	GetAverageTemperature(groupName string) (float64, error)
	GetGroupByName(groupName string) (*models.SensorGroup, error)
}

type GroupService struct {
	repo groupRepository
}

func NewGroupService(groupRepository groupRepository) *GroupService {
	return &GroupService{repo: groupRepository}
}

// GetAverageTransparency - .
func (s *GroupService) GetAverageTransparency(groupName string) (float64, error) {
	group, err := s.repo.GetGroupByName(groupName)
	if err != nil || group == nil {
		return 0, err
	}
	return s.repo.GetAverageTransparency(groupName)
}

func (s *GroupService) GetAverageTemperature(groupName string) (float64, error) {
	group, err := s.repo.GetGroupByName(groupName)
	if err != nil || group == nil {
		return 0, err
	}
	return s.repo.GetAverageTemperature(groupName)
}
