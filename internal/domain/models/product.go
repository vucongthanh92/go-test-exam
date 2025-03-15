package models

import (
	"github.com/vucongthanh92/go-test-exam/internal/domain/entities"
)

type ProductListRequest struct {
	Limit  int    `json:"limit" form:"limit"`
	Offset int    `json:"offset" form:"offset"`
	Status string `json:"status" form:"status"`
}

type ProductListFilter struct {
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
	Status  []string `json:"status"`
	Columns []string `json:"columns"`
}

type ProductListResponse struct {
	entities.Product
	CategoryName string `json:"category_name"`
	SupplierName string `json:"supplier_name"`
}
