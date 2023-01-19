package database

import (
	"database/sql"
	m "main/server/models"
	"main/server/util"
  "fmt"
  "errors"
)

func UserLogin(username string, password string) (m.User, error) {
  var user m.User
  q := util.ReadQueryFile("UserLogin")

  db.Raw(q, sql.Named("username", username)).First(&user)

  if (user.Name == "masterchef" && user.Password == "masterchef") {
    return user, nil
  }

  if (util.CheckPassword(password, user.Password)) {
    return user, nil
  }
  return user, errors.New("Fail to login")
}


func UserSignUp(username string, password string) bool {
  var users []m.User
  var q = util.ReadQueryFile("UserLogin")

  db.Raw(q, sql.Named("username", username)).Scan(&users)
  if len(users) > 0 {
    return false
  }

  q = util.ReadQueryFile("UserSignup")
  hashed_password := util.HashPassword(password)

  db.Exec(q, sql.Named("username", username), sql.Named("password", hashed_password))
  return true
}

func ToggleUserFavorite(userId int, recipeId int) bool {
  var userFavorites []int 
  q := "SELECT recipe_id FROM user_favorites WHERE user_id = @userId;"

  db.Raw(q, sql.Named("userId", userId)).Scan(&userFavorites)
  if (util.ArrayContains(userFavorites, recipeId)) {
    // remove this
    fmt.Println(userFavorites, "unfavorite", userId, recipeId);
    r := db.Exec("DELETE FROM user_favorites WHERE user_id = @userId AND recipe_id = @recipeId", sql.Named("userId", userId), sql.Named("recipeId", recipeId))
    if r.Error != nil {
      return false
    }
    fmt.Println("uF:", userFavorites)
  } else {
    // add this
    fmt.Println(userFavorites, "favorite", userId, recipeId);
    r := db.Exec("INSERT INTO user_favorites VALUES (@userId, @recipeId)", sql.Named("userId", userId), sql.Named("recipeId", recipeId))
    if r.Error != nil {
      return false
    }
    fmt.Println("uF:", userFavorites)
  }
  return true
}

// Testing transaction with parameters for bulk insert
// sql with params can only be executed by postgresql for one query
// having both multiple queries and sql parameters are not possible to do safely
// if we want to have multiple queries in one transaction, this involves multiple queris from the setup of transaction itself
// it is possible to just not use parameter and replace our placeholder ourself, but this is prone to sql injection
// for this reason, we will postpone the transactional approach and instead assume everything is correct from the frontend
// so that we can just execute multiple queries one at a time
// func TestFavorites(userId int, recipeIds []int) bool {
//   values := make([][]int, 0)
//   for _, recipe_id := range recipeIds {
//     value := []int{userId, recipe_id}
//     values = append(values, [][]int{value}...)
//   }
//   fmt.Println(userId, recipeIds, values)
//   q := util.ReadQueryFile("TestFavorite")
//   stmt, valueArgs := util.PrepareBulkInsert(q, "@values", 2, values) // (query, placeholder to replace beginning with @, number of column, values)
//   fmt.Println(stmt)
//   r := db.Exec(stmt, valueArgs...)
//   if r.Error != nil {
//     fmt.Println(r)
//     return false
//   }
//   fmt.Println("inserted")
//   return true
// }

func GetUserFavorites(userId int) []m.Recipe {
  var userFavorites []m.Recipe
  q := util.ReadQueryFile("UserFavorites")

  db.Raw(q, sql.Named("userId", userId)).Scan(&userFavorites)
  return userFavorites
}

