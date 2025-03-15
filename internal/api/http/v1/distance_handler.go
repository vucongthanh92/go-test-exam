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

// API get products list godoc
// @Tags Supplier
// @Summary search products with filter and return pagination
// @Accept json
// @Produce json
// @Param  params body models.CreateCategoryReq true "CreateCategoryReq"
// @Router 	/api/v1/products [get]
// @Success	200
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
