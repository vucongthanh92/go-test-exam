package v1

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(
	router *gin.Engine,
	productHandler *ProductHandler,
	categoryHandler *CategoryHandler,
	supplierHandler *SupplierHandler,
	distanceHandler *DistanceHandler,
) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/categories", categoryHandler.GetCategoryList)

		v1.GET("/products", productHandler.GetProductList)
		v1.GET("/product/gen-pdf", productHandler.GenProductListToPDF)

		v1.GET("/statistics/products-per-category", productHandler.StatisticsProductPerCategory)
		v1.GET("/statistics/products-per-supplier", productHandler.StatisticsProductPerSupplier)

		v1.GET("/distance/stock_city", distanceHandler.CalculateDistanceStockCity)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
