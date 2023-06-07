package repository

import (
	"log"

	"github.com/mogueno/burrito-shop/models"
	"github.com/mogueno/burrito-shop/repository/dbqueries"
	database "github.com/mogueno/burrito-shop/utils"
)

func GetBurritos() []models.Burrito {
	rows, err := database.DB.Query(dbqueries.GET_BURRITOS)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var burritos []models.Burrito
	for rows.Next() {
		var burrito models.Burrito
		err := rows.Scan(&burrito.ID, &burrito.Name, &burrito.Size, &burrito.Price)
		if err != nil {
			log.Fatal(err)
		}

		burritos = append(burritos, burrito)
	}

	return burritos
}

func GetBurrito(burritoID uint) models.Burrito {
	var burrito models.Burrito
	err := database.DB.QueryRow(dbqueries.GET_BURRITO, burritoID).Scan(&burrito.ID, &burrito.Name, &burrito.Size, &burrito.Price)
	if err != nil {
		log.Fatal(err)
	}

	return burrito
}

func SaveBurrito(burrito models.Burrito) (int64, error) {
	res, err := database.DB.Exec(dbqueries.SAVE_BURRITO, burrito.Name, burrito.Size, burrito.Price)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}
