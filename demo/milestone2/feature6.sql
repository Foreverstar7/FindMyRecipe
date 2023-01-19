((SELECT id FROM recipes)
  EXCEPT
  (SELECT DISTINCT recipe_id FROM recipe_ingredient
    WHERE NOT (ingredient_id = ANY (CAST('{4621,4684,6557,8899,10583,12245,14832, 12149}' AS INT[])))))
EXCEPT
((SELECT id FROM recipes)
  EXCEPT
  (SELECT DISTINCT recipe_id FROM recipe_ingredient
    WHERE NOT (ingredient_id = ANY (CAST('{4621,4684,6557,8899,10583,12245,14832}' AS INT[])))));
