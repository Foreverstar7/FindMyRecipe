<script>
  import RecipeBanner from '../components/RecipeBanner.svelte'
  import { getSimilarRecipes, getRecipeById } from '../utils/recipes.js'
  export let ingredientIds

  let recipes = []
  getSimilarRecipes(ingredientIds).then((res) => {
    recipes = res.data.map((rec) => getRecipeById(rec.Id))
    Promise.all(recipes).then((res) => {
      recipes = res.map((r) => r.data)
    })
  })
</script>

<div>
  {#if recipes.length != 0}
    <div class="prose"><h4>Similar Recipes</h4></div>
    <ul>
      {#each recipes as recipe, i}
        <RecipeBanner recipeId={recipe.Id} recipeLabel={recipe.Name} />
      {/each}
    </ul>
  {/if}
</div>
