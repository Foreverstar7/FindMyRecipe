<script>
  import PaginationWrapper from '../lib/components/PaginationWrapper.svelte';
  import { getRecipesForPage } from '../lib/utils/recipes.js'
  import { getTagsOfRecipe } from '../lib/utils/tags.js'
  import { querystring } from 'svelte-spa-router'
  import { RECIPE_NUMS } from '../lib/utils/utils';

  let params = new URLSearchParams($querystring)
  let page = params.get('page') ? parseInt(params.get('page'), 10) : 1 

  async function fetchRecipeData(page) {
    let recipes = []
    let tags = []

    await getRecipesForPage(page)
      .then((res) => {
        return res.data
      })
      .then(async (data) => {
        recipes = data
        let ptags = await Promise.all(
          data.map(async (r) => {
            return (await getTagsOfRecipe(r.Id)).data
          }),
        )
        tags = ptags
      })

    return { recipes: recipes, tags: tags }
  }
  let pageData = fetchRecipeData(page)

  async function handleChangePage(event) {
    page = event.detail.page
    pageData = fetchRecipeData(page)
    params.set('page', String(page))
    const currentRoute = window.location.href.split('?')[0]
    window.location.href = currentRoute + '?' + params.toString()
  }
</script>

<div class="flex flex-col flex-grow pt-10 px-10 bg-base-200" style="min-height: 85vh;">
  <div class="flex w-full justify-center">
      <h1 class="text-5xl font-bold">Recipes</h1>
  </div>
  <div class="form-control justify-center hidden">
    <!-- <SearchBar /> -->
    <!-- QQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQ -->

    <div class="flex justify-center flex-col">
      <div class="input-group max-w-3xl">
        <input
          type="text"
          placeholder="Search for Recipes"
          class="input input-bordered w-full"
        />
        <button
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
    </div>
  </div>
  {#await pageData}
    <div class="flex justify-center text-2xl mt-3">
      Loading...
    </div>
  {:then data}
    {#if data.recipes.length == 0}
      <div class="flex justify-center text-2xl mt-3">
        Sorry! Unfortunately we have insufficient data! Come back later!
      </div>
    {/if}
    <PaginationWrapper
      recipes={data.recipes}
      tags={data.tags}
      page={page}
      client={true}
      totalPages={Math.ceil(RECIPE_NUMS / 6)}
      on:pageChange={handleChangePage}
    />
  {/await}
</div>
