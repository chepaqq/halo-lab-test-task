package region

import (
	"net/http"
	"strconv"

	"github.com/chepaqq99/halo-lab-test-task/pkg/utils"
	"github.com/gin-gonic/gin"
)

type regionService interface {
	GetMinTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error)
	GetMaxTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error)
}

type RegionHandler struct {
	region regionService
}

func NewRegionHandler(region regionService) *RegionHandler {
	return &RegionHandler{region: region}
}

// GetMinTemperatureInRegion - .
func (h *RegionHandler) GetMinTemperatureInRegion(c *gin.Context) {
	xMin, err := strconv.Atoi(c.Query("xMin"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid xMin parameter")
		return
	}

	yMin, err := strconv.Atoi(c.Query("yMin"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid yMin parameter")
		return
	}

	zMin, err := strconv.Atoi(c.Query("zMin"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid zMin parameter")
		return
	}

	xMax, err := strconv.Atoi(c.Query("xMax"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid xMax parameter")
		return
	}

	yMax, err := strconv.Atoi(c.Query("yMax"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid yMax parameter")
		return
	}

	zMax, err := strconv.Atoi(c.Query("zMax"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid zMax parameter")
		return
	}

	minTemperature, err := h.region.GetMinTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, minTemperature)
}

// GetMaxTemperatureInRegion - .
func (h *RegionHandler) GetMaxTemperatureInRegion(c *gin.Context) {
	xMin, err := strconv.Atoi(c.Query("xMin"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid xMin parameter")
		return
	}

	yMin, err := strconv.Atoi(c.Query("yMin"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid yMin parameter")
		return
	}

	zMin, err := strconv.Atoi(c.Query("zMin"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid zMin parameter")
		return
	}

	xMax, err := strconv.Atoi(c.Query("xMax"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid xMax parameter")
		return
	}

	yMax, err := strconv.Atoi(c.Query("yMax"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid yMax parameter")
		return
	}

	zMax, err := strconv.Atoi(c.Query("zMax"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid zMax parameter")
		return
	}

	maxTemperature, err := h.region.GetMaxTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, maxTemperature)
}
