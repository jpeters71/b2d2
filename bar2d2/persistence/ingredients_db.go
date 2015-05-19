package persistence

import (
	"database/sql"
	"bozos.on.parade.com/bar2d2/dto"
)


func GetAllIngredients() ([]dto.Ingredient, error) {

	aIngs := []dto.Ingredient{}
	db, err := sql.Open("sqlite3", "/home/johnp/b2d2/b2d2db.sqlite3")
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

