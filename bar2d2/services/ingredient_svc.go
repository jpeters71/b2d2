package services

import (
	"encoding/json"	
	"bozos.on.parade.com/bar2d2/persistence"
	"fmt"
	"net/http"
	"strings"
)

func IngredientHandler(w http.ResponseWriter, r *http.Request) {
	
	if (strings.EqualFold(r.Method, "GET")) {
		aIngs, _ := persistence.GetAllIngredients()
		jsonIngs, _ := json.Marshal(aIngs)

		fmt.Fprintf(w, string(jsonIngs))
		w.Header().Set("Content-Type", "application/json")
	}

}