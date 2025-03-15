package category

import (
	"context"

	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req models.CreateCategoryReq) error
	GetCategoryByID(ctx context.Context) error
	UpdateCategoryByID(ctx context.Context) error
}
