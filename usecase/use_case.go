package usecase

import "github.com/deyweson/stock-micro-service/repository"

type StockUseCase struct {
	StockRepository repository.StockRepository
}

func NewStockUseCase(repo repository.StockRepository) StockUseCase {
	return StockUseCase{
		StockRepository: repo,
	}
}
