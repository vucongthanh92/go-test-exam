package product

import (
	"context"

	"github.com/vucongthanh92/go-base-utils/tracing"
	"github.com/vucongthanh92/go-test-exam/database"
	"gorm.io/gorm"

	httpcommon "github.com/vucongthanh92/go-test-exam/helper/http_common"
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
	"github.com/vucongthanh92/go-test-exam/internal/domain/interfaces"
	"github.com/vucongthanh92/go-test-exam/internal/domain/models"
)

type productQueryRepository struct {
	readDb *gorm.DB
}

func NewProductQueryRepository(readDb *database.GormReadDb) interfaces.ProductQueryRepoI {
	return &productQueryRepository{
		readDb: *readDb,
	}
}

func (repo *productQueryRepository) GetProductByFilter(ctx context.Context, filter models.ProductListFilter) (
	response []models.ProductListResponse, totalRows int64, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "GetProductByFilter")
	defer span.End()

	tx := repo.readDb.WithContext(ctx).Model(&entities.Product{}).
		Select(filter.Columns).
		Joins("LEFT JOIN categories ON products.category_id = categories.id").
		Joins("LEFT JOIN suppliers ON products.supplier_id = suppliers.id")

	if len(filter.Status) > 0 {
		tx = tx.Where("products.status IN ?", filter.Status)
	}

	if filter.CategoryID != "" {
		tx = tx.Where("products.category_id = ?", filter.CategoryID)
	}

	if filter.FromDate != nil {
		tx = tx.Where("products.added_date >= ?", filter.FromDate)
	}

	if filter.ToDate != nil {
		tx = tx.Where("products.added_date <= ?", filter.ToDate)
	}

	if filter.Limit > 0 {
		tx = tx.Limit(filter.Limit)
	}

	tx = tx.Count(&totalRows).Offset(filter.Offset).Order("products.added_date DESC")

	err := tx.Find(&response).Error
	if err != nil {
		errRes.IsSystemError = true
		errRes.Error = err
		return response, totalRows, errRes
	}

	return response, totalRows, errRes
}

func (repo *productQueryRepository) StatisticsProductPerCategory(ctx context.Context) (
	response []models.StatisticsProductPerCategory, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "StatisticsProductPerCategory")
	defer span.End()

	err := repo.readDb.WithContext(ctx).Raw(`
		WITH category_totals AS (
    		SELECT 
        		category_id, 
        		SUM(quantity) AS total_products_in_category
    		FROM products
    		GROUP BY category_id)
		SELECT 
			p.category_id,
    		p.id AS product_id,
    		p.name AS product_name,
    		p.quantity,
    		ROUND((p.quantity * 100.0) / ct.total_products_in_category, 2) AS percentage
		FROM products p
		JOIN category_totals ct ON p.category_id = ct.category_id;`).Scan(&response).Error

	if err != nil {
		errRes.IsSystemError = true
		errRes.Error = err
		return response, errRes
	}

	return response, errRes
}

func (repo *productQueryRepository) StatisticsProductPerSupplier(ctx context.Context) (
	response []models.StatisticsProductPerSupplier, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "StatisticsProductPerSupplier")
	defer span.End()

	err := repo.readDb.WithContext(ctx).Raw(`
		WITH supplier_totals AS (
    		SELECT 
        		supplier_id, 
        		SUM(quantity) AS total_products_in_supplier
    		FROM products
    		GROUP BY supplier_id)
		SELECT 
    		p.supplier_id,
    		p.id AS product_id,
    		p.name AS product_name,
    		p.quantity,
    		ROUND((p.quantity * 100.0) / ct.total_products_in_supplier, 2) AS percentage
		FROM products p
		JOIN supplier_totals ct ON p.supplier_id  = ct.supplier_id;`).Scan(&response).Error

	if err != nil {
		errRes.IsSystemError = true
		errRes.Error = err
		return response, errRes
	}

	return response, errRes
}
