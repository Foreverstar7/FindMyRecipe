<script>
  import { autocompleteTags } from '../utils/tags'
  import { autocompleteIngredients } from '../utils/ingredients'
  import { ENTER_KEY } from '../utils/utils'

  // parent should bind these values
  export let resultArray = []
  export let searchBar = ''
  export function stopSuggestion() {
    searchResults = []
  }

  // values don't need any bind
  export let searchIng = true // true for ingredients search, false for tags

  // local vars
  let autocompleteFunc = searchIng ? autocompleteIngredients : autocompleteTags
  let inputTimeout = null
  let searchResults = []

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
    if (!resultArray.includes(input)) {
      resultArray.push(input)
      resultArray = resultArray
    }
    searchBar = ''
    searchResults = []
  }

  const handleAutoComp = async () => {
    searchResults = []
    let response
    if (searchBar == '') {
      return
    } else {
      response = await autocompleteFunc(searchBar)
    }

    if (response.data == null) {
      return
    }
    searchResults = response.data.map((res) => res.name)
  }

  const handleInputEnd = () => {
    let input = searchBar
    if (searchResults.length == 1) {
      input = searchResults[0]
    }
    if (searchResults.includes(input)) {
      if (!resultArray.includes(input)) {
        resultArray.push(input)
        resultArray = resultArray
      }
    } else {
      return
    }
    searchBar = ''
    searchResults = []
  }

</script>

<div class="form-control items-center">
  <div class="input-group mb-2">
    <input
      on:click={() => {
        handleInputEnd()
      }}
      bind:value={searchBar}
      type="text"
      placeholder="Type {searchIng ? 'ingredients' : 'tags'}"
      class="input input-bordered rounded w-full max-w-xs"
    />
  </div>
</div>
{#if searchResults.length > 0}
  <div class="py-2">
  <ul>
    {#each searchResults as result, i}
      <div
        class="btn btn-ghost normal-case"
        on:click={() => setInputVal(result)}
      >
        <li>{result}</li>
      </div>
    {/each}
  </ul>
</div>
{/if}

<svelte:window
  on:keyup={(e) => {
    handleAutoCompTO(e)
  }}
/>
