package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/helper/utils"
	"github.com/vucongthanh92/go-test-exam/helper/validation"
	"github.com/vucongthanh92/go-test-exam/internal/application/distance"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type DistanceHandler struct {
	distanceService distance.DistanceService
}

func NewDistanceHandler(
	distanceService distance.DistanceService,
) *DistanceHandler {
	return &DistanceHandler{
		distanceService: distanceService,
	}
}

// API CalculateDistanceStockCity godoc
// @Tags Distance
// @Summary calculate distance from IP to Stock city
// @Accept json
// @Produce json
// @Param city query string false "city name"
// @Router 	/api/v1/distance/stock_city [get]
// @Success 200 {object} models.CalculateDistanceStockCityRes
func (h *DistanceHandler) CalculateDistanceStockCity(c *gin.Context) {

	var req models.CalculateDistanceStockCityReq

	err := validation.GetQueryParamsHTTP(c, &req)
	if err != nil {
		return
	}

	req.ClientID = utils.GetIPFromClient(c)

	res, errorCommon := h.distanceService.CalculateDistanceStockCity(c, req)
	if errorCommon.Error != nil {
		httpcommon.ExposeError(c, errorCommon)
		return
	}

	c.JSON(http.StatusOK, res)
}
