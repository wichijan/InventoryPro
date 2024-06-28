<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";

  let username = "";
  let successMessage = "";
  let errorMessage = "";

  const sendResetLink = () => {
    if (username) {
      fetch(API_URL + "email-forget-password", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: username,
        }),
      }).then((response) => {
        if (response.ok) {
          successMessage =
            "Ein Link zum Zurücksetzen des Passworts wurde an Ihre E-Mail-Adresse gesendet.";
        } else {
          errorMessage =
            "Es ist ein Fehler aufgetreten. Bitte überprüfen Sie Ihren Nutzernamen.";
        }
      });
    } else {
      successMessage = "";
    }
  };
</script>

<div class="container mx-auto p-6">
  <h1 class="text-3xl font-bold mb-8 text-center">Forgot Password</h1>

  <div class="bg-tertiary shadow-lg rounded-lg p-6 max-w-md mx-auto">
    <p class="mb-4 text-gray-600 text-center">
      Bitte geben Sie Ihren Nutzernamen ein, um einen Link zum Zurücksetzen des
      Passworts zu erhalten.
    </p>

    <form on:submit|preventDefault={sendResetLink} class="space-y-4">
      <div>
        <label for="username" class="block text-gray-700">Nutzernamen:</label>
        <input
          type="text"
          id="username"
          bind:value={username}
          class="w-full px-4 py-2 border rounded-md"
          required
        />
      </div>
      {#if successMessage}
        <p class="text-green-500 text-center">{successMessage}</p>
      {/if}
      {#if errorMessage}
        <p class="text-red-500 text-center">{errorMessage}</p>
      {/if}
      <button
        type="submit"
        class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full"
        >Send Reset Link</button
      >
    </form>
  </div>
</div>

<style>
  .container {
    max-width: 600px;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
      "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji",
      "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
  }
</style>
