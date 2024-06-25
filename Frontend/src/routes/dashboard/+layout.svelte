<script lang="ts">
  import { browser } from "$app/environment";
  import { isUserLoggedIn, isUserAdmin } from "$lib/_services/UserService";
  import { onMount } from "svelte";

  onMount(() => {
    isUserLoggedIn().then(async (isLoggedIn) => {
      if (!isLoggedIn) {
        browser ? (window.location.href = "/auth/login") : null;
      } else {
        const isAdmin = await isUserAdmin();
        if (isAdmin) {
          browser ? (window.location.href = "/admin") : null;
        }
      }
    });
  });
</script>

<slot />
