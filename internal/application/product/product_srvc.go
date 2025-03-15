package product

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type ProductService interface {
	GetProductsByFilter(ctx context.Context, req models.ProductListRequest) (response []models.ProductListResponse, totalRows int64, errRes httpcommon.ErrorDTO)
	GenProductListToPDF(ctx context.Context, req []models.ProductListResponse) (filePath, fileName string, errRes httpcommon.ErrorDTO)
	StatisticsProductPerCategory(ctx context.Context) (response []models.StatisticsProductPerCategory, errRes httpcommon.ErrorDTO)
	StatisticsProductPerSupplier(ctx context.Context) (response []models.StatisticsProductPerSupplier, errRes httpcommon.ErrorDTO)
}
