package services

import (
	"encoding/json"	
	"bozos.on.parade.com/b2d2/bar2d2/persistence"
	"fmt"
	"net/http"
	"strings"
)

func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	
	if (strings.EqualFold(r.Method, "GET")) {
		aRecipes, _ := persistence.GetAllRecipes()
		jsonRecipes, _ := json.Marshal(aRecipes)

		fmt.Fprintf(w, string(jsonRecipes))
		w.Header().Set("Content-Type", "application/json")
	}

}
