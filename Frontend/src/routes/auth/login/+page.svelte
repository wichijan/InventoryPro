<script lang="ts">
  let email: string = "";
  let password: string = "";
  let passwordError: string = "";

  const validatePassword = () => {
    const regex =
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    passwordError = regex.test(password)
      ? ""
      : "Invalid password! Password should contain at least one uppercase letter, one lowercase letter, one number, one special character and should be at least 8 characters long.";
  };

  function login(e: any) {
    e.preventDefault();
    console.log(email);
    console.log(password);
  }
</script>

<div class="min-h-screen bg-gray-800 flex justify-center items-center">
  <div
    class="max-w-xs w-full m-auto bg-gray-700 rounded-xl shadow-md overflow-hidden md:max-w-2xl px-4 py-6"
  >
    <h2
      class="text-center text-4xl font-extrabold text-white tracking-widest hover:text-blue-300 mb-4 duration-300"
    >
      Login
    </h2>
    <form class="mt-8 space-y-6" on:submit={login}>
      <div>
        <label for="username" class="text-white text-lg font-bold">
          Email
        </label>
        <div class="mt-1 rounded-md shadow-sm">
          <input
            bind:value={email}
            type="text"
            class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
            placeholder="Email"
            required
          />
        </div>
      </div>
      <div>
        <label for="password" class="text-white text-lg font-bold">
          Password
        </label>
        <div class="mt-1 rounded-md shadow-sm">
          <input
            bind:value={password}
            on:input={validatePassword}
            type="password"
            class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
            placeholder="Password"
            required
          />
          <p class="text-red-600 text-sm mt-2">{passwordError}</p>
        </div>
      </div>
      <!-- Need to be registerd and remember me-->
      <div class="flex justify-between">
        <div class="flex flex-row w-max gap-1">
          <input type="checkbox" class="mx-auto my-auto rounded" />
          <label for="remember" class="text-white text-sm font-bold my-auto"
            >Remember me</label
          >
        </div>
        <a
          href="/auth/register"
          class="text-white text-sm font-bold hover:text-blue-300 duration-300"
          >Need to register?</a
        >
      </div>

      <button
        type="submit"
        disabled={passwordError}
        class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-lg font-bold text-white bg-purple-500 hover:bg-purple-500 hover:text-black rounded-md duration-300"
        >Login</button
      >
    </form>
  </div>
</div>
