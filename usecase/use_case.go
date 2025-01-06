package usecase

import (
	"github.com/deyweson/stock-micro-service/model"
	"github.com/deyweson/stock-micro-service/repository"
)

type StockUseCase struct {
	StockRepository repository.StockRepository
}

func NewStockUseCase(repo repository.StockRepository) StockUseCase {
	return StockUseCase{
		StockRepository: repo,
	}
}

func (suc *StockUseCase) GetStockUseCase() ([]model.Stock, error) {
	return suc.StockRepository.GetStocks()
}

func (suc *StockUseCase) GetStockByProdIdUseCase(id string) ([]model.Stock, error) {
	return suc.StockRepository.GetStocksByProdId(id)
}

func (suc *StockUseCase) GetStockByShopIdUseCase(id string) ([]model.Stock, error) {
	return suc.StockRepository.GetStocksByShopId(id)
}
