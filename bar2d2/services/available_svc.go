package services

import (
	"bozosonparade/b2d2/bar2d2/persistence"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AvailableHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("AvailableHandler called\n")
	if strings.EqualFold(r.Method, "GET") {
		aRecipes, _ := persistence.GetAvailableRecipes()
		jsonRecipes, _ := json.Marshal(aRecipes)

		fmt.Fprintf(w, string(jsonRecipes))
		w.Header().Set("Content-Type", "application/json")
	}
}

func AvailableHandlerLight(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("AvailableHandlerLight called\n")

	if strings.EqualFold(r.Method, "GET") {
		aRecipes, _ := persistence.GetAvailableRecipesLight()
		jsonRecipes, _ := json.Marshal(aRecipes)

		fmt.Fprintf(w, string(jsonRecipes))
		w.Header().Set("Content-Type", "application/json")
	}
}
