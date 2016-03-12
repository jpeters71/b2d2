package services

import (
	"bozosonparade/b2d2/bar2d2/dto"
	"bozosonparade/b2d2/bar2d2/gpio_rpi"
	"bozosonparade/b2d2/bar2d2/persistence"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
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
	} else if strings.EqualFold(r.Method, "POST") {
		if strings.HasPrefix(r.URL.Path, "/svc/recipes/") {
			// Get the recipe ID
			id := r.URL.Path[len("/svc/recipes/"):]
			iID, _ := strconv.Atoi(id)
			makeDrink(w, iID)
		}
	}
}

func makeDrink(w http.ResponseWriter, iID int) {
	var wg sync.WaitGroup

	if iID <= 0 {
		http.Error(w, "Invalid recipe ID", 400)
		return
	}
	// Pull up the recipe definition from the DB
	if repIng, err := persistence.GetRecipeWithIngredients(iID); err != nil {
		http.Error(w, "DB error occurred making drink: "+err.Error(), 500)
	} else {
		// Make drink! (this is where the magic happens) Here's the plan:
		// * Start a separate go proc for each ingredient
		// * Go procs will turn on the pump and wait the appropriate time
		// * Main thread will wait for all threads to be completed before responding

		for _, si := range repIng.Ingredients {
			fmt.Printf("Add!\n")
			wg.Add(1)
			go processIngredient(&wg, si)
		}
		wg.Wait()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{'status':'OK'}")
	}
}

func processIngredient(wg *sync.WaitGroup, si dto.SimpleIngredient) {
	defer trace("processIngredient - " + strconv.Itoa(int(si.GPIOLine)))()

	// Calculate sleep time
	iSleepTime := time.Duration(math.Ceil(float64(si.OunceTiming) * float64(si.Amount)))

	gpio.SetMode(si.GPIOLine, gpio.Write)
	gpio.SetHigh(si.GPIOLine)
	time.Sleep(iSleepTime * time.Millisecond)
	gpio.SetLow(si.GPIOLine)
	wg.Done()
}

func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("%s started at: %s\n", msg, start.Local())
	return func() {
		end := time.Now()
		fmt.Printf("%s finished at: %s; elapsed time: %s\n", msg, end.Local(), end.Sub(start))
	}
}
