package dto

// Ingredient structure holds data about the components of drinks.
type Ingredient struct {
	ID             int
	Title          string
	Description    string
	Brand          string
	AlcoholContent int
	PumpID         int
}

// SimpleIngredient is a simplier form of Ingredient designed to hold just what's needed
// to actually make the drink
type SimpleIngredient struct {
	ID          int
	PumpID      int
	GPIOLine    int16
	Title       string
	Amount      float32
	OunceTiming int
}
