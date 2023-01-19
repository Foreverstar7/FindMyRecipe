<script>
  import { logout, refresh } from '../utils/loginSignup.js'
  import { getUserInfo, RECIPE_NUMS } from '../utils/utils.js'

  let userinfo = getUserInfo()
  var checkCookie = (function () {
    return function () {
      return
      // TODO: make refresh work
      let currentCookie = getUserInfo()
      // something useful like parse cookie, run a callback fn, etc.
      let now = Math.round(new Date().getTime() / 1000)
      let fiveMinutes = 60 * 5
      if (currentCookie && now > currentCookie.exp - fiveMinutes) {
        refresh(currentCookie).then((res) => {
          console.log('refreshed', getUserInfo())
        })
      }
    }
  })()

  window.setInterval(checkCookie, 1000)

  let onLogin = userinfo == null ? false : true

  const handleLogout = () => {
    logout().then((res) => {
      console.log('logout successfully')
      window.location.href = '#/login'
      window.location.reload()
    })
  }

  const getRandom = () => {
    const id = 1 + Math.floor(Math.random() * RECIPE_NUMS)
    window.location.href = `#/recipes/id/${id}`
    window.location.reload()
  }
</script>

<div class="navbar bg-base-100 justify-between flex">
  <div class="flex">
    <a class="btn btn-ghost normal-case text-xl" href="#/"
      ><img
        src="https://cdn-icons-png.flaticon.com/512/526/526190.png"
        class="w-6 h-6 p-1"
        alt="chefhat"
      />Find My Recipe</a
    >
    <a class="btn btn-ghost normal-case" href="#/recipes">
      <svg
        class="w-6 h-6 p-1"
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="#ffffff"
        stroke-width="2"
        stroke-linecap="butt"
        stroke-linejoin="bevel"
        ><rect x="3" y="3" width="7" height="7" /><rect
          x="14"
          y="3"
          width="7"
          height="7"
        /><rect x="14" y="14" width="7" height="7" /><rect
          x="3"
          y="14"
          width="7"
          height="7"
        /></svg
      >
      Recipes</a
    >
    <a class="btn btn-ghost normal-case" href="#/recipes/weekly">
      <svg
        class="w-6 h-6 p-1"
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="#ffffff"
        stroke-width="2"
        stroke-linecap="butt"
        stroke-linejoin="bevel"
        ><path d="M20.2 7.8l-7.7 7.7-4-4-5.7 5.7" /><path d="M15 7h6v6" /></svg
      >
      Trending</a
    >
    <!-- Top Weekly Recipes -->
    {#if onLogin}
      <a class="btn btn-ghost normal-case" href="#/recipes/favorite">
        <svg
          class="w-6 h-6 p-1"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="#ffffff"
          stroke-width="2"
          stroke-linecap="butt"
          stroke-linejoin="bevel"
          ><polygon
            points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"
          /></svg
        >
        Favorites</a
      >
      <a class="btn btn-ghost normal-case" href="#/recipes/recommended">
        <svg
          class="w-6 h-6 p-1"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="#ffffff"
          stroke-width="2"
          stroke-linecap="butt"
          stroke-linejoin="bevel"
          ><path
            d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"
          /></svg
        >
        Recommended</a
      >
      <a class="btn btn-ghost normal-case" href="#/recipes/add">
        <svg
          class="w-6 h-6 p-1"
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="#ffffff"
          stroke-width="2"
          stroke-linecap="butt"
          stroke-linejoin="bevel"
          ><path
            d="M20 14.66V20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h5.34"
          /><polygon points="18 2 22 6 12 16 8 16 8 12 18 2" /></svg
        >
        Add</a
      >
    {/if}
  </div>
  <div class="flex">
    <div tabindex="0" class="btn btn-ghost btn-circle" on:click={getRandom}>
      <div class="rounded-full">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="#ffffff"
          stroke-width="2"
          stroke-linecap="butt"
          stroke-linejoin="bevel"
          ><path
            d="M16 3h5v5M4 20L20.2 3.8M21 16v5h-5M15 15l5.1 5.1M4 4l5 5"
          /></svg
        >
      </div>
    </div>
    {#if onLogin}
      <div class="dropdown dropdown-end">
        <div tabindex="0" class="btn btn-ghost btn-circle">
          <div class="rounded-full">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="#ffffff"
              stroke-width="2"
              stroke-linecap="butt"
              stroke-linejoin="bevel"
              ><line x1="3" y1="12" x2="21" y2="12" /><line
                x1="3"
                y1="6"
                x2="21"
                y2="6"
              /><line x1="3" y1="18" x2="21" y2="18" /></svg
            >
          </div>
        </div>
        <ul
          tabindex="0"
          class="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
        >
          <li class="prose">{userinfo.username}</li>
          <li>
            <a class="btn btn-ghost normal-case" href='#/recipes/author/{userinfo.userid}'
              >Your Recipes</a
            >
            <button class="btn btn-ghost normal-case" on:click={handleLogout}
              >Logout</button
            >
          </li>
        </ul>
      </div>
    {:else}
      <a class="btn" href="#/signup">Get Started</a>
    {/if}
  </div>
</div>
