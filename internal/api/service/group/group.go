package group

type groupRepository interface {
	GetAverageTransparency(groupName string) (float64, error)
	GetAverageTemperature(groupName string) (float64, error)
}

type GroupService struct {
	repo groupRepository
}

func NewGroupService(groupRepository groupRepository) *GroupService {
	return &GroupService{repo: groupRepository}
}

func (s *GroupService) GetAverageTransparency(groupName string) (float64, error) {
	return s.repo.GetAverageTransparency(groupName)
}

func (s *GroupService) GetAverageTemperature(groupName string) (float64, error) {
	return s.repo.GetAverageTemperature(groupName)
}
