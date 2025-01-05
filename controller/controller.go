package controller

import (
	"net/http"

	"github.com/deyweson/stock-micro-service/usecase"
	"github.com/gin-gonic/gin"
)

type StockController struct {
	StockUseCase usecase.StockUseCase
}

func NewStockController(StockUseCase usecase.StockUseCase) StockController {
	return StockController{
		StockUseCase: StockUseCase,
	}
}

func (sc *StockController) GetStocks(c *gin.Context) {
	stocks, err := sc.StockUseCase.GetStockUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stocks)
}
