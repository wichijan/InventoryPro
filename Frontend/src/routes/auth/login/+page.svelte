<script lang="ts">
  let email = "";
  let password = "";
  let passwordError = "";

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

<div class="min-h-screen flex justify-center items-center">
  <div
    class="max-w-xs w-full m-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl"
  >
    <h2 class="text-center py-5 text-4xl font-bold text-blue-800 mb-4">
      Login
    </h2>
    <form class="py-8 px-4 md:p-8" on:submit={login}>
      <div class="my-5 text-sm">
        <label for="username" class="block text-black"
          >Email <span class="-ml-1 text-red-500">*</span></label
        >
        <input
          bind:value={email}
          type="text"
          class="w-full px-5 py-3 border rounded-lg focus:outline-none"
          placeholder="Email"
          required
        />
      </div>
      <div class="my-5 text-sm;">
        <label for="password" class="block text-black"
          >Password <span class="-ml-1 text-red-500">*</span></label
        >
        <input
          bind:value={password}
          on:input={validatePassword}
          type="password"
          class="w-full px-5 py-3 border rounded-lg focus:outline-none"
          placeholder="Password"
          required
        />
        <p class="text-red-500 text-sm mt-2">{passwordError}</p>
      </div>
      <button
        type="submit"
        disabled={passwordError}
        class="block w-full rounded-lg bg-gradient-to-r from-purple-500 to-blue-500 py-3 text-white text-center font-semibold text-sm focus:outline-none"
        >Login</button
      >
    </form>
  </div>
</div>
