package repository

import (
	"fmt"

	"github.com/deyweson/stock-micro-service/model"
)

func (sr *StockRepository) GetStocksByShopId(id string) ([]model.Stock, error) {
	query := "SELECT * FROM stock WHERE shop_id = $1"
	rows, err := sr.connection.Query(query, id)

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

	if len(stockList) < 1 {
		return []model.Stock{}, nil
	}

	return stockList, nil
}
