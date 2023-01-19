#!/bin/bash

# Populate DB with Production Data
psql findmyrecipe -c "\copy users from './data/users_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy recipes from './data/recipes_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy ingredients from './data/ingredients_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy tags from './data/tags_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy recipe_tag from './data/recipe_tag_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy recipe_ingredient from './data/recipe_ingredient_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy user_favorites from './data/user_favorites_table.csv' DELIMITER ',' CSV HEADER;"

# Change Serial Sequence
psql findmyrecipe -c "alter sequence users_id_seq restart with 100;"
psql findmyrecipe -c "alter sequence recipes_id_seq restart with 231638;"
