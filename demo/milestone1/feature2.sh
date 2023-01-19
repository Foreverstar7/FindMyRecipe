#!/bin/bash

# Autocomplete Search Ingredients
psql findmyrecipe -c "select autocompleteSearch('bac%');"

