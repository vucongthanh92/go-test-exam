package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductListRequest struct {
	Limit      int    `json:"limit" form:"limit"`
	Offset     int    `json:"offset" form:"offset"`
	Status     string `json:"status" form:"status"`
	CategoryID string `json:"category_id" form:"category_id"`
	FromDate   string `json:"from_date" form:"from_date"`
	ToDate     string `json:"to_date" form:"to_date"`
}

type ProductListFilter struct {
	Limit      int        `json:"limit"`
	Offset     int        `json:"offset"`
	Status     []string   `json:"status"`
	Columns    []string   `json:"columns"`
	CategoryID string     `json:"category_id"`
	FromDate   *time.Time `json:"from_date" form:"from_date"`
	ToDate     *time.Time `json:"to_date" form:"to_date"`
}

type ProductListResponse struct {
	ID           uuid.UUID `json:"id"`
	Reference    string    `json:"reference"`
	Name         string    `json:"name"`
	AddedDate    string    `json:"added_date"`
	Status       string    `json:"status"`
	CategoryID   uuid.UUID `json:"category_id"`
	Price        float64   `json:"price"`
	StockCity    string    `json:"stock_city"`
	SupplierID   uuid.UUID `json:"supplier_id"`
	Quantity     int       `json:"quantity"`
	CategoryName string    `json:"category_name"`
	SupplierName string    `json:"supplier_name"`
}

type StatisticsProductPerCategory struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
	ProductID    uuid.UUID `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	Percentage   float64   `json:"percentage"`
}

type StatisticsProductPerSupplier struct {
	SupplierID   uuid.UUID `json:"supplier_id"`
	SupplierName string    `json:"supplier_name"`
	ProductID    uuid.UUID `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	Percentage   float64   `json:"percentage"`
}
