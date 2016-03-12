package dto

// RecipeIngredients structure contains basic information about the recipe along
// with the ingredients needed to make it.  Mainly used for actually creating a
// drink.
type RecipeIngredients struct {
	RecipeID    int
	RecipeTitle string
	Ingredients []SimpleIngredient
}
