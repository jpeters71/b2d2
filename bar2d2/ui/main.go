package main

import (
	"bozosonparade/b2d2/bar2d2/persistence"
	"bozosonparade/b2d2/bar2d2/services"
	"flag"
	"fmt"
	"net/http"
)

var dbPath string
var docPath string

func main() {
	pRootPath := flag.String("root", ".", "root path to Bar2D2.")
	flag.Parse()

	dbPath = fmt.Sprintf("%s/b2d2db.sqlite3", *pRootPath)
	docPath = fmt.Sprintf("%s/ui/", *pRootPath)

	fmt.Printf("Bar2D2 coming online.\n")
	fmt.Printf("DB Path: %s\n", dbPath)
	fmt.Printf("Doc Path: %s\n", docPath)

	persistence.InitPers(&dbPath)
	fmt.Printf("Hello, world.\n")
	http.HandleFunc("/ui/", UIHandler)
	http.HandleFunc("/svc/ingredients", services.IngredientHandler)
	http.HandleFunc("/svc/recipes", services.RecipeHandler)
	http.HandleFunc("/svc/recipes/", services.RecipeHandler)
	http.HandleFunc("/svc/available-recipes", services.AvailableHandler)
	http.HandleFunc("/svc/available-recipes-light", services.AvailableHandlerLight)
	fmt.Printf("Starting service on port 8080")
	http.ListenAndServe(":8080", nil)
}
