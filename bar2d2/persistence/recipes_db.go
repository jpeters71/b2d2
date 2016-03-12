package persistence

import (
	"bozosonparade/b2d2/bar2d2/dto"
	"database/sql"
)

// GetAllRecipes returns all recipes
func GetAllRecipes() ([]dto.Recipe, error) {

	aRecipes := []dto.Recipe{}
	db, err := sql.Open("sqlite3", *pDBPath)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT recipeId, title, description, image FROM recipes ORDER BY title")
	checkErr(err)

	for rows.Next() {
		rep := dto.Recipe{}
		err = rows.Scan(&rep.ID, &rep.Title, &rep.Description, &rep.Image)
		checkErr(err)
		aRecipes = append(aRecipes, rep)
	}
	db.Close()

	return aRecipes, err
}

// GetRecipe retrieves a recipe by it's ID
func GetRecipe(id int) (dto.Recipe, error) {
	var rep dto.Recipe
	db, err := sql.Open("sqlite3", *pDBPath)
	checkErr(err)

	// query
	rows, err := db.Query("SELECT recipeId, title, description, image FROM recipes WHERE recipeId=?", id)
	checkErr(err)

	for rows.Next() {
		rep = dto.Recipe{}
		err = rows.Scan(&rep.ID, &rep.Title, &rep.Description, &rep.Image)
		checkErr(err)
	}
	db.Close()

	return rep, err
}
