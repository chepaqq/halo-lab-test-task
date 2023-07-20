package group

import "github.com/chepaqq99/halo-lab-test-task/internal/models"

type groupRepository interface {
	Create(groupName string) (int, error)
	GetAverageTransparency(groupName string) (float64, error)
	GetAverageTemperature(groupName string) (float64, error)
	GetGroupByName(groupName string) (*models.SensorGroup, error)
	GetListOfSpecies(groupName string) (map[string]int, error)
	GetTopListOfSpecies(groupName string, top int) (map[string]int, error)
	GetLastIDInGroup(groupID int) (int, error)
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

func (s *GroupService) GetListOfSpecies(groupName string) (map[string]int, error) {
	group, err := s.repo.GetGroupByName(groupName)
	if err != nil || group == nil {
		return nil, err
	}
	return s.repo.GetListOfSpecies(groupName)
}

func (s *GroupService) GetTopListOfSpecies(groupName string, top int) (map[string]int, error) {
	group, err := s.repo.GetGroupByName(groupName)
	if err != nil || group == nil {
		return nil, err
	}
	return s.repo.GetTopListOfSpecies(groupName, top)
}

func (s *GroupService) CreateGroup(groupName string) (int, error) {
	group, err := s.repo.GetGroupByName(groupName)
	if group == nil {
		return s.repo.Create(groupName)
	}
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (s *GroupService) GetLastIDInGroup(groupID int) (int, error) {
	return s.repo.GetLastIDInGroup(groupID)
}
