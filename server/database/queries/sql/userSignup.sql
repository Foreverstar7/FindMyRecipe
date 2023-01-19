create or replace function userSignup(text, text)
  returns void
  as $$
  insert into users(name, password) values ($1, $2);
  $$ language sql;

