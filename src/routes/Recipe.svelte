<script>
  export let params = {}

  import SimilarRecipes from '../lib/components/SimilarRecipes.svelte'
  import RecipeTags from '../lib/components/RecipeTags.svelte'
  import { getRecipeById } from '../lib/utils/recipes.js'
  import { getIngredientsOfRecipe } from '../lib/utils/ingredients.js'
  import { toggleUserFavorite, getUserFavorites } from '../lib/utils/favorites.js'
  import { USER_FEATURES, NUTRITION_NAMES, capitalizeFirstChar, getUserInfo } from '../lib/utils/utils.js'
  import NotFound from './NotFound.svelte'

  let recipe = {}
  let ingredients = []
  let steps = []
  let nutritions = []
  let checkFavorite = false
  let starFill = 'none'
  let notfound = false

  let userinfo = getUserInfo()
  let onLogin = userinfo == null ? false : true
  if (onLogin) {
    getUserFavorites(userinfo.userid).then((res) => {
      if (res.data.map(r => r.Id).includes(parseInt(params.id))) {
        checkFavorite = true
        toggleStar()
      }
    })
  }

  const toggleStar = () => {
    if (checkFavorite) {
      starFill = 'white'
    } else {
      starFill = 'none'
    }
  }

  getRecipeById(params.id).then((res) => {
    if (res.status == "failed") {
      notfound = true;
      return
    }
    recipe = res.data
    nutritions = recipe.Nutrition
    steps = recipe.Steps
  })
  let ingredientPromise = getIngredientsOfRecipe(params.id).then((res) => {
    console.log(res)
    if (res.status != 'failed') { 
      ingredients = res.data
    } else {
      ingredients = []
    }
  })

  const toggleFavorite = () => {
    toggleUserFavorite(userinfo.userid, parseInt(params.id)).then((res) => {
      console.log(res)
      if (res.status == 'success') {
        checkFavorite = !checkFavorite
        toggleStar()
      }
    })
  }
</script>

{#if notfound}
  <NotFound />
{:else}
<div class="min-h-screen bg-base-200" style="width:100%">
  <div class="flex flex-row w-full">
    <div class="items-center align-middle text-center md:p-2 basis-1/4 bg-base-300 rounded-box" >
      {#if onLogin}
        <button on:click={toggleFavorite} class="btn mt-2 gap-2 flex w-full">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill={starFill}
            stroke="#ffffff"
            stroke-width="2"
            stroke-linecap="butt"
            stroke-linejoin="bevel"
            ><polygon
              points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"
            /></svg
          >
          FAVORITE
        </button>
      {:else}
        <label for="my-modal" class="btn gap-2 mt-2 flex mx-4">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill={starFill}
            stroke="#ffffff"
            stroke-width="2"
            stroke-linecap="butt"
            stroke-linejoin="bevel"
            ><polygon
              points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"
            /></svg
          >FAVORITE</label
        >
        <input type="checkbox" id="my-modal" class="modal-toggle" />
        <div class="modal">
          <div class="modal-box align-middle">
            <div class="flex mb-2 mt-4 justify-center">
              <h3 class="font-bold text-2xl">Login to access more features!</h3>
            </div>
            <p class="py-4">
              {#each USER_FEATURES as feat}
                {feat}<br />
              {/each}
            </p>
            <div class="modal-action">
              <label for="my-modal" class="btn">Close</label>
            </div>
          </div>
        </div>
      {/if}
      <div class="divider divider-vertical my-1" />
      <RecipeTags recipeId={params.id} />
      <div class="divider divider-vertical" />
      {#await ingredientPromise then { }}
        <SimilarRecipes ingredientIds={ingredients.map((i) => i.id)} />
      {/await}
    </div>
    <div class="flex-grow w-full justify-center p-10">
      <div class="flex-col">
        <div class="flex p-5 text-center justify-center">
          <h1 class="text-5xl font-bold">{recipe.Name}</h1>
        </div>
        <div class="flex mb-3 text-center justify-center">
          <p class="text-lg font-semibold">Prep Time: {recipe.PrepMinutes} minutes</p>
        </div>
        <hr>
        <div class="flex-col my-4">
          <div class="flex text-lg font-semibold">Description:</div>
          <div class="flex text-justify">{recipe.Description}</div>
        </div>
        <div class="flex flex-col md:flex-row my-3">
          <div class="flex flex-col justify-center md:basis-1/2">
            <div class="flex text-lg font-semibold">Ingredients:</div>
            <div class="flex mt-1 text-start">
              <ol>
                {#each ingredients as ingredient, i}
                  <li>{i + 1}. {capitalizeFirstChar(ingredient.name)}</li>
                {/each}
              </ol>
            </div>
          </div>

          <div class="flex flex-col justify-center md:basis-1/2 mt-3 md:mt-0">
            <div class="flex text-lg font-semibold">Nutrition:</div>
            <div class="flex mt-1 text-start">
              <ul>
                {#each nutritions as nutrition, i}
                  <li>{NUTRITION_NAMES[i]}: {nutrition}</li>
                {/each}
              </ul>
            </div>
          </div>
        </div>
        <div class="flex flex-col my-5">
          <div class="flex text-lg font-semibold">Steps:</div>
          <div class="flex mt-2 text-start">
            <ol type="1">
              {#each steps as step, i}
                <li>{i + 1}. {capitalizeFirstChar(step.substr(1, step.length - 2))}</li>
              {/each}
            </ol>
          </div>
        </div>
        <br />
      </div>
    </div>
  </div>
</div>
{/if}
