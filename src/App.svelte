<script>
  import Router from 'svelte-spa-router'
  import Home from './routes/Home.svelte'
  import Recipes from './routes/Recipes.svelte'
  import Recipe from './routes/Recipe.svelte'
  import NotFound from './routes/NotFound.svelte'
  import WeeklyRecipes from './routes/WeeklyRecipes.svelte'
  import FavoriteRecipes from './routes/FavoriteRecipes.svelte'
  import RecommendedRecipes from './routes/RecommendedRecipes.svelte';
  import AddRecipes from './routes/AddRecipes.svelte'
  import LoginSignupV2 from './routes/LoginSignupV2.svelte'
  import NavBar from './lib/components/Navbar.svelte'
  import Foot from './lib/components/Foot.svelte'
  import UserAuthoredRecipes from './routes/UserAuthoredRecipes.svelte'

  let routes = {
    '/': Home,
    '/recipes': Recipes,
    '/recipes/weekly': WeeklyRecipes,
    '/recipes/favorite': FavoriteRecipes,
    '/recipes/recommended': RecommendedRecipes,
    '/recipes/add': AddRecipes,
    '/recipes/author/:id': UserAuthoredRecipes,
    '/recipes/id/:id': Recipe,
    '/login': LoginSignupV2,
    '/signup': LoginSignupV2,
    '/notfound': NotFound,
    '*': NotFound,
  }
</script>

<NavBar />
<main>
  <Router {routes} />
</main>
<Foot />

<style>
</style>
