package persistence

import (
	"database/sql"
	"bozos.on.parade.com/bar2d2/dto"
)


func GetAllRecipes() ([]dto.Recipe, error) {

	aRecipes := []dto.Recipe{}
	db, err := sql.Open("sqlite3", "/home/johnp/b2d2/b2d2db.sqlite3")
	checkErr(err)

    // query
    rows, err := db.Query("SELECT recipeId, title, description, image FROM recipes")
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

