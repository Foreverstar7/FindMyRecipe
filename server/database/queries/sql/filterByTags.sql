create or replace function filterByTags(int)
  returns recipes
  as $$
  select * from recipes
    where id in (select recipe_id from recipe_tag
      where tag_id = $1);
  $$ language sql;

