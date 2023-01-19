<script>
  export let recipeLabel
  export let recipeId
  export let description
  export let prepminutes
  export let tags = []
  // export let checkFavorite

  const snipText = (text, length) => {
    if (text.length > length) {
      return text.substring(0, length) + '...'
    }
    return text
  }
</script>

<div class="flex flex-col w-auto h-80 bg-base-100 rounded-box shadow-xl m-3 p-5">
  <div class="flex flex-col justify-center py-3 px-5">
    <h2 class="text-2xl font-semibold">
      {recipeLabel}
      <!--<div>
  <input type="checkbox" name="favorite" class="star" bind:checked={checkFavorite} />
    </div>-->
    </h2>
    <h4 class="mt-1">{prepminutes} minutes</h4>
  </div>
  <p class="text-justify px-4">{snipText(description, 100)}</p>
  <div class="flex flex-row justify-center mt-2">
    {#if tags?.length > 0}
      {#each tags as tag, i}
        {#if i < 3}
          <div class="badge badge-sm lg:badge-md badge-outline mx-0.5">{snipText(tag.name, 16)}</div>
        {/if}
      {/each}
    {/if}
  </div>
  <div class="flex-grow h-auto"></div>
  <div class="flex justify-end">
    <button
      onclick="javascript:location.href='#/recipes/id/{recipeId}';location.reload()"
      class="btn btn-primary">Cook Now</button
    >
  </div>
</div>

