package repository

import (
	"fmt"

	"github.com/deyweson/stock-micro-service/model"
)

func (sr *StockRepository) GetStocks() ([]model.Stock, error) {
	query := "SELECT * FROM stock"
	rows, err := sr.connection.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		return []model.Stock{}, err
	}

	var stockList []model.Stock
	var stockObj model.Stock

	for rows.Next() {
		err = rows.Scan(
			&stockObj.Id,
			&stockObj.ShopId,
			&stockObj.ProdId,
			&stockObj.OperationType,
			&stockObj.Amount,
			&stockObj.Description,
			&stockObj.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Stock{}, err
		}

		stockList = append(stockList, stockObj)
	}
	return stockList, nil
}
