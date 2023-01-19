package database

import (
	"database/sql"
	m "main/server/models"
	"main/server/util"
	"math/rand"
	"time"
  "errors"
  "fmt"
)

func GetRecipes() []m.Recipe {
  var recipes []m.Recipe
  db.Raw("select * from recipes limit 10").Scan(&recipes)
  return recipes
}

func GetRecipesPage(page int) []m.Recipe {
  if page < 0 {
    return nil
  }
  var recipes []m.Recipe
  db.Raw("select * from recipes order by id limit 6 offset @page_offset", sql.Named("page_offset", (page - 1) * 6)).Scan(&recipes)
  return recipes
}

func GetRecipeById(id int) (m.Recipe, error) {
  var recipe m.Recipe
  db.Raw("select * from recipes where id = @id limit 1", sql.Named("id", id)).Scan(&recipe)
  if (recipe.Id == 0) {
    return recipe, errors.New("Recipe not found")
  }
  return recipe, nil
}

func GetRecipesByTags(tagId int) []m.Recipe {
  var recipesByTags []m.Recipe
  q := util.ReadQueryFile("FilterByTags")

  db.Raw(q, sql.Named("tag_id", tagId)).Scan(&recipesByTags)
  if len(recipesByTags) == 0 {
    return nil
  }
  return recipesByTags
}

func GetRecipesByManyTags(tagIds []int) []m.Recipe {
  var recipesByTags []m.Recipe
  q := util.ReadQueryFile("FilterByManyTags")

  prettyQuery := "{" + util.ArrayToString(tagIds, ",") + "}"

  db.Raw(q, sql.Named("tag_ids", prettyQuery), sql.Named("tag_count", len(tagIds))).Scan(&recipesByTags)
  if len(recipesByTags) == 0 {
    return nil
  }
  return recipesByTags
}

func GetMatchingRecipesByIngredients(ingredientIds []int) []m.Recipe {
  var recipes []m.Recipe
  q := util.ReadQueryFile("MatchRecipesByIngredients")

  prettyQuery := "{" + util.ArrayToString(ingredientIds, ",") + "}"

  db.Raw(q, sql.Named("ingredient_ids", prettyQuery)).Scan(&recipes)
  if len(recipes) == 0 {
    return nil
  }
  return recipes
}

func GetMatchingRecipes(ingredientIds []int, tagIds []int) []m.Recipe {
  var recipes []m.Recipe
  q := util.ReadQueryFile("MatchRecipes")

  ingredients := "{" + util.ArrayToString(ingredientIds, ",") + "}"
  tags := "{" + util.ArrayToString(tagIds, ",")+ "}"

  db.Raw(q, sql.Named("ingredient_ids", ingredients), sql.Named("tag_ids", tags),
    sql.Named("tag_count", len(tagIds))).Scan(&recipes)
  if len(recipes) == 0 {
    return nil
  }
  return recipes
}

func GetRecipesByAuthor(author int) []m.Recipe {
  var recipes []m.Recipe

  db.Raw("SELECT * FROM recipes WHERE author = @id", sql.Named("id", author)).Scan(&recipes)
  if len(recipes) == 0 {
    return nil
  }
  return recipes
}

func WeeklyFavorites() []m.Recipe {
  var recipes []m.Recipe
  var q = util.ReadQueryFile("WeeklyFavorites")

  current := time.Now()
  offset := (int(time.Monday) - int(current.Weekday()) - 7) % 7
  startWeek := current.Add(time.Duration(offset*24) * time.Hour).Add(time.Duration(-current.Hour()) * time.Hour)

  db.Raw(q, sql.Named("week_timestamp", startWeek)).Scan(&recipes)
  if len(recipes) == 0 {
    return nil
  }
  return recipes
}

func GetRecommendation(userId int) []m.Recipe {
  var recipes []m.Recipe
  var q = util.ReadQueryFile("RecommendRecipes")

  db.Raw(q, sql.Named("user_id", userId)).Scan(&recipes)
  return recipes
}

func GetSimilarRecipes(ingredientIds []int) []m.Recipe {
  var recipes []m.Recipe
  q := util.ReadQueryFile("SimilarRecipes")

  ingredients := "{" + util.ArrayToString(ingredientIds, ",") + "}"
  staples := []int{6557,14832,9454,8189,8223,2237,4621,8792,13436,10905,409,7206,1616,13150,14809,1097,3282,14020,12149,9604}
  randomIndex := rand.Intn(len(staples))
  appendedIngredientsId := append(ingredientIds, staples[randomIndex])
  appendedIngrdients := "{" + util.ArrayToString(appendedIngredientsId, ",") + "}"

  db.Raw(q, sql.Named("fridge1", appendedIngrdients), sql.Named("fridge2", ingredients)).Scan(&recipes)
  if len(recipes) == 0 {
    r := rand.Intn(200000)
    db.Raw("SELECT * FROM recipes WHERE id = @id", sql.Named("id", r)).Scan(&recipes)
    return recipes
  }
  return recipes
}


  // TODO:
  // insert into recipes values () returning id
  // get the recipe id
  // get ingredient id by name
  // insert into recipe_ingredient
  // get tag id by name
  // insert into recipe_tag
  // make sure all is all or nothing
func AddNewRecipe(recipe m.RecipeRequest) bool {
    // name text,
    // description text,
    // prep_minutes int,
    // author int,
    // steps text[],
    // nutrition float8[],

  // UserId int `json:"userId"`
  // RecipeId int `json:"Id"`
  // RecipeName string `json:"Name"`
  // PrepMinutes int `json:"PrepMinutes"`
  // Description string `json:"Description"`
  // Ingredients []string `json:"Ingredients"`
  // Nutrition []int `json:"Nutrition"`
  // Tags []string `json:"Tags"`
  // Steps []string `json:"Steps"`
  insertRecipe := "INSERT INTO recipes(name, description, prep_minutes, author, steps, nutrition) VALUES(@1, @2, @3, @4, cast(@5 as text[]), cast(@6 as float8[])) RETURNING id;"
  recipeId := 0
  recipeSteps := "{" + util.ArrayStringToString(recipe.Steps, ",") + "}"
  recipeNutrition := "{" + util.ArrayToString(recipe.Nutrition, ",") + "}"
  r1 := db.Raw(insertRecipe, sql.Named("1", recipe.RecipeName), sql.Named("2", recipe.Description), 
  sql.Named("3", recipe.PrepMinutes), sql.Named("4", recipe.UserId), sql.Named("5", recipeSteps), 
  sql.Named("6", recipeNutrition)).Scan(&recipeId)
    //  (cast(@tag_ids as int[]))
  if (r1.Error != nil || recipeId == 0) {
    return false
  }

  // insert recipe_ingredient
  ingredientIds := make([]int, 0)
  for _, ingredientName := range recipe.Ingredients {
    ingredientId := 0
    db.Raw("SELECT id FROM ingredients WHERE name = @ingredientName", sql.Named("ingredientName", ingredientName)).Scan(&ingredientId)
    ingredientIds = append(ingredientIds, ingredientId)
  }
  insertRecipeIngredient := "INSERT INTO recipe_ingredient(recipe_id, ingredient_id) VALUES @values;"
  riValues := make([][]int, 0)
  for _, ingredientId := range ingredientIds {
    value := []int{recipeId, ingredientId}
    riValues = append(riValues, [][]int{value}...)
  }
  stmt2, valueArgs2 := util.PrepareBulkInsert(insertRecipeIngredient, "@values", 2, riValues)
  r2 := db.Exec(stmt2, valueArgs2...)
  if r2.Error != nil {
    // rollback insert recipe
    rerr := db.Exec("DELETE FROM recipes WHERE id = @recipeId", sql.Named("recipeId", recipeId))
    if (rerr.Error != nil) {
      fmt.Println("Failed to rollback delete from recipes where id ", recipeId)
      return false
    }
    return false
  }

  // insert recipe_tag
  tagIds := make([]int, 0)
  for _, tagName := range recipe.Tags {
    tagId := 0
    db.Raw("SELECT id FROM tags WHERE name = @tagName", sql.Named("tagName", tagName)).Scan(&tagId)
    tagIds = append(tagIds, tagId)
  }
  insertRecipeTag := "INSERT INTO recipe_tag(recipe_id, tag_id) VALUES @values;"
  rtValues := make([][]int, 0)
  for _, tagId:= range tagIds {
    value := []int{recipeId, tagId}
    rtValues = append(rtValues, [][]int{value}...)
  }
  stmt3, valueArgs3 := util.PrepareBulkInsert(insertRecipeTag, "@values", 2, rtValues)
  r3 := db.Exec(stmt3, valueArgs3...)
  if r3.Error != nil {
    // rollback insert recipe and recipe_tag (delete recipe auto cascades recipe_tag by table def)
    rerr := db.Exec("DELETE FROM recipes WHERE id = @recipeId", sql.Named("recipeId", recipeId))
    if (rerr.Error != nil) {
      fmt.Println("Failed to rollback delete from recipes where id ", recipeId)
      return false
    }
    return false
  }
  return true
}
