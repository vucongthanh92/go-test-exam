package interfaces

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type ProductQueryRepoI interface {
	GetProductByFilter(ctx context.Context, filter models.ProductListFilter) (response []models.ProductListResponse, totalRows int64, errRes httpcommon.ErrorDTO)
	StatisticsProductPerCategory(ctx context.Context) (response []models.StatisticsProductPerCategory, errRes httpcommon.ErrorDTO)
	StatisticsProductPerSupplier(ctx context.Context) (response []models.StatisticsProductPerSupplier, errRes httpcommon.ErrorDTO)
}

type ProductCommandRepoI interface {
}
