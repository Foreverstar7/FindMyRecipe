<script>
  export let recipeLabel
  export let recipeId
</script>

{#if recipeId != undefined}
  <li
    class="btn gap-2 btn-ghost normal-case btn-wide" style="font-family:'Helvetica',accent-color:#A6ADBACC"
    onclick="javascript:location.href='#/recipes/id/{recipeId}';location.reload()"
  >
    {@html recipeLabel}
  </li>
{/if}
