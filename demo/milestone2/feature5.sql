SELECT * FROM recipes WHERE id IN ( SELECT recipe_id FROM (
  SELECT recipe_id, COUNT(*) AS favCount FROM user_favorites
  WHERE liked_at >= '2022-11-16 21:50:56.44249'
  GROUP BY recipe_id
  ORDER BY favCount DESC 
) AS top_recipes
LIMIT 5 )
