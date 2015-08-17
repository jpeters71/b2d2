package main

import (
	"flag"
	"fmt"
	"net/http"
	"bozos.on.parade.com/bar2d2/persistence"
	"bozos.on.parade.com/bar2d2/services"
)
var pDBPath *string

func main() {
	pDBPath = flag.String("db", "./b2d2db.sqlite3", "path to the SQLite3 Bar2D2 DB.")

	fmt.Printf("Bar2D2 coming online.\nDB Path: %s\n", *pDBPath)

	persistence.InitPers(pDBPath)
	fmt.Printf("Hello, world.\n")
	http.HandleFunc("/ui/", UiHandler)
	http.HandleFunc("/svc/ingredients", services.IngredientHandler)
	http.HandleFunc("/svc/recipes", services.RecipeHandler)
	http.HandleFunc("/svc/available-recipes", services.AvailableHandler)
	http.HandleFunc("/svc/available-recipes-light", services.AvailableHandlerLight)
	http.ListenAndServe(":8080", nil)
}
