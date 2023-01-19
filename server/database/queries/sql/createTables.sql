create or replace function createTables()
  returns void
  as $$
  CREATE TABLE if not exists users (
    id serial PRIMARY KEY,
    name text,
    password text
  );

  create table if not exists recipes (
    id serial primary key, 
    name text,
    description text,
    prep_minutes int,
    author int,
    steps text[],
    nutrition float8[],

    constraint fk_author
        foreign key(author) references users(id)
  );

  create table if not exists user_favorites (
    user_id int references users,
    recipe_id int references recipes,
    liked_at timestamp default current_timestamp,
    primary key (user_id, recipe_id)
  );

  CREATE TABLE if not exists tags (
    id serial PRIMARY KEY,
    name text
  );

  CREATE TABLE if not exists ingredients (
    id serial PRIMARY KEY, 
    name text
  );

  create table if not exists recipe_ingredient (
    recipe_id int references recipes,
    ingredient_id int references ingredients,
    primary key (recipe_id, ingredient_id)
  );

  create table if not exists recipe_tag (
    recipe_id int references recipes,
    tag_id int references tags,
    primary key (recipe_id, tag_id)
  );

  create index FilterTags on recipe_tag(tag_id);
  create index MatchRecipe on recipe_ingredient(ingredient_id);
  create index RecipeByName on recipes(name);
  create index RecipeById on recipes(id);
  create index IngredientByName on ingredients(name);
  create index UserFavoritesByUser on user_favorites(user_id);
  create index UserFavoritesByRecipe on user_favorites(recipe_id);
  create index RecipeByAuthor on recipes(author);
  $$ language sql;

