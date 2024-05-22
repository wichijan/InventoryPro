<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService.js";
  import Swal from "sweetalert2";

  let username: string = "";
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

    fetch(API_URL + "auth/login", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username,
        password: password,
      }),
    }).then((response) => {
      if (response.ok) {
        Swal.fire({
          position: "top-end",
          icon: "success",
          title: "Login successful!",
          showConfirmButton: false,
          timer: 1500,
        });
      } else {
        Swal.fire({
          position: "top-end",
          icon: "error",
          title: "Login failed!",
          showConfirmButton: false,
          timer: 1500,
        });
      }
    });
  }
</script>

<div class="min-h-screen flex justify-center items-center">
  <div
    class="max-w-xs w-full m-auto bg-tertiary rounded-xl shadow-md overflow-hidden md:max-w-2xl px-4 py-6"
  >
    <h2
      class="text-center text-4xl font-extrabold text-white tracking-widest hover:text-blue-300 mb-4 duration-300"
    >
      Login
    </h2>
    <form class="mt-8 space-y-6" on:submit={login}>
      <div>
        <label for="username" class="text-white text-lg font-bold">
          Benutzername
        </label>
        <div class="mt-1 rounded-md shadow-sm">
          <input
            bind:value={username}
            type="text"
            class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
            placeholder="Benutzername"
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
          <p class="text-red-500 text-sm mt-2">{passwordError}</p>
        </div>
      </div>
      <div class="flex justify-between">
        <div class="flex flex-row w-max gap-1">
          <input type="checkbox" class="mx-auto my-auto rounded" />
          <label for="remember" class="text-white text-sm font-bold my-auto"
            >Angemeldet bleiben?</label
          >
        </div>
        <a
          href="/auth/register"
          class="text-white text-sm font-bold hover:text-blue-300 duration-300"
          >Noch keinen Account?</a
        >
      </div>

      <button
        disabled={passwordError}
        type="submit"
        class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full"
      >
        Login
      </button>
    </form>
  </div>
</div>
