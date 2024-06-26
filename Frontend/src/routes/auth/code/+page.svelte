<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  let code = "";
  let password = "";
  let passwordError = "";

  const validatePassword = () => {
    const regex =
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    passwordError = regex.test(password)
      ? ""
      : "Falsches Passwort! Das Passwort muss mindestens einen Großbuchstaben, einen Kleinbuchstaben, eine Zahl, ein Sonderzeichen enthalten und mindestens 8 Zeichen lang sein.";
  };

  async function finalRegister() {
    const response = await fetch(`${API_URL}auth/register/${code}`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ password: password }),
    });

    if (response.ok) {
      Swal.fire({
        icon: "success",
        title: "Die Registrierung war erfolgreich",
      });
    } else {
      Swal.fire({
        icon: "error",
        title: "Error",
        text: "Die Registrierung ist fehlgeschlagen",
      });
    }
  }
</script>

<main class="flex items-center justify-center min-h-screen bg-gray-100 p-6">
  <div class="max-w-md w-full bg-white rounded-lg shadow-md">
    <div class="p-6">
      <h2 class="text-2xl font-bold mb-6 text-center">Registration</h2>
      <div class="mb-4">
        <label for="code" class="block text-sm font-medium text-gray-700"
          >Code</label
        >
        <input
          type="text"
          id="code"
          bind:value={code}
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        />
      </div>
      <div class="mb-6">
        <label for="password" class="block text-sm font-medium text-gray-700"
          >Password</label
        >
        <input
          type="password"
          id="password"
          on:input={validatePassword}
          bind:value={password}
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        />
        <p class="text-red-500 text-sm mt-2">{passwordError}</p>
      </div>
      <button
        class="w-full py-2 px-4 bg-indigo-600 text-white font-medium rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        on:click={finalRegister}
      >
        Register
      </button>
    </div>
  </div>
</main>
