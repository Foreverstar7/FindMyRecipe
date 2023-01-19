#!/usr/bin/env python3

import sys
import ast
import pandas as pd

def parse_list(s):
    return str([x.replace("{", "").replace("}", "").replace(",", ";").replace("'", "").replace('"', ' ') for x in ast.literal_eval(s)])

def to_curly(s):
    return "{" + s[1:len(s)-1] + "}"

def main():
    if len(sys.argv) != 2:
        print("Usage: parse_recipe_csv.py CSV_FILE")
        exit(1)

    data = pd.read_csv(sys.argv[1])

    # Produce .csv for users table
    user_id = [1]
    name = ["masterchef"]
    password = ["masterchef"]
    users = pd.DataFrame({"id": user_id, "name": name, "password": password})
    users.to_csv("./users_table.csv", index=False)

    # Produce .csv for recipes table
    recipes_id = range(1,data.shape[0]+1)
    name = data["name"]
    description = data["description"]
    prep_minutes = data["minutes"]
    author = [1]*data.shape[0]
    steps = [to_curly(parse_list(i)) for i in data["steps"]]
    nutrition = [to_curly(i) for i in data["nutrition"]]
    recipes = pd.DataFrame({
        "id": recipes_id,
        "name": name,
        "description": description,
        "prep_minutes": prep_minutes,
        "author": author,
        "steps": steps,
        "nutrition": nutrition
    })
    recipes.to_csv("./recipes_table.csv", index=False)

    # Produce .csv for tags_name table
    tags_name = set()
    for t in data["tags"]:
        for item in ast.literal_eval(t):
            tags_name.add(item)
    tags_name = list(tags_name)
    tags_id = range(1,len(tags_name)+1)
    tags = pd.DataFrame({
        "id": tags_id,
        "name": tags_name
        })
    tags.to_csv("./tags_table.csv", index=False)

    # Produce .csv for ingredients table
    ingredients_name = set()
    for i in data["ingredients"]:
        for item in ast.literal_eval(i):
            ingredients_name.add(item)
    ingredients_name = list(ingredients_name)
    ingredients_id = range(1,len(ingredients_name)+1)
    ingredients = pd.DataFrame({
        "id": ingredients_id,
        "name": ingredients_name
        })
    ingredients.to_csv("./ingredients_table.csv", index=False)

    # Produce .csv for recipe_ingredient table
    ingredient_to_id = {a: b for a, b in zip(ingredients_name, ingredients_id)}
    first, second = [], []
    for idx, i in enumerate(data["ingredients"]):
        first += [idx+1]*len(ast.literal_eval(i))
        for yes in ast.literal_eval(i):
            second.append(ingredient_to_id[yes])
    assert len(first) == len(second)
    recipe_ingredient = pd.DataFrame({
        "recipe_id": first,
        "ingredient_id": second
        })
    recipe_ingredient.to_csv("./recipe_ingredient_table.csv", index=False)

    # Produce .csv for recipe_tag table
    tag_to_id = {a: b for a, b in zip(tags_name, tags_id)}
    first, second = [], []
    for idx, i in enumerate(data["tags"]):
        first += [idx+1]*len(ast.literal_eval(i))
        for yes in ast.literal_eval(i):
            second.append(tag_to_id[yes])
    assert len(first) == len(second)
    recipe_tag = pd.DataFrame({
        "recipe_id": first,
        "tag_id": second
        })
    recipe_tag.to_csv("./recipe_tag_table.csv", index=False)


if __name__ == "__main__":
    main()
