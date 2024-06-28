<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";

  export let data;

  let userid = data.userID;

  let password = "";
  let confirmPassword = "";
  let successMessage = "";
  let errorMessage = "";

  $: errorMessage = errorMessage;
  const validatePassword = () => {
    const regex =
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    errorMessage = regex.test(password)
      ? ""
      : "Falsches Passwort! Das Passwort muss mindestens einen Großbuchstaben, einen Kleinbuchstaben, eine Zahl, ein Sonderzeichen enthalten und mindestens 8 Zeichen lang sein.";
  };

  const resetPassword = () => {
    if (password === confirmPassword) {
      fetch(API_URL + "request-forgot-password", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          userId: userid,
          password: password,
        }),
      }).then((response) => {
        if (response.ok) {
          successMessage = "Das Passwort wurde erfolgreich zurückgesetzt.";
          setTimeout(() => {
            goto("/auth/login");
          }, 2000);
        } else {
          errorMessage =
            "Es ist ein Fehler aufgetreten. Bitte versuchen Sie es erneut.";
        }
      });
    } else {
      errorMessage = "Die Passwörter stimmen nicht überein.";
    }
  };
</script>

<div class="container mx-auto p-6">
  <h1 class="text-3xl font-bold mb-8 text-center">Password zurücksetzen</h1>

  <div class="bg-gray-100 shadow-lg rounded-lg p-6 max-w-md mx-auto">
    <p class="mb-4 text-gray-600 text-center">
      Bitte geben Sie Ihr neues Passwort ein, um Ihr Passwort zurückzusetzen.
    </p>

    <form on:submit|preventDefault={resetPassword} class="space-y-4">
      <div>
        <label for="password" class="block text-gray-700">Neues Passwort:</label
        >
        <input
          type="password"
          id="password"
          on:input={validatePassword}
          bind:value={password}
          class="w-full px-4 py-2 border rounded-md"
          required
        />
      </div>
      <div>
        <label for="confirmPassword" class="block text-gray-700"
          >Bestätigen Sie das Passwort:</label
        >
        <input
          type="password"
          id="confirmPassword"
          bind:value={confirmPassword}
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
        class="w-full bg-blue-600 text-white py-2 rounded-md"
        >Passwort zurücksetzen</button
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
