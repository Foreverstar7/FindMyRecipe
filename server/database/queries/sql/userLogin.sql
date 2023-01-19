create or replace function userLogin(text)
  returns text
  as $$
  select password from users
    where name = $1;
  $$ language sql;

