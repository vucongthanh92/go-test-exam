package product

import (
	"context"
	"fmt"

	"github.com/jung-kurt/gofpdf"
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

func (s *ProductImpl) GenProductListToPDF(ctx context.Context, req []models.ProductListResponse) (
	filePath, fileName string, errRes httpcommon.ErrorDTO) {

	_, span := tracing.StartSpanFromContext(ctx, "GenProductListToPDF")
	defer span.End()

	pdf := gofpdf.New("L", "mm", "A3", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Product List")
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)

	colWidths := []float64{40, 60, 30, 30, 50, 30, 50, 50, 40}
	cols := []string{
		"Product Reference",
		"Product Name",
		"Date Added",
		"Status",
		"Product Category",
		"Price",
		"Stock Location",
		"Supplier",
		"Available Quantity",
	}

	for i, col := range cols {
		pdf.CellFormat(colWidths[i], 10, col, "1", 0, "C", false, 0, "")
	}

	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 12)
	for _, product := range req {
		pdf.CellFormat(colWidths[0], 10, product.Reference, "1", 0, "L", false, 0, "")
		pdf.CellFormat(colWidths[1], 10, product.Name, "1", 0, "L", false, 0, "")
		// pdf.CellFormat(colWidths[2], 10, product.AddedDate.Format("2006-01-02"), "1", 0, "L", false, 0, "")
		pdf.CellFormat(colWidths[3], 10, product.Status, "1", 0, "L", false, 0, "")
		pdf.CellFormat(colWidths[4], 10, product.CategoryName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(colWidths[5], 10, fmt.Sprintf("%f", product.Price), "1", 0, "C", false, 0, "")
		pdf.CellFormat(colWidths[6], 10, product.StockCity, "1", 0, "L", false, 0, "")
		pdf.CellFormat(colWidths[7], 10, product.SupplierName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(colWidths[8], 10, fmt.Sprintf("%d", product.Quantity), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	filePath = "file/products.pdf"
	fileName = "products.pdf"

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		errRes.Error = err
		errRes.IsSystemError = true
		return filePath, fileName, errRes
	}

	return filePath, fileName, errRes
}

func (s *ProductImpl) StatisticsProductPerCategory(ctx context.Context) (
	response []models.StatisticsProductPerCategory, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "StatisticsProductPerCategory")
	defer span.End()

	response, errRes = s.productReadRepo.StatisticsProductPerCategory(ctx)
	return response, errRes
}

func (s *ProductImpl) StatisticsProductPerSupplier(ctx context.Context) (
	response []models.StatisticsProductPerSupplier, errRes httpcommon.ErrorDTO) {

	ctx, span := tracing.StartSpanFromContext(ctx, "StatisticsProductPerSupplier")
	defer span.End()

	response, errRes = s.productReadRepo.StatisticsProductPerSupplier(ctx)
	return response, errRes
}
