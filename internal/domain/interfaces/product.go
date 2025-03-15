package interfaces

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type ProductQueryRepoI interface {
	GetProductByFilter(ctx context.Context, filter models.ProductListFilter) (response []entities.Product, totalRows int64, errRes httpcommon.ErrorDTO)
}

type ProductCommandRepoI interface {
}
