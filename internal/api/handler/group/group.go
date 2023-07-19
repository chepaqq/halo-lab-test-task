package group

import (
	"net/http"

	"github.com/chepaqq99/halo-lab-test-task/pkg/utils"
	"github.com/gin-gonic/gin"
)

type groupService interface {
	GetAverageTransparency(groupName string) (float64, error)
	GetAverageTemperature(groupName string) (float64, error)
	GetListOfSpecies(groupName string) (map[string]int, error)
}

type GroupHandler struct {
	group groupService
}

func NewGroupHandler(group groupService) *GroupHandler {
	return &GroupHandler{group: group}
}

// GetAverageTransparency - .
func (h *GroupHandler) GetAverageTransparency(c *gin.Context) {
	groupName := c.Param("groupName")
	averageTransparency, err := h.group.GetAverageTransparency(groupName)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, averageTransparency)
}

// GetAverageTemperature - .
func (h *GroupHandler) GetAverageTemperature(c *gin.Context) {
	groupName := c.Param("groupName")
	averageTemperature, err := h.group.GetAverageTemperature(groupName)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, averageTemperature)
}

// GetListOfSpecies - .
func (h *GroupHandler) GetListOfSpecies(c *gin.Context) {
	groupName := c.Param("groupName")
	detectedFishes, err := h.group.GetListOfSpecies(groupName)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"fishSpecies": detectedFishes})
}
