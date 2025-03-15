package category

import (
	"context"

	"github.com/vucongthanh92/go-base-utils/tracing"
	"github.com/vucongthanh92/go-test-exam/database"
	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
	"gorm.io/gorm"
)

type categoryQueryRepository struct {
	readDb *gorm.DB
}

func NewCategoryQueryRepository(readDb *database.GormReadDb) interfaces.CategoryQueryRepoI {
	return &categoryQueryRepository{
		readDb: *readDb,
	}
}

func (repo *categoryQueryRepository) GetCategoryList(ctx context.Context) (response []entities.Category, errRes httpcommon.ErrorDTO) {
	ctx, span := tracing.StartSpanFromContext(ctx, "GetCategoryList")
	defer span.End()

	err := repo.readDb.WithContext(ctx).Model(&entities.Category{}).
		Select("id, name").Find(&response).Error
	if err != nil {
		errRes.IsSystemError = true
		errRes.Error = err
		return response, errRes
	}

	return response, errRes
}
