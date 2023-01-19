<script>
  export let recipeId
  export let knownTags = []
  import { getTagsOfRecipe } from '../utils/tags.js'

  let tags = []
  if (knownTags.length == 0) {
    getTagsOfRecipe(recipeId).then((res) => {
      tags = res.data
    })
  } else {
    tags = knownTags.map((i) => ({ name: i }))
  }
</script>

<div class="prose">
  <h4>Tags</h4>
  {#if tags.length != 0}
    <ul style="list-style:none;margin:0;padding:0">
      {#each tags as tag, i}
        <li>#{tag.name}</li>
      {/each}
    </ul>
  {/if}
</div>
