<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import { Eye, EyeSlash } from "svelte-bootstrap-icons";
  import Swal from "sweetalert2";

  let username: string = "";
  let password: string = "";
  let passwordError: string = "";
  let passwordVisible = false;
  const validatePassword = () => {
    const regex =
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    passwordError = regex.test(password)
      ? ""
      : "Falsches Passwort! Das Passwort muss mindestens einen Großbuchstaben, einen Kleinbuchstaben, eine Zahl, ein Sonderzeichen enthalten und mindestens 8 Zeichen lang sein.";
  };

  function typeCheck(node) {
    node.type = passwordVisible ? "text" : "password";
  }

  function togglePasswordVisibility() {
    let pElement = document.getElementById("password");

    passwordVisible = pElement.type === "password";
    if (pElement.type === "password") {
      pElement.type = "text";
    } else {
      pElement.type = "password";
    }
  }
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
        goto("/dashboard");
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
        <div class="mt-1 rounded-md shadow-sm relative">
          {#key passwordVisible}
            <input
              bind:value={password}
              on:input={validatePassword}
              id="password"
              use:typeCheck
              class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
              placeholder="Password"
              required
            />
          {/key}
          <button
            type="button"
            on:click={togglePasswordVisibility}
            class="absolute right-2 top-1/2 transform -translate-y-1/2 bg-transparent border-none text-gray-500 focus:outline-none"
          >
            {#if passwordVisible}
              <EyeSlash class="h-6 w-6" />
            {:else}
              <Eye class="h-6 w-6" />
            {/if}
          </button>
        </div>
        <p class="text-red-500 text-sm mt-2">{passwordError}</p>
      </div>
      <div class="flex justify-between">
        <div class="flex flex-row w-max gap-1">
          <input type="checkbox" class="mx-auto my-auto rounded" />
          <label for="remember" class="text-white text-sm font-bold my-auto"
            >Angemeldet bleiben?</label
          >
        </div>
        <a
          href="/auth/forgot-password"
          class="text-white text-sm font-bold hover:text-blue-400 duration-300"
          >Passwort vergessen?</a
        >

        <a
          href="/auth/register"
          class="text-white text-sm font-bold hover:text-blue-400 duration-300"
          >Noch keinen Account?</a
        >
      </div>

      <button
        type="submit"
        class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full"
      >
        Login
      </button>
    </form>
  </div>
</div>
