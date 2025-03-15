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
) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/products", productHandler.GetProductList)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
