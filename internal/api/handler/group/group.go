package group

import (
	"net/http"

	"github.com/chepaqq99/halo-lab-test-task/pkg/utils"
	"github.com/gin-gonic/gin"
)

type groupService interface {
	GetAverageTransparency(groupName string) (float64, error)
	GetAverageTemperature(groupName string) (float64, error)
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
