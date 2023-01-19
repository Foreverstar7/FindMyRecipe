create or replace function matchRecipes(int[])
  returns recipes
  as $$
  SELECT * FROM recipes AS r
  WHERE r.id NOT IN (
    SELECT recipe_id FROM recipe_ingredient AS ri
      WHERE ri.ingredient_id NOT IN (
        SELECT * FROM UNNEST($1)
      )
  ); 
  $$ language sql;

