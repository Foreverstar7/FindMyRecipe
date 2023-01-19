package util

import (
	"os"
  "log"
  "fmt"
	"strings"
  "golang.org/x/crypto/bcrypt"
)

func ReadQueryFile(f string) string {
  var q, err = os.ReadFile("./server/database/queries/" + f)
  if err != nil {
    log.Fatal(err)
  }
  return string(q)
}

func ArrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func ArrayStringToString(a []string, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), "' '", "'" + delim + "'", -1), "[]")
}

func ArrayContains(a []int, c int) bool {
  for _, s := range a {
    if (s == c) {
      return true
    }
  }
  return false
}

func HashPassword(password string) string {
  bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14);
  return string(bytes);
}

func CheckPassword(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password));
  return err == nil;
}


// used for bulk inserting to recipe_ingredient and recipe_tag
func PrepareBulkInsert(query string, placeholder string, numColumn int, values [][]int) (string, []interface{}) {
  query = strings.Replace(query, placeholder, "%s", 1)
  valueStrings := make([]string, 0, len(values))
  valueArgs := make([]interface{}, 0, len(values) * numColumn)
  valuePlaceholder := "("
  for i := 0; i < numColumn; i++ {
    if i != 0 {
      valuePlaceholder += ", "
    } 
    valuePlaceholder += "?"
  }
  valuePlaceholder += ")"
  for _, value := range values {
      valueStrings = append(valueStrings, valuePlaceholder)
      for i := 0; i < numColumn; i++ {
        valueArgs = append(valueArgs, value[i])
      }
  }
  stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))
  // fmt.Println(stmt, valueStrings, valueArgs, valuePlaceholder)
  return stmt, valueArgs
}
