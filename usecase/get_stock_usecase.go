package usecase

import "github.com/deyweson/stock-micro-service/model"

func (suc *StockUseCase) GetStockUseCase() ([]model.Stock, error) {
	return suc.StockRepository.GetStocks()
}
