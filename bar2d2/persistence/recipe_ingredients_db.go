package persistence

import (
	"bozosonparade/b2d2/bar2d2/dto"
	"database/sql"
	"fmt"
)

// GetAvailableRecipes returns a list of the available recipes.
func GetAvailableRecipes() ([]dto.Recipe, error) {
	aRecipes := []dto.Recipe{}
	db, err := sql.Open("sqlite3", *pDBPath)
	defer db.Close()

	checkErr(err)

	// query
	rows, err := db.Query("SELECT DISTINCT r.recipeId, r.title, r.description, r.image from recipes r, recipeIngredients ri, ingredients i, pumps p " +
		"WHERE r.recipeId=ri.recipeId AND i.ingredientId=ri.ingredientId AND i.pumpId=p.pumpId " +
		"ORDER BY r.title")
	checkErr(err)

	for rows.Next() {
		rep := dto.Recipe{}
		err = rows.Scan(&rep.ID, &rep.Title, &rep.Description, &rep.Image)
		checkErr(err)
		aRecipes = append(aRecipes, rep)
	}

	return aRecipes, err
}

// GetAvailableRecipesLight retrieves the list of available recipes without the images.
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
		err = rows.Scan(&rep.ID, &rep.Title, &rep.Description)
		checkErr(err)
		aRecipes = append(aRecipes, rep)
	}

	db.Close()
	return aRecipes, err
}

// GetRecipeWithIngredients returns all that's needed to make a particular recipe...the RecipeIngredients structure.
func GetRecipeWithIngredients(iID int) (dto.RecipeIngredients, error) {
	db, err := sql.Open("sqlite3", *pDBPath)

	defer db.Close()
	checkErr(err)

	// query
	qry := "SELECT DISTINCT r.recipeId, r.title, i.ingredientId, i.title, i.pumpId, p.gpioLine, ri.amount, p.ounceTiming from recipes r, recipeIngredients ri, ingredients i, pumps p " +
		"WHERE r.recipeId=? AND r.recipeId=ri.recipeId AND i.ingredientId=ri.ingredientId AND i.pumpId=p.pumpId " +
		"ORDER BY r.title"
	fmt.Printf("QUERY: %s\nID: %v\n", qry, iID)
	rows, err := db.Query(qry, iID)
	checkErr(err)

	var pRepIng *dto.RecipeIngredients
	for rows.Next() {
		var recipeID int
		var recipeTitle string
		var ingredientID int
		var ingredientTitle string
		var ingredientPumpID int
		var ingredientGPIOLine int16
		var ingredientAmount float32
		var ounceTiming int

		err = rows.Scan(&recipeID, &recipeTitle, &ingredientID, &ingredientTitle, &ingredientPumpID, &ingredientGPIOLine, &ingredientAmount, &ounceTiming)

		if pRepIng == nil {
			pRepIng = &dto.RecipeIngredients{}
			pRepIng.RecipeID = recipeID
			pRepIng.RecipeTitle = recipeTitle
		}
		if ingredientID > 0 && ingredientPumpID > 0 {
			ing := dto.SimpleIngredient{ingredientID, ingredientPumpID, ingredientGPIOLine, ingredientTitle, ingredientAmount, ounceTiming}
			pRepIng.Ingredients = append(pRepIng.Ingredients, ing)
		}
		checkErr(err)
	}

	return *pRepIng, err
}
