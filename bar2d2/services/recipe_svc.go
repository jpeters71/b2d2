package services

import (
	"bozosonparade/b2d2/bar2d2/persistence"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// RecipeHandler handles requests sent to /svc/recipes
func RecipeHandler(w http.ResponseWriter, r *http.Request) {

	if strings.EqualFold(r.Method, "GET") {
		if strings.HasPrefix(r.URL.Path, "/svc/recipes/") {
			// Get the recipe ID
			id := r.URL.Path[len("/svc/recipes/"):]
			iID, _ := strconv.Atoi(id)
			rep, _ := persistence.GetRecipe(iID)
			jsonRecipe, _ := json.Marshal(rep)
			fmt.Fprintf(w, string(jsonRecipe))

			fmt.Printf("Recipe ID=%s\n", id)
		} else {
			aRecipes, _ := persistence.GetAllRecipes()
			jsonRecipes, _ := json.Marshal(aRecipes)

			fmt.Fprintf(w, string(jsonRecipes))
		}
		w.Header().Set("Content-Type", "application/json")
	}
}
