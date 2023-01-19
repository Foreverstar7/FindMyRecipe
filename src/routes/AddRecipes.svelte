<script>
  import { getUserInfo, NUTRITION_NAMES } from '../lib/utils/utils'
  import { addNewRecipe } from '../lib/utils/recipes.js'

  import RecipePageTemplate from '../lib/components/RecipePageTemplate.svelte'
  import AutocompleteSearch from '../lib/components/AutocompleteSearch.svelte'
  import NotFound from './NotFound.svelte'

  let recipeName = ''
  let prepMinutes = ''
  let ingredients = []
  let tags = []
  let nutrition = Array(7).fill(0)
  let description = ''
  let steps = []

  let args = {
    recipe: {
      Id: 0,
      Name: 'Recipe Name',
      PrepMinutes: 0,
      Description: '',
      Ingredients: [],
      Nutrition: Array(7).fill(0),
      Tags: [],
      Steps: [],
    },
  }

  let ingredientSearchBar
  let tagSearchBar

  let field
  let userinfo = getUserInfo()
  let onLogin = userinfo == null ? false : true
  let inputError = false
  let errorMsgs = []
  let addRecipeSuccess = false
  let addRecipeFail = false

  const resetErrors = () => {
    inputError = false
    errorMsgs = []
  }

  const checkInputs = () => {
    if (recipeName == '') {
      errorMsgs.push('Missing recipe name')
    }
    if (prepMinutes <= 0 || prepMinutes == '') {
      errorMsgs.push('Preparation time must be positive')
    }
    if (description == '') {
      errorMsgs.push('Missing description')
    }
    if (ingredients.length == 0) {
      errorMsgs.push('Missing ingredients')
    }
    if (tags.length == 0) {
      errorMsgs.push('Missing tags')
    }
    if (steps.length == 0) {
      errorMsgs.push('Missing steps')
    }
    if (nutrition === Array(7).fill(0)) {

    }
    let nSum = 0
    if (nutrition.reduce((acc, n) => {
        nSum += n
        return acc || n < 0
      }, false)) {
      errorMsgs.push('Nutrition value must be non-negative')
    } else if (nSum == 0) {
      errorMsgs.push('Missing nutrition values')
    }
    // now check
    if (errorMsgs.length != 0) {
      errorMsgs = errorMsgs
      inputError = true
      return false
    }
    return true
  }
  const addRecipe = () => {
    console.log('add recipe')
    if (checkInputs()) {
      console.log(formatRecipe())
      addNewRecipe(userinfo.userid, formatRecipe()).then((res) => {
        console.log(res.message)
        if (res.status != 'failed') { 
          addRecipeSuccess = true
        } else {
          addRecipeFail = true
        }
      })
    } else {
      console.log('tes')
      inputError = true
    }
  }

  const removeFromIngredients = (ing) => {
    const index = ingredients.indexOf(ing)
    if (index > -1) {
      ingredients.splice(index, 1)
    }
    ingredients = ingredients
  }

  const removeFromTags = (tag) => {
    const index = tags.indexOf(tag)
    if (index > -1) {
      tags.splice(index, 1)
    }
    tags = tags
  }

  const removeFromSteps = (stepNum) => {
    const index = stepNum
    if (index > -1) {
      steps.splice(index, 1)
    }
    steps = steps
  }

  const formatRecipe = () => {
    return {
      Id: 0,
      Name: recipeName || 'Recipe Name',
      PrepMinutes: prepMinutes || 0,
      Description: description || 'Description',
      Ingredients: ingredients,
      Nutrition: nutrition.length == 0 ? Array(7).fill(0) : nutrition,
      Tags: tags,
      Steps: steps.map((s) => "'" + s + "'"),
    }
  }

  const updatePreview = () => {
    steps.map((s) => {
      if (!s) return ''
      return s
    })
    args = {
      recipe: formatRecipe(),
    }
  }
</script>

{#if !onLogin}
  <NotFound />
{:else}
  <div>
    <div
      class="flex-col flex-grow pt-10 px-10 bg-base-200"
      style="min-height: 85vh"
      on:click={() => {
        field.stopSuggestion()
        updatePreview()
      }}
    >
      <div class="flex my-2 w-full justify-center">
        <h1 class="text-5xl font-bold">Add Recipes</h1>
      </div>
      <div class="flex flex-row justify-center p-5 " style="min-height: 70vh;">
        <div
          class="form-control w-auto card grid basis-2/5 flex-none bg-base-300 rounded-box items-start"
          on:change={updatePreview}
          on:click={updatePreview}
        >
          <div class="label">
            <span class="label-text">Recipe Name</span>
          </div>
          <input
            type="text"
            placeholder="Type here"
            class="input input-bordered w-full max-w-xs"
            bind:value={recipeName}
          />
          <div class="label">
            <span class="label-text">Preparation Time (minutes)</span>
          </div>
          <input
            type="number"
            min="0"
            placeholder="Type a number"
            class="input input-bordered w-full max-w-xs"
            bind:value={prepMinutes}
          />
          <div class="label">
            <span class="label-text">Ingredients</span>
          </div>
          <AutocompleteSearch
            bind:resultArray={ingredients}
            bind:searchBar={ingredientSearchBar}
            bind:this={field}
          />
          {#if ingredients.length > 0}
            <ul class="">
              {#each ingredients as ing, i}
                <div class="badge gap-2 badge-lg">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    class="inline-block w-4 h-4 stroke-current"
                    on:click={() => removeFromIngredients(ing)}
                    ><path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M6 18L18 6M6 6l12 12"
                    /></svg
                  >
                  {ing}
                </div>
              {/each}
            </ul>
          {/if}
          <div class="label">
            <span class="label-text">Tags</span>
          </div>
          <AutocompleteSearch
            bind:resultArray={tags}
            bind:searchBar={tagSearchBar}
            searchIng={false}
            bind:this={field}
          />
          {#if tags.length > 0}
            <ul class="">
              {#each tags as tag, i}
                <div class="badge badge-lg gap-2">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    class="inline-block w-4 h-4 stroke-current"
                    on:click={() => removeFromTags(tag)}
                    ><path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M6 18L18 6M6 6l12 12"
                    /></svg
                  >
                  #{tag}
                </div>
              {/each}
            </ul>
          {/if}
          <div class="label">
            <span class="label-text">Nutrition</span>
          </div>
          <div class="flex flex-col">
            {#each Array(7) as _, i}
              <label class="input-group">
                <input
                  type="number"
                  min="0"
                  placeholder="0"
                  class="input input-bordered"
                  bind:value={nutrition[i]}
                />
                <span>{NUTRITION_NAMES[i]}</span>
              </label>
            {/each}
          </div>
          <div class="label">
            <span class="label-text">Description</span>
          </div>
          <textarea
            placeholder="Type here"
            class="textarea textarea-bordered w-full max-w-xs"
            bind:value={description}
          />
          <div class="label">
            <span class="label-text">Steps</span>
          </div>
          {#each Array(steps.length + 1) as _, i}
            <label class="input-group">
              <span>{i + 1}</span>
              <textarea
                placeholder="Type here"
                class="textarea textarea-bordered w-full max-w-xs"
                bind:value={steps[i]}
              />
              <button class="btn h-full" on:click={() => removeFromSteps(i)}>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="#ffffff"
                  stroke-width="2"
                  stroke-linecap="butt"
                  stroke-linejoin="bevel"
                  ><line x1="18" y1="6" x2="6" y2="18" /><line
                    x1="6"
                    y1="6"
                    x2="18"
                    y2="18"
                  /></svg
                >
              </button>
            </label>
          {/each}
          <div class="flex justify-end mt-8">
            <button class="btn" on:click={addRecipe}>Add Your Recipe!</button>
          </div>
        </div>
        <div
          class="flex-grow ml-4 bg-base-200 rounded-box"
          style="pointer-events:none"
        >
          <RecipePageTemplate bind:recipe={args.recipe} />
        </div>
      </div>
    </div>
  </div>
  {#if inputError}
    <input type="checkbox" id="my-modal" class="modal-toggle" checked />
    <div class="modal">
      <div class="modal-box align-middle">
        <div class="flex mb-2 mt-4 justify-center">
          <h3 class="font-bold text-2xl">
            Oops! Please make sure your recipe is written properly.
          </h3>
        </div>
        <p class="py-4">
          {#each errorMsgs as msg}
            {msg} <br />
          {/each}
        </p>
        <div class="modal-action">
          <label for="my-modal" class="btn" on:click={resetErrors}>Close</label>
        </div>
      </div>
    </div>
  {/if}
  {#if addRecipeSuccess}
        <input type="checkbox" id="my-modal" class="modal-toggle" checked />
        <div class="modal">
          <div class="modal-box align-middle">
            <div class="flex mb-2 mt-4 justify-center">
              <h3 class="font-bold text-2xl">Add Recipe Success!</h3>
            </div>
            <p class="py-4">
              You will be redirected home
            </p>
            <div class="modal-action">
              <label for="my-modal" class="btn" on:click={() => { window.location.href = "#/"; window.location.reload() }}>Close</label>
            </div>
          </div>
        </div>
      {:else if addRecipeFail}
        <input type="checkbox" id="my-modal" class="modal-toggle" checked />
        <div class="modal">
          <div class="modal-box align-middle">
            <div class="flex mb-2 mt-4 justify-center">
              <h3 class="font-bold text-2xl">Add Recipe Fail!</h3>
            </div>
            <p class="py-4">
              Failed to add recipe
            </p>
            <div class="modal-action">
              <label for="my-modal" class="btn" >Close</label>
            </div>
          </div>
        </div>
  {/if}
{/if}
