package database

import (
	"database/sql"
	m "main/server/models"
	"main/server/util"
)

func GetAutocompleteTag(substring string) []m.Tag {
  var tags []m.Tag
  q := util.ReadQueryFile("AutocompleteTag")

  prettyQuery := substring + "%"
  db.Raw(q, sql.Named("substring", prettyQuery)).Scan(&tags)
  if len(tags) == 0 {
    return nil
  }
  return tags
}

func GetTagById(id int) m.Tag {
  var tag m.Tag
  db.Raw("select * from tags where id = @tag_id limit 1", sql.Named("tag_id", id)).Scan(&tag)
  return tag;
}

func GetTagByName(name string) m.Tag {
  var tag m.Tag
  db.Raw("select * from tags where name = @tag_name limit 1", sql.Named("tag_name", name)).Scan(&tag)
  return tag;
}

func GetAllTags() []m.Tag {
  var tags []m.Tag
  db.Raw("SELECT * FROM tags").Scan(&tags)
  return tags
}

func GetTagsOfRecipe(recipeId int) []m.Tag {
  var tags []m.Tag
  db.Raw("SELECT * FROM tags WHERE id IN (SELECT tag_id FROM recipe_tag WHERE recipe_id = @recipeId)", sql.Named("recipeId", recipeId)).Scan(&tags)
  return tags
}

