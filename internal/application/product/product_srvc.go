package product

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type ProductService interface {
	GetProductsByFilter(ctx context.Context, req models.ProductListRequest) (response []models.ProductListResponse, totalRows int64, errRes httpcommon.ErrorDTO)
}
