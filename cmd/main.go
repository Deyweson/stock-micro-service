package main

import (
	"github.com/deyweson/stock-micro-service/controller"
	db "github.com/deyweson/stock-micro-service/database"
	"github.com/deyweson/stock-micro-service/repository"
	"github.com/deyweson/stock-micro-service/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	conn, err := db.OpenConnection()
	if err != nil {
		panic(err.Error())
	}

	// repository layer
	stockRepository := repository.NewStockRepository(conn)

	// usecase layer
	stockUseCase := usecase.NewStockUseCase(stockRepository)

	// controller layer
	stockController := controller.NewStockController(stockUseCase)

	server := gin.Default()

	server.GET("/v1/stocks", stockController.GetStocks)
	server.GET("/v1/stocks/prod/:prod_id", stockController.GetStocksByProdId)
	server.GET("/v1/stocks/shop/:shop_id", stockController.GetStocksByProdId)

	server.Run(":8080")
}
