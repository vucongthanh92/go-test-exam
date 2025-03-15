package product

import (
	"context"

	"github.com/vucongthanh92/go-base-utils/tracing"
	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type ProductImpl struct {
	productReadRepo interfaces.ProductQueryRepoI
}

func NewProductService(productReadRepo interfaces.ProductQueryRepoI) ProductService {
	return &ProductImpl{
		productReadRepo: productReadRepo,
	}
}

func (s *ProductImpl) GetProductsByFilter(ctx context.Context, req models.ProductListRequest) (
	response []models.ProductListResponse, totalRows int64, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "GetProductsByFilter")
	defer span.End()

	filter := GetFilterProductList(req)

	response, totalRows, errRes = s.productReadRepo.GetProductByFilter(ctx, filter)
	return response, totalRows, errRes
}

func (s *ProductImpl) GetProductByID(ctx context.Context) error {
	return nil
}
