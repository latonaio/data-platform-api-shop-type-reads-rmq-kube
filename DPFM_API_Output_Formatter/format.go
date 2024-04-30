package dpfm_api_output_formatter

import (
	"data-platform-api-shop-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToShopType(rows *sql.Rows) (*[]ShopType, error) {
	defer rows.Close()
	shopType := make([]ShopType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ShopType{}

		err := rows.Scan(
			&pm.ShopType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &shopType, nil
		}

		data := pm
		shopType = append(shopType, ShopType{
			ShopType:				data.ShopType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &shopType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.ShopType,
			&pm.Language,
			&pm.ShopTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			ShopType:     			data.ShopType,
			Language:          		data.Language,
			ShopTypeName:			data.ShopTypeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
