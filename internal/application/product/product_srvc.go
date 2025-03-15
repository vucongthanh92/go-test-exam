package product

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type ProductService interface {
	CreateProduct(ctx context.Context) error
	GetProductsByFilter(ctx context.Context, req models.ProductListFilter) (response []entities.Product, totalRows int64, errRes httpcommon.ErrorDTO)
	GetProductByID(ctx context.Context) error
	UpdateProductByID(ctx context.Context) error
	DeleteProductByID(ctx context.Context) error
}
