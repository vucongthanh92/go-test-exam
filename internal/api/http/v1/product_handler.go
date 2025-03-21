package v1

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vucongthanh92/go-test-exam/helper/constants"
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

// API GetProductList godoc
// @Tags Product
// @Summary search products with filter and return pagination
// @Accept json
// @Produce json
// @Param pageSize query string false "page size"
// @Param pageIndex query string false "page index"
// @Param status query string false "get product by status" default(Available, Out of Stock, On Order)
// @Router 	/api/v1/product/search [get]
// @Success 200 {object} []models.ProductListResponse
func (h *ProductHandler) GetProductList(c *gin.Context) {
	var (
		req    models.ProductListRequest
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

	c.JSON(http.StatusOK, httpcommon.NewPagingSuccessResponse(res, int(totalRows), constants.ProductTypeName, req.Limit))
}

// API GenProductListToPDF godoc
// @Tags Product
// @Summary generate file PDF by product list
// @Accept json
// @Produce json
// @Router 	/api/v1/product/gen-pdf [get]
// @Success	200
func (h *ProductHandler) GenProductListToPDF(c *gin.Context) {
	var (
		req models.ProductListRequest
	)

	err := validation.GetQueryParamsHTTP(c, &req)
	if err != nil {
		return
	}

	req.Limit = 0
	req.Offset = 0

	// get product list
	productList, _, errRes := h.productService.GetProductsByFilter(c, req)
	if errRes.Error != nil {
		httpcommon.ExposeError(c, errRes)
		return
	}

	filePath, fileName, errRes := h.productService.GenProductListToPDF(c, productList)
	if errRes.Error != nil {
		httpcommon.ExposeError(c, errRes)
		return
	}

	defer os.Remove(filePath)
	c.FileAttachment(filePath, fileName)
}

// API StatisticsProductPerCategory godoc
// @Tags Product
// @Summary get statistics product per category
// @Accept json
// @Produce json
// @Router 	/api/v1/statistics/products-per-category [get]
// @Success 200 {object} []models.StatisticsProductPerCategory
func (h *ProductHandler) StatisticsProductPerCategory(c *gin.Context) {

	res, errorCommon := h.productService.StatisticsProductPerCategory(c)
	if errorCommon.Error != nil {
		httpcommon.ExposeError(c, errorCommon)
		return
	}

	c.JSON(http.StatusOK, res)
}

// API StatisticsProductPerSupplier godoc
// @Tags Product
// @Summary get statistics product per supplier
// @Accept json
// @Produce json
// @Router 	/api/v1/statistics/products-per-supplier [get]
// @Success 200 {object} []models.StatisticsProductPerSupplier
func (h *ProductHandler) StatisticsProductPerSupplier(c *gin.Context) {

	res, errorCommon := h.productService.StatisticsProductPerSupplier(c)
	if errorCommon.Error != nil {
		httpcommon.ExposeError(c, errorCommon)
		return
	}

	c.JSON(http.StatusOK, res)
}
