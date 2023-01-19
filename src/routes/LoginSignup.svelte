<script>
  import { location } from 'svelte-spa-router'
  import { loginUser, signupUser } from '../lib/utils/loginSignup.js'
  let pageType
  if ($location == '/login') {
    pageType = 'Login'
  } else if ($location == '/signup') {
    pageType = 'Sign Up'
  }

  let username = ''
  let password = ''
  let repassword = ''
  let error = false
  const DEF_ERR_MSG = 'Invalid username and/or password'
  let errmsg = ''
  let success = false
  let sucmsg = ''

  const handleSubmit = (e) => {
    e.preventDefault()
    console.log('handleSubmit')
    error = false
    if (username == '' || password == '') {
      error = true
      errmsg = 'Missing username and/or password'
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
          }
        })
      } else {
        error = true
        errmsg = 'Passwords do not match, please retype'
      }
    }
  }
  //Invalid username <span id="error-msg-second-line">and/or password</span>
</script>

<body>
  <main id="main-holder">
    <h1 id="login-header">
      {pageType}
    </h1>

    {#if error}
      <div id="login-error-msg-holder">
        <p id="login-error-msg">{errmsg}</p>
      </div>
    {/if}
    {#if success}
      <div id="login-success-msg-holder">
        <p id="login-success-msg">{sucmsg}</p>
      </div>
    {/if}

    <form id="login-form">
      <input
        bind:value={username}
        type="text"
        name="username"
        id="username-field"
        class="login-form-field"
        placeholder="Username"
      />
      <input
        bind:value={password}
        type="password"
        name="password"
        id="password-field"
        class="login-form-field"
        placeholder="Password"
      />
      {#if pageType == 'Sign Up'}
        <input
          bind:value={repassword}
          type="password"
          name="repassword"
          id="retype-password-field"
          class="login-form-field"
          placeholder="Re-type Password"
        />
      {/if}
      <input
        type="submit"
        on:click={handleSubmit}
        value={pageType}
        id="login-form-submit"
      />
    </form>
  </main>
</body>

<style>
  html {
    height: 100%;
  }

  body {
    height: 100%;
    margin: 0;
    font-family: Arial, Helvetica, sans-serif;
    display: grid;
    justify-items: center;
    align-items: center;
    background-color: #242933;
  }

  #main-holder {
    width: 50%;
    height: 70%;
    display: grid;
    justify-items: center;
    align-items: center;
    background-color: #2a303c;
    border-radius: 7px;
    box-shadow: 0px 0px 5px 2px black;
  }

  #login-success-msg-holder {
    width: 100%;
    height: 100%;
    display: grid;
    justify-items: center;
    align-items: center;
  }

  #login-success-msg {
    width: 23%;
    text-align: center;
    margin: 0;
    padding: 5px;
    font-size: 12px;
    font-weight: bold;
    color: #0f5704;
    border: 1px solid #0f5704;
    background-color: #9ded91;
  }

  #login-error-msg-holder {
    width: 100%;
    height: 100%;
    display: grid;
    justify-items: center;
    align-items: center;
  }

  #login-error-msg {
    width: 23%;
    text-align: center;
    margin: 0;
    padding: 5px;
    font-size: 12px;
    font-weight: bold;
    color: #8a0000;
    border: 1px solid #8a0000;
    background-color: #e58f8f;
  }

  #error-msg-second-line {
    display: block;
  }

  #login-form {
    align-self: flex-start;
    display: grid;
    justify-items: center;
    align-items: center;
  }

  .login-form-field::placeholder {
    color: white;
  }

  .login-form-field {
    border: none;
    border-bottom: 1px solid #3a3a3a;
    margin-bottom: 10px;
    border-radius: 3px;
    outline: none;
    padding: 0px 0px 5px 5px;
  }

  #login-form-submit {
    width: 100%;
    padding: 7px;
    border: none;
    border-radius: 5px;
    color: white;
    font-weight: bold;
    background-color: #3a3a3a;
    cursor: pointer;
    outline: none;
  }
</style>
