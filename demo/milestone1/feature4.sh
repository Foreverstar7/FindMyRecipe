#!/bin/bash

# Match Recipes based on Ingredients 
psql findmyrecipe -c "select filterByTags(124);"

