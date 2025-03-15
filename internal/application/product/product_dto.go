package product

import (
	"strings"

	"github.com/vucongthanh92/go-test-exam/helper/constants"
	"github.com/vucongthanh92/go-test-exam/helper/utils"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

func GetFilterProductList(req models.ProductListRequest) models.ProductListFilter {
	var (
		filter = models.ProductListFilter{
			Limit:  req.Limit,
			Offset: req.Offset,
			Status: []string{},
			Columns: []string{
				"products.id",
				"products.reference",
				"products.name",
				"products.added_date",
				"products.status",
				"products.category_id",
				"products.price",
				"products.stock_city",
				"products.supplier_id",
				"products.quantity",
				"categories.name as category_name",
				"suppliers.name as supplier_name",
			},
		}
	)

	status := strings.Split(req.Status, ",")
	utils.IterateSlice(status, func(i int, value string) {
		if utils.CompareEqualFold(value,
			constants.ProductTypeName.Available,
			constants.ProductTypeName.OnOrder,
			constants.ProductTypeName.OutOfStock) {
			filter.Status = append(filter.Status, value)
		}
	})

	return filter
}
