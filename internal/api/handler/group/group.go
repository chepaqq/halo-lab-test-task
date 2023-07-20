package group

import (
	"net/http"
	"strconv"

	"github.com/chepaqq99/halo-lab-test-task/pkg/utils"
	"github.com/gin-gonic/gin"
)

type groupService interface {
	GetAverageTransparency(groupName string) (float64, error)
	GetAverageTemperature(groupName string) (float64, error)
	GetListOfSpecies(groupName string) (map[string]int, error)
	GetTopListOfSpecies(groupName string, top int) (map[string]int, error)
}

type GroupHandler struct {
	group groupService
}

func NewGroupHandler(group groupService) *GroupHandler {
	return &GroupHandler{group: group}
}

// GetAverageTransparency - Get the current average transparency inside the group.
// @Summary Get the current average transparency inside the group.
// @Tags group
// @Param groupName path string true "Name of the group"
// @Produce json
// @Success 200 {number} float64 "Average transparency"
// @Failure 500 "Internal Server Error"
// @Router /group/{groupName}/transparency/average [get]
func (h *GroupHandler) GetAverageTransparency(c *gin.Context) {
	groupName := c.Param("groupName")
	averageTransparency, err := h.group.GetAverageTransparency(groupName)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, averageTransparency)
}

// GetAverageTemperature - Get the current average temperature inside the group.
// @Summary Get the current average temperature inside the group.
// @Tags group
// @Param groupName path string true "Name of the group"
// @Produce json
// @Success 200 {number} float64 "Average temperature"
// @Failure 500 "Internal Server Error"
// @Router /group/{groupName}/temperature/average [get]
func (h *GroupHandler) GetAverageTemperature(c *gin.Context) {
	groupName := c.Param("groupName")
	averageTemperature, err := h.group.GetAverageTemperature(groupName)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, averageTemperature)
}

// GetListOfSpecies - Get the full list of fish species (with counts) currently detected inside the group.
// @Summary Get the full list of fish species (with counts) currently detected inside the group.
// @Tags group
// @Param groupName path string true "Name of the group"
// @Produce json
// @Success 200 {object} map[string]int "Map of fish species with counts"
// @Failure 500 "Internal Server Error"
// @Router /group/{groupName}/species [get]
func (h *GroupHandler) GetListOfSpecies(c *gin.Context) {
	groupName := c.Param("groupName")
	detectedFishes, err := h.group.GetListOfSpecies(groupName)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"fishSpecies": detectedFishes})
}

// GetTopListOfSpecies - Get the list of top N fish species (with counts) currently detected inside the group.
// @Summary Get the list of top N fish species (with counts) currently detected inside the group.
// @Tags group
// @Param groupName path string true "Name of the group"
// @Param N path int true "Number of top fish species to retrieve"
// @Produce json
// @Success 200 {object} map[string]int "Map of top fish species with counts"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /group/{groupName}/species/top/{N} [get]
func (h *GroupHandler) GetTopListOfSpecies(c *gin.Context) {
	groupName := c.Param("groupName")
	top, err := strconv.Atoi(c.Param("N"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid top parameter")
		return
	}
	detectedFishes, err := h.group.GetTopListOfSpecies(groupName, top)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"fishSpecies": detectedFishes})
}
