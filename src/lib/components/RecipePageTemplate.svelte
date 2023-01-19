<!-- TODO: Refactor recipe page with this preview component -->
<script>
  export const toggleFavorite = () => {}
  // export let onLogin = false
  // export let starFill = 'none'
  export let recipe = {
    Id: 0,
    Name: 'Recipe Name',
    PrepMinutes: 0,
    Description: '',
    Ingredients: [],
    Nutrition: Array(7).fill(0),
    Tags: [],
    Steps: [],
  }

  import {
    capitalizeFirstChar,
    NUTRITION_NAMES,
    USER_FEATURES,
  } from '../utils/utils'
  import RecipeTags from './RecipeTags.svelte'
  import SimilarRecipes from './SimilarRecipes.svelte'
</script>

<!-- ONLY FOR PREVIEW PURPOSE -->

<div class="min-h-screen bg-base-200" style="width:100%">
  <div class="flex flex-row w-full">
    <div
      class="items-center align-middle text-center md:p-2 basis-1/4 bg-base-300 rounded-box"
    >
      <div class="prose">
        <h4 class="py-2">Tags</h4>
        <hr>
        {#if recipe.Tags.length != 0}
          <ul style="list-style:none;margin:0;padding:0">
            {#each recipe.Tags as tag, i}
              <li>#{tag}</li>
            {/each}
          </ul>
        {/if}
      </div>
    </div>
    <div class="flex-grow w-full justify-center p-10">
      <div class="flex-col">
        <div class="flex p-5 text-center justify-center">
          <h1 class="text-5xl font-bold">{recipe.Name}</h1>
        </div>
        <div class="flex mb-3 text-center justify-center">
          <p class="text-lg font-semibold">
            Prep Time: {recipe.PrepMinutes} minutes
          </p>
        </div>
        <hr />
        <div class="flex-col my-4">
          <div class="flex text-lg font-semibold">Description:</div>
          <div class="flex text-justify">{recipe.Description}</div>
        </div>
        <div class="flex flex-col md:flex-row my-3">
          <div class="flex flex-col justify-center md:basis-1/2">
            <div class="flex text-lg font-semibold">Ingredients:</div>
            <div class="flex mt-1 text-start">
              <ol>
                {#each recipe.Ingredients as ingredient, i}
                  <li>{i + 1}. {capitalizeFirstChar(ingredient)}</li>
                {/each}
              </ol>
            </div>
          </div>

          <div class="flex flex-col justify-center md:basis-1/2 mt-3 md:mt-0">
            <div class="flex text-lg font-semibold">Nutrition:</div>
            <div class="flex mt-1 text-start">
              <ul>
                {#each recipe.Nutrition as nutrition, i}
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
              {#each recipe.Steps as step, i}
                <li>
                  {i + 1}. {step.length > 2
                    ? capitalizeFirstChar(step.substr(1, step.length - 2))
                    : ''}
                </li>
              {/each}
            </ol>
          </div>
        </div>
        <br />
      </div>
    </div>
  </div>
</div>
