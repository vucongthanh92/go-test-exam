package category

import (
	"context"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
)

type CategoryService interface {
	GetCategoryList(ctx context.Context) (response []entities.Category, errRes httpcommon.ErrorDTO)
}
