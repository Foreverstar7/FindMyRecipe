package database

import (
	"database/sql"
	m "main/server/models"
	"main/server/util"
)

func GetIngredients() []m.Ingredient {
  var ingredients []m.Ingredient
  db.Raw("select * from ingredients limit 10").Scan(&ingredients)
  return ingredients
}

func GetIngredientById(id int) m.Ingredient {
  var ingredient m.Ingredient
  db.Raw("select * from ingredients where id = @id limit 1", sql.Named("id", id)).Scan(&ingredient)
  return ingredient
}

func GetIngredientByName(name string) m.Ingredient {
  var ingredient m.Ingredient
  db.Raw("select * from ingredients where name = @name limit 1", sql.Named("name", name)).Scan(&ingredient)
  return ingredient
}

func GetIngredientsByRecipeId(id int) []m.Ingredient {
  var ingredients []m.Ingredient
  db.Raw("select * from ingredients where id in (select ingredient_id from recipe_ingredient where recipe_id = @id)", sql.Named("id", id)).Scan(&ingredients)
  return ingredients
}

func GetAutocompleteIngredient(substring string) []m.Ingredient {
  var ingredients []m.Ingredient
  q := util.ReadQueryFile("AutocompleteIngredient")

  prettyQuery := "%" + substring + "%"
  db.Raw(q, sql.Named("substring", prettyQuery)).Scan(&ingredients)
  if len(ingredients) == 0 {
    return nil
  }
  return ingredients
}

