<script>
  import PaginationWrapper from '../lib/components/PaginationWrapper.svelte'
  import AutocompleteSearch from '../lib/components/AutocompleteSearch.svelte'
  import { matchRecipes } from '../lib/utils/recipes.js'
  import {
    autocompleteIngredients,
    getIngredientByName,
  } from '../lib/utils/ingredients.js'
  import {
    getTagsOfRecipe,
    getTagByName,
    autocompleteTags,
  } from '../lib/utils/tags.js'
  import { ENTER_KEY } from '../lib/utils/utils.js'

  let loaded = false
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

    if (e.keyCode == ENTER_KEY) {
      handleInputEnd()
      return
    }
    inputTimeout = setTimeout(async () => {
      searchBar = searchBar.trim()
      await handleAutoComp()
    }, 200)
  }

  const setInputVal = (input) => {
    if (searchBar != '' && searchBar[0] == '#') {
      !tags.includes(input) ? tags.push(input) : ''
      tags = tags
    } else {
      !fridge.includes(input) ? fridge.push(input) : ''
      fridge = fridge
    }
    searchBar = ''
    searchMatches = []
  }

  const mapTags = async (tagNames) => {
    let ids = []
    ids = await Promise.all(
      tagNames.map(async (i) => {
        console.log(i)
        return (await getTagByName(i)).data.id
      }),
    )
    console.log('tagIds: ', ids)
    return ids
  }

  const mapIngredients = async (ingredientNames) => {
    let ids = []
    ids = await Promise.all(
      ingredientNames.map(async (i) => {
        console.log(i)
        return (await getIngredientByName(i)).data.id
      }),
    )
    console.log('ingredientIds: ', ids)
    return ids
  }

  const handleAutoComp = async () => {
    searchMatches = []
    searchStatus = ''
    let response
    if (searchBar == '') {
      return
    } else if (searchBar[0] == '#') {
      response = await autocompleteTags(searchBar.substring(1))
    } else {
      response = await autocompleteIngredients(searchBar)
    }

    if (response.data == null) {
      return
    }
    response.data.forEach((i) => {
      searchMatches.push(i.name)
    })
    searchMatches = searchMatches
  }

  const searchRecipes = async () => {
    loaded = false
    recipes = []
    let req = await Promise.all([mapIngredients(fridge), mapTags(tags)])
    console.log(req[0], req[1])
    let res = await matchRecipes(req[0], req[1])
    console.log(res)
    if (res.status == 'success') {
      recipes = res.data
      let ptags = await Promise.all(
        recipes.map(async (r) => {
          return (await getTagsOfRecipe(r.Id)).data
        }),
      )
      matchedTags = ptags
      console.log(matchedTags)
      loaded = true
    } else {
      searchStatus = res.message
    }
  }

  const handleInputEnd = () => {
    let isIngredient = true
    let input = searchBar
    if (searchBar != '' && searchBar[0] == '#') {
      isIngredient = false
      input = searchBar.substring(1)
    }
    if (searchMatches.length == 1) {
      input = searchMatches[0]
    }
    if (searchMatches.includes(input)) {
      if (isIngredient && !fridge.includes(input)) {
        fridge.push(input)
      } else if (!isIngredient && !tags.includes(input)) {
        tags.push(input)
      }
    } else {
      return
    }
    searchBar = ''
    fridge = fridge
    tags = tags
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

  const stopSuggestion = () => {
    searchMatches = []
  }

  $: searchBarMargin = recipes.length == 0 ? 'hidden' : 'h-20'
  let page = 1;
</script>

<div on:click={stopSuggestion}>
  <div class="hero min-h-screen bg-base-200">
    <div class="flex flex-col hero-content p-0 text-center">
      <div class="flex flex-grow {searchBarMargin}"></div>
      <div class="max-w-xl">
        <h1 class="text-5xl font-bold">Find What You Can Cook</h1>
        <p class="py-6">
          Just type your ingredients, and we'll show you what YOU can cook!
        </p>
        <div class="form-control items-center">
          <!-- <SearchBar /> -->
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

          <div class="fridge-tags">
            <br style="margin-bottom:0" />
            {#if tags.length > 0}
              <ul class="">
                {#each tags as tag, i}
                  <div class="badge gap-2 badge-lg">
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
              <ul class="">
                {#each fridge as ing, i}
                  <div class="badge badge-outline badge-lg gap-2">
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
        </div>
      </div>
      {#if searchMatches.length > 0}
        <ul>
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
      <div class="w-screen">
        {#await searchRecipes}
          <div class="flex justify-center text-2xl mt-3">
            Loading...
          </div>
        {:then}
          {#if loaded == true && recipes.length > 0}
            <PaginationWrapper 
              recipes={recipes} 
              tags={matchedTags} 
              page={page}
              totalPages={Math.ceil(recipes.length / 6)}
            />
          {:else}
            {searchStatus}
          {/if}
        {/await}
      </div>
    </div>
  </div>
</div>

<svelte:window
  on:keyup={(e) => {
    handleAutoCompTO(e)
  }}
/>
