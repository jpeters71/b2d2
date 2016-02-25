package persistence

import (
	"bozosonparade/b2d2/bar2d2/dto"
	"database/sql"
)

func GetAvailableRecipes() ([]dto.Recipe, error) {
	aRecipes := []dto.Recipe{}
	db, err := sql.Open("sqlite3", *pDBPath)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT DISTINCT r.recipeId, r.title, r.description, r.image from recipes r, recipeIngredients ri, ingredients i, pumps p " +
		"WHERE r.recipeId=ri.recipeId AND i.ingredientId=ri.ingredientId AND i.pumpId=p.pumpId " +
		"ORDER BY r.title")
	checkErr(err)

	for rows.Next() {
		rep := dto.Recipe{}
		err = rows.Scan(&rep.Id, &rep.Title, &rep.Description, &rep.Image)
		checkErr(err)
		aRecipes = append(aRecipes, rep)
	}

	db.Close()
	return aRecipes, err
}

func GetAvailableRecipesLight() ([]dto.Recipe, error) {

	aRecipes := []dto.Recipe{}
	db, err := sql.Open("sqlite3", *pDBPath)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT DISTINCT r.recipeId, r.title, r.description from recipes r, recipeIngredients ri, ingredients i, pumps p " +
		"WHERE r.recipeId=ri.recipeId AND i.ingredientId=ri.ingredientId AND i.pumpId=p.pumpId " +
		"ORDER BY r.title")
	checkErr(err)

	for rows.Next() {
		rep := dto.Recipe{}
		err = rows.Scan(&rep.Id, &rep.Title, &rep.Description)
		checkErr(err)
		aRecipes = append(aRecipes, rep)
	}

	db.Close()
	return aRecipes, err
}
