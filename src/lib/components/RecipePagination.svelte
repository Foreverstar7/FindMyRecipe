<script>
  import { createEventDispatcher } from 'svelte';
  export let page;
  export let totalPages;

  $: totalRange = [...Array(7).keys()].map((i) => i + page - 3);
  $: if (totalRange[0] <= 0 ) {
    totalRange = totalRange.map((i) => i + 1 - totalRange[0]);
  }

  $: prevPages = totalRange.filter((i) => i < page);
  $: nextPages = totalRange.filter((i) => i > page && i <= totalPages); 

  const dispatch = createEventDispatcher();

  function handlePreviousPage() {
    dispatch('pageChange', { page: page - 1 })
  }

  function handleChangePage(newPage) {
    dispatch('pageChange', { page: newPage })
  }

  function handleNextPage() {
    dispatch('pageChange', { page: page + 1 })
  }
</script>

<div class="flex flex-row justify-center md:text-xl">
  {#if page > 4}
    <button
      class="btn font-semibold mx-4 bg-base-100"
      on:click={handlePreviousPage}
    >
      Prev
    </button>
  {/if}
  {#each prevPages as i}
    <button
      class="btn mx-4 bg-base-100"
      on:click={() => handleChangePage(i)}
    >
      {i}
    </button>
  {/each}
  <button class="btn no-animation font-extrabold mx-4 underline bg-base-100">
    {page}
  </button>
  {#each nextPages as i}
    <button
      class="btn mx-4 bg-base-100"
      on:click={() => handleChangePage(i)}
    >
      {i}
    </button>
  {/each}
    {#if page < totalPages - 3}
    <button
      class="btn font-semibold mx-4 bg-base-100"
      on:click={handleNextPage}
    >
      Next
    </button>
  {/if}
</div>
