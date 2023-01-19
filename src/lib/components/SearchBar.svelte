<!-- TODO: refactor Home with this somehow -->
<script>
  import RecipeBannerV2 from '../components/RecipeBannerV2.svelte'
  import { matchRecipes } from '../utils/recipes.js'
  import {
    autocompleteIngredients,
    getIngredientByName,
  } from '../utils/ingredients.js'
  import { getTagsOfRecipe } from '../utils/tags.js'

  let searchBar = ''
  let fridge = []
  let tags = []
  let recipes = []
  let searchMatches = []
  let matchedTags = []
  let searchStatus = ''
  let inputTimeout = null

  const handleAutoCompTO = (e) => {
    clearTimeout(inputTimeout)

    if (e.keyCode == 13) {
      handleInputEnd()
      return
    }
    // Make a new timeout set to go off in 1000ms (1 second)
    inputTimeout = setTimeout(async () => {
      if (searchBar != '' && searchBar[0] != '#') {
        await handleAutoComp()
      }
    }, 200)
  }

  const setInputVal = (ingredient) => {
    console.log('setInputValClick')
    fridge.push(ingredient)
    fridge = fridge
    searchBar = ''
    searchMatches = []
  }

  const mapIngredients = async (ingredientNames) => {
    let ids = []
    //ids = [4621,4684,6557,8899,10583,12245,14832]
    ids = await Promise.all(
      ingredientNames.map(async (i) => {
        console.log(i)
        return (await getIngredientByName(i)).data.id
      }),
    )
    console.log('mapping')
    console.log(ids)
    return ids
  }

  const handleAutoComp = async () => {
    searchMatches = []
    searchStatus = ''
    const response = await autocompleteIngredients(searchBar)
    if (response.data == null) {
      return
    }
    response.data.forEach((i) => {
      searchMatches.push(i.name)
    })
    searchMatches = searchMatches
  }

  const searchRecipes = async () => {
    recipes = []
    matchRecipes(await mapIngredients(fridge))
      .then(async (res) => {
        console.log(res)
        if (res.status == 'success') {
          recipes = res.data
          let ptags = await Promise.all(
            recipes.map(async (r) => {
              return (await getTagsOfRecipe(r.Id)).data
            }),
          )
          matchedTags = ptags
        } else {
          searchStatus = res.message
        }
      })
      .catch((err) => {
        console.log(err)
      })
  }
  const handleInputEnd = () => {
    console.log('handleInputEnd')
    console.log(searchBar)
    if (
      searchBar != '' &&
      searchBar[0] == '#' &&
      !tags.includes(searchBar.substring(1))
    ) {
      tags.push(searchBar.substring(1))
      tags = tags
    } else if (
      !fridge.includes(searchBar) &&
      searchMatches.includes(searchBar)
    ) {
      fridge.push(searchBar)
    } else if (
      searchMatches.length == 1 &&
      !fridge.includes(searchMatches[0])
    ) {
      fridge.push(searchMatches[0])
    } else {
      return
    }
    searchBar = ''
    fridge = fridge
    searchMatches = []
  }

  const removeFromFridge = (ing) => {
    const index = fridge.indexOf(ing)
    if (index > -1) {
      fridge.splice(index, 1)
    }
    fridge = fridge
  }

  const removeFromTags = (tag) => {
    const index = tags.indexOf(tag)
    if (index > -1) {
      tags.splice(index, 1)
    }
    tags = tags
  }
</script>

<!-- abandoned -->


<!-- QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ -->

<div class="input-group">
  <input
    on:click={() => {
      handleInputEnd()
    }}
    bind:value={searchBar}
    type="text"
    placeholder="Type ingredients and/or add a '#' in front to filter by tags"
    class="input input-bordered w-full"
  />
  <button
    on:click={async () => {
      await searchRecipes()
    }}
    class="btn btn-square"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-6 w-6"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
      ><path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
      /></svg
    >
  </button>
</div>

<!-- ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ -->

{#if searchMatches.length > 0}
  <ul id="autocomplete-items-list">
    {#each searchMatches as ingredient, i}
      <div
        class="btn btn-ghost btn-wide normal-case"
        on:click={() => setInputVal(ingredient)}
      >
        <li>{ingredient}</li>
      </div>
    {/each}
  </ul>
{/if}

<div class="fridge-tags">
  <br style="margin-bottom:0" />
  {#if tags.length > 0}
    <ul class="">
      {#each tags as tag, i}
        <div class="badge gap-2">
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
  {#if fridge.length > 0}
    <h3><b>Your Fridge:</b></h3>
    <ul class="">
      {#each fridge as ing, i}
        <div class="badge badge-info gap-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="inline-block w-4 h-4 stroke-current"
            on:click={() => removeFromFridge(ing)}
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
</div>

<!-- WWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWW -->

<div class="items-center">
  {#await tags then}
    {#if recipes.length > 0}
      <div class="items-center text-center">
        {#each recipes as recipe, i}
          <RecipeBannerV2
            recipeId={recipe.Id}
            recipeLabel={recipe.Name}
            description={recipe.Description}
            prepminutes={recipe.PrepMinutes}
            tags={matchedTags[i]}
            checkFavorite="false"
          />
        {/each}
      </div>
    {:else}
      {searchStatus}
    {/if}
  {/await}
</div>

<svelte:window
  on:keyup={(e) => {
    handleAutoCompTO(e)
  }}
/>
