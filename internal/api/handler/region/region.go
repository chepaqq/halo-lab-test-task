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

// GetMinTemperatureInRegion - Get the minimum temperature in the specified region.
// @Summary Get the minimum temperature in the specified region.
// @Tags region
// @Param xMin query int true "Minimum x-coordinate of the region"
// @Param yMin query int true "Minimum y-coordinate of the region"
// @Param zMin query int true "Minimum z-coordinate of the region"
// @Param xMax query int true "Maximum x-coordinate of the region"
// @Param yMax query int true "Maximum y-coordinate of the region"
// @Param zMax query int true "Maximum z-coordinate of the region"
// @Produce json
// @Success 200 {number} float64 "Minimum temperature in the region"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /region/temperature/min [get]
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

// GetMaxTemperatureInRegion - Get the maximum temperature in the specified region.
// @Summary Get the maximum temperature in the specified region.
// @Tags region
// @Param xMin query int true "Minimum x-coordinate of the region"
// @Param yMin query int true "Minimum y-coordinate of the region"
// @Param zMin query int true "Minimum z-coordinate of the region"
// @Param xMax query int true "Maximum x-coordinate of the region"
// @Param yMax query int true "Maximum y-coordinate of the region"
// @Param zMax query int true "Maximum z-coordinate of the region"
// @Produce json
// @Success 200 {number} float64 "Maximum temperature in the region"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /region/temperature/max [get]
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
