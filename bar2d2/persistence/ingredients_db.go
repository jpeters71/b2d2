package persistence

import (
	"bozosonparade/b2d2/bar2d2/dto"
	"database/sql"
)

func GetAllIngredients() ([]dto.Ingredient, error) {

	aIngs := []dto.Ingredient{}
	db, err := sql.Open("sqlite3", *pDBPath)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT ingredientId, title, description, brand, alcoholContent, pumpId FROM ingredients")
	checkErr(err)

	for rows.Next() {
		ing := dto.Ingredient{}
		err = rows.Scan(&ing.Id, &ing.Title, &ing.Description, &ing.Brand, &ing.AlcoholContent, &ing.PumpId)
		checkErr(err)
		aIngs = append(aIngs, ing)
	}
	db.Close()

	return aIngs, err
}
