create or replace function dropTables()
  returns void
  as $$
  drop table recipe_ingredient cascade;
  drop table recipe_tag cascade;
  drop table user_favorites cascade;
  drop table recipes cascade       ;
  drop table users cascade         ;
  drop table tags cascade          ;
  drop table ingredients cascade   ;
  $$ language sql;

