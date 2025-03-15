package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/application/category"
)

type CategoryHandler struct {
	categoryService category.CategoryService
}

func NewCategoryHandler(
	categoryService category.CategoryService,
) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// API GetCategoryList godoc
// @Tags Category
// @Summary get list categories
// @Accept json
// @Produce json
// @Router /api/v1/category [get]
// @Success 200 {object} []entities.Category
func (h *CategoryHandler) GetCategoryList(c *gin.Context) {

	res, errorCommon := h.categoryService.GetCategoryList(c)
	if errorCommon.Error != nil {
		httpcommon.ExposeError(c, errorCommon)
		return
	}

	c.JSON(http.StatusOK, res)
}
