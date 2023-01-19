create or replace function autocompleteSearch(text)
  returns ingredients
  as $$
  SELECT * FROM ingredients WHERE name LIKE $1;
  $$ language sql;

