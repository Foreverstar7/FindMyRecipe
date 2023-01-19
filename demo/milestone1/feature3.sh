#!/bin/bash

# Match Recipes based on Ingredients 
psql findmyrecipe -c "select matchRecipes('{38,102,29,33,91,82,39,10}');"

