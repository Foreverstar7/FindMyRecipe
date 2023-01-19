SELECT * FROM recipes WHERE id IN ( SELECT recipesResult.recipe_id FROM (
  SELECT recipe_id, COUNT(*) AS favCount FROM user_favorites
  GROUP BY recipe_id
  HAVING COUNT(*) >= (
    SELECT AVG(favCount2) FROM (
      SELECT uf.recipe_id, COUNT(*) AS favCount2 FROM user_favorites as uf
      GROUP BY recipe_id
    ) as favRecipeWithCounts, user_favorites as uf2
    WHERE uf2.user_id = 2 AND uf2.recipe_id = favRecipeWithCounts.recipe_id
  ) 
) AS recipesResult
ORDER BY RANDOM()
LIMIT 3)
