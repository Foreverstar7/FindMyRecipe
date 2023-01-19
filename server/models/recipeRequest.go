package models

type RecipeRequest struct {
  UserId int `json:"userId"`
  RecipeId int `json:"Id"`
  RecipeName string `json:"Name"`
  PrepMinutes int `json:"PrepMinutes"`
  Description string `json:"Description"`
  Ingredients []string `json:"Ingredients"`
  Nutrition []int `json:"Nutrition"`
  Tags []string `json:"Tags"`
  Steps []string `json:"Steps"`
}

