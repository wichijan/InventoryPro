<script lang="ts">
  import { browser } from "$app/environment";
  import { API_URL } from "$lib/_services/ShelfService";
  import { onMount } from "svelte";

  let isLoggedIn = false;

  onMount(() => {
    fetch(API_URL + "auth/logged-in", {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((response) => {
      if (response.ok) {
        response.json().then((data) => {
          isLoggedIn = !data.loggedIn;
        });
      }
    });
  });
</script>

{#if isLoggedIn}
  <div class="flex flex-col text-left">
    <div class="font-semibold text-lg">Password</div>
    <input
      type="password"
      class="border-2 border-[#d5bdaf] rounded-md px-3 py-1 mt-1 w-full
      focus:outline-none focus:ring-2 focus:ring-[#d5bdaf] focus:border-transparent"
      placeholder="Password"
    />
    <div class="font-semibold text-lg mt-3">New Password</div>
    <input
      type="password"
      class="border-2 border-[#d5bdaf] rounded-md px-3 py-1 mt-1 w-full
      focus:outline-none focus:ring-2 focus:ring-[#d5bdaf] focus:border-transparent"
      placeholder="New Password"
    />
    <button
      class="bg-[#d5bdaf] hover:bg-d6ccc2 enabled:hover:text-black enabled:hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full
      disabled:opacity-50 disabled:cursor-not-allowed"
    >
      Speichern
    </button>
    <div class="text-sm mt-2 text-red-500 mx-auto">Änderungen speichern!</div>
  </div>
{:else}
  <div class="container">
    <div class="row">
      <div class="col-12">
        <h1 class="text-center">Please log in to view this page</h1>
      </div>
    </div>
  </div>
{/if}
