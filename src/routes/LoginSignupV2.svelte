<script>
  import { location } from 'svelte-spa-router'
  import { loginUser, signupUser } from '../lib/utils/loginSignup.js'
  import { USER_FEATURES } from '../lib/utils/utils.js'
  let pageType
  if ($location == '/login') {
    pageType = 'Login'
  } else if ($location == '/signup') {
    pageType = 'Sign Up'
  }
  let username = ''
  let password = ''
  let repassword = ''
  let usernameError = false
  let passwordError = false
  let mismatchPwdError = false
  let error = false
  let signupModalActive = false
  let loading = false

  const DEF_ERR_MSG = 'Invalid username and/or password'
  let errmsg = ''
  let success = false
  let sucmsg = ''
  $: inputError = error ? 'input-error' : ''
  $: inputUsernameError = usernameError ? 'input-error' : ''
  $: inputPasswordError = passwordError ? 'input-error' : ''
  $: inputMismatch = mismatchPwdError ? 'input-error' : ''

  const handleSignupModal = (path) => {
    window.location.href = `#/${path}`
    window.location.reload()
  }

  const handleSubmit = (e) => {
    e.preventDefault()
    loading = true
    console.log('handleSubmit')
    error = false
    mismatchPwdError = false
    usernameError = false
    passwordError = false
    if (username == '' || password == '') {
      if (username == '') {
        usernameError = true
      }
      if (password == '') {
        passwordError = true
        mismatchPwdError = true
      }
      errmsg = 'Missing username and/or password'
      console.log(errmsg)
      loading = false
      return
    }
    if (pageType == 'Login') {
      loginUser(username, password).then((res) => {
        console.log(res)
        if (res.status == 'failed') {
          error = true
          errmsg = res.message
        } else {
          success = true
          sucmsg = 'Login success! You will be redirected'
          loading = false
          window.location.href = '#/'
          window.location.reload()
        }
      })
    } else if (pageType == 'Sign Up') {
      if (password == repassword) {
        signupUser(username, password).then((res) => {
          console.log(res)
          if (res.status == 'failed') {
            error = true
            errmsg = res.message
          } else {
            success = true
            sucmsg = 'Sign Up Success! You can now Login'
            signupModalActive = true
          }
        })
      } else {
        mismatchPwdError = true
        errmsg = 'Passwords do not match, please retype'
      }
    }
    console.log(errmsg)
    loading = false
  }
</script>

<div class="hero bg-base-200" style="min-height: 85vh;">
  <div class="hero-content w-full flex-col justify-center lg:flex-row">
    <div class="flex flex-col basis-1/3 text-center lg:text-left">
      <h1 class="text-5xl font-bold">{pageType} Now!</h1>
      <p class="py-6">
        {#each USER_FEATURES as feat}
          {feat}<br />
        {/each}
      </p>
    </div>
    <div class="card flex-shrink-0 max-w-md basis-1/2 shadow-2xl bg-base-100">
      <div class="card-body">
        <div class="form-control">
          <div class="label">
            <span class="label-text">Username</span>
          </div>
          <input
            bind:value={username}
            type="text"
            placeholder="username"
            class="input input-bordered {inputError} {inputUsernameError}"
          />
        </div>
        <div class="form-control">
          <div class="label">
            <span class="label-text">Password</span>
          </div>
          <input
            bind:value={password}
            type="password"
            placeholder="password"
            class="input input-bordered {inputError} {inputPasswordError} {inputMismatch}"
          />
          {#if pageType == 'Sign Up'}
            <div class="label">
              <span class="label-text">Re-type Password</span>
            </div>
            <input
              bind:value={repassword}
              type="password"
              placeholder="re-type password"
              class="input input-bordered {inputError} {inputMismatch}"
            />
            <div class="label">
              <a href="#/login" class="label-text-alt link link-hover"
                >Already have an account?</a
              >
            </div>
          {/if}
        </div>
        <div class="form-control mt-6">
          <button on:click={handleSubmit} class={loading ? "btn loading" : 'btn btn-primary'}>
            {pageType}
          </button
          >
        </div>
      </div>
      {#if signupModalActive}
        <input
          type="checkbox"
          id="my-modal"
          class="modal-toggle"
          bind:checked={signupModalActive}
        />
        <div class="modal">
          <div class="modal-box relative">
            <label
              for="my-modal"
              class="btn btn-sm btn-circle absolute right-2 top-2"
              on:click={() => handleSignupModal('')}>✕</label
            >
            <h3 class="font-bold text-lg">Sign Up Success!</h3>
            <p class="py-4">Now you can login or continue to browse freely!</p>
            <div class="modal-action">
              <label
                for="my-modal"
                class="btn"
                on:click={() => handleSignupModal('login')}>Login</label
              >
            </div>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>
