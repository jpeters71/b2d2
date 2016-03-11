package persistence

import (
	"bozosonparade/b2d2/bar2d2/dto"
	"database/sql"
)

// GetAllIngredients returns all ingredients in the DB.
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

// GetIngredient retrieves an ingredient by it's ID
func GetIngredient(id int) (dto.Ingredient, error) {
	var ing dto.Ingredient
	db, err := sql.Open("sqlite3", *pDBPath)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT ingredientId, title, description, brand, alcoholContent, pumpId FROM ingredients WHERE ingredientId=?", id)
	checkErr(err)

	for rows.Next() {
		ing = dto.Ingredient{}
		err = rows.Scan(&ing.Id, &ing.Title, &ing.Description, &ing.Brand, &ing.AlcoholContent, &ing.PumpId)
		checkErr(err)
	}
	db.Close()

	return ing, err
}
