package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/helper/validation"
	"github.com/vucongthanh92/go-test-exam/internal/application/product"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type ProductHandler struct {
	productService product.ProductService
}

func NewProductHandler(
	productService product.ProductService,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// API get products list godoc
// @Tags Product
// @Summary search products with filter and return pagination
// @Accept json
// @Produce json
// @Param  params body models.CreateCategoryReq true "CreateCategoryReq"
// @Router 	/api/v1/products [get]
// @Success	200
func (h *ProductHandler) GetProductList(c *gin.Context) {
	var (
		req    models.ProductListFilter
		paging = httpcommon.ParseParams(c)
	)

	err := validation.GetQueryParamsHTTP(c, &req)
	if err != nil {
		return
	}

	req.Limit = paging.Limit
	req.Offset = paging.Offset

	res, totalRows, errorCommon := h.productService.GetProductsByFilter(c, req)
	if errorCommon.Error != nil {
		httpcommon.ExposeError(c, errorCommon)
		return
	}

	c.JSON(http.StatusOK, httpcommon.NewPagingSuccessResponse(res, int(totalRows), nil, req.Limit))
}
