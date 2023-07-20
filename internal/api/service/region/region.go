package region

type regionRepository interface {
	GetMinTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error)
	GetMaxTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error)
}

type RegionService struct {
	repo regionRepository
}

func NewRegionService(regionRepository regionRepository) *RegionService {
	return &RegionService{repo: regionRepository}
}

func (s *RegionService) GetMinTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error) {
	return s.repo.GetMinTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax)
}

func (s *RegionService) GetMaxTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error) {
	return s.repo.GetMaxTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax)
}
