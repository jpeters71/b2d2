package services

import (
	"bozosonparade/b2d2/bar2d2/persistence"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// IngredientHandler handles requests to /svc/ingredients
func IngredientHandler(w http.ResponseWriter, r *http.Request) {

	if strings.EqualFold(r.Method, "GET") {
		if strings.HasPrefix(r.URL.Path, "/svc/ingredients/") {
			// Get the recipe ID
			id := r.URL.Path[len("/svc/ingredients/"):]
			iID, _ := strconv.Atoi(id)
			rep, _ := persistence.GetIngredient(iID)
			jsonIng, _ := json.Marshal(rep)
			fmt.Fprintf(w, string(jsonIng))

			fmt.Printf("Ingredient ID=%s\n", id)
		} else {
			aIngs, _ := persistence.GetAllIngredients()
			jsonIngs, _ := json.Marshal(aIngs)

			fmt.Fprintf(w, string(jsonIngs))
		}
		w.Header().Set("Content-Type", "application/json")
	}
}
