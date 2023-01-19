#!/bin/bash

# Create Tables
psql findmyrecipe -f './server/database/queries/sql/createTables.sql'
psql findmyrecipe -c "select createTables();"

# Define Functions
psql findmyrecipe -f './server/database/queries/sql/dropTables.sql'
psql findmyrecipe -f './server/database/queries/sql/filterByTags.sql'
psql findmyrecipe -f './server/database/queries/sql/matchRecipes.sql'
psql findmyrecipe -f './server/database/queries/sql/userLogin.sql'
psql findmyrecipe -f './server/database/queries/sql/userSignup.sql'
psql findmyrecipe -f './server/database/queries/sql/autocompleteSearch.sql'

# Populate DB with Sample Data
psql findmyrecipe -c "\copy users from './demo/data/users_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy recipes from './demo/data/recipes_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy ingredients from './demo/data/ingredients_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy tags from './demo/data/tags_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy recipe_tag from './demo/data/recipe_tag_table.csv' DELIMITER ',' CSV HEADER;"
psql findmyrecipe -c "\copy recipe_ingredient from './demo/data/recipe_ingredient_table.csv' DELIMITER ',' CSV HEADER;"

# Change Serial Sequence
psql findmyrecipe -c "alter sequence users_id_seq restart with 100;"

