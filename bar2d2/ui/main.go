package main

import (
	"fmt"
	"net/http"
	"bozos.on.parade.com/bar2d2/persistence"
	"bozos.on.parade.com/bar2d2/services"
)

func main() {
	persistence.InitPers()
	fmt.Printf("Hello, world.\n")
	http.HandleFunc("/ui/", UiHandler)
	http.HandleFunc("/svc/ingredients", services.IngredientHandler)
	http.HandleFunc("/svc/recipes", services.RecipeHandler)
	http.HandleFunc("/svc/available-recipes", services.AvailableHandler)
	http.ListenAndServe(":8080", nil)
}
