<script>
  import PaginationWrapper from '../lib/components/PaginationWrapper.svelte';
  import { getUserFavorites } from '../lib/utils/favorites'
  import { getTagsOfRecipe } from '../lib/utils/tags.js'
  import { getUserInfo } from '../lib/utils/utils'
  import NotFound from './NotFound.svelte'

  let userinfo = getUserInfo()
  let onLogin = userinfo == null ? false : true
  let recipes = []
  let tags = new Promise((res, rej) => res)
  let page = 1

  if (onLogin) {
    getUserFavorites(userinfo.userid)
      .then((res) => {
        return res
      })
      .then(async (res) => {
        console.log(res)
        if (res.status != 'failed') {
          recipes = res.data
        } else {
          tags = []
          return
        }
        let ptags = await Promise.all(
          recipes.map(async (r) => {
            return (await getTagsOfRecipe(r.Id)).data
          }),
        )
        tags = ptags
      })
  }
</script>

{#if !onLogin}
  <NotFound />
{:else}
<div class="flex flex-col bg-base-200 pt-10 px-10 items-center" style="min-height: 85vh;">
  <div class="flex w-full justify-center">
    <h1 class="text-5xl font-bold">Favorites</h1>
  </div>
  {#await tags}
    <div class="flex justify-center text-2xl mt-3">
      Loading...
    </div>
  {:then}
    {#if recipes.length == 0}
      <div class="justify-center flex text-2xl mt-3">
        So empty! Add recipes to your favorites!
      </div>
    {:else}
      <PaginationWrapper
        recipes={recipes}
        tags={tags}
        page={page}
        totalPages={Math.ceil(recipes.length / 6)}
      />
    {/if}
  {/await}
</div>
{/if}
