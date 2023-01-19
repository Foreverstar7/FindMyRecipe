<script>
  import RecipeBannerV2 from '../components/RecipeBannerV2.svelte'
  import RecipePagination from '../components/RecipePagination.svelte'
  import { createEventDispatcher } from 'svelte';
  export let recipes;
  export let tags;
  export let page;
  export let totalPages;
  export let client = false;

  const dispatch = createEventDispatcher();

  function handleChangePage(event) {
    page = event.detail.page;
    dispatch('pageChange', { page: page })
  }

  let itemsPerPage = 6;

  function paginate(items) {
    if (items.length == 0) {
      return [[]]
    }
    const pagination = []
    for (let i = 0; i < items.length; i += itemsPerPage) {
      const chunk = items.slice(i, i + itemsPerPage)
      pagination.push(chunk)
    }
    return pagination
  }

  let recipePaginated = client ? [] : paginate(recipes)
  let tagsPaginated = client ? [] : paginate(tags)
</script>

<div class="flex flex-col justify-center p-8" style="min-height: 70vh;">
  <div class="md:grid-cols-3 grid grid-cols-1 grid-flow-row mb-4">
    <!-- {#each recipesPaginated[page - 1] as recipe, i} -->
    {#if client}
      {#each recipes as recipe, i}
        <RecipeBannerV2
          recipeId={recipe.Id}
          recipeLabel={recipe.Name}
          description={recipe.Description}
          prepminutes={recipe.PrepMinutes}
          tags={tags[i]}
          checkFavorite="false"
        />
      {/each}
    {:else}
      {#each recipePaginated[page - 1] as recipe, i}
        <RecipeBannerV2
          recipeId={recipe.Id}
          recipeLabel={recipe.Name}
          description={recipe.Description}
          prepminutes={recipe.PrepMinutes}
          tags={tagsPaginated[page - 1][i]}
          checkFavorite="false"
        />
      {/each}
    {/if}
  </div>
  <div class="flex flex-grow h-auto"></div>
  <RecipePagination
    page={page}
    totalPages={totalPages}
    on:pageChange={handleChangePage}
  />
</div>
