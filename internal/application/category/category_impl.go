package category

import (
	"context"

	"github.com/vucongthanh92/go-base-utils/tracing"
	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
)

type CategoryImpl struct {
	categoryReadRepo interfaces.CategoryQueryRepoI
}

func NewCategoryService(
	categoryReadRepo interfaces.CategoryQueryRepoI,
) CategoryService {
	return &CategoryImpl{
		categoryReadRepo: categoryReadRepo,
	}
}

func (s *CategoryImpl) GetCategoryList(ctx context.Context) (response []entities.Category, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "GetProductsByFilter")
	defer span.End()

	response, errRes = s.categoryReadRepo.GetCategoryList(ctx)
	return response, errRes
}

func (s *CategoryImpl) GetCategoryByID(ctx context.Context) error {
	return nil
}

func (s *CategoryImpl) UpdateCategoryByID(ctx context.Context) error {
	return nil
}
