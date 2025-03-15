package v1

import (
	"github.com/vucongthanh92/go-test-exam/internal/application/supplier"
)

type SupplierHandler struct {
	supplierService supplier.SupllierService
}

func NewSupplierHandler(
	supplierService supplier.SupllierService,
) *SupplierHandler {
	return &SupplierHandler{
		supplierService: supplierService,
	}
}
