<script>
  import { browser } from "$app/environment";
  import { isUserLoggedIn } from "$lib/_services/UserService";

  import { onMount } from "svelte";

  onMount(() => {
    isUserLoggedIn().then((isLoggedIn) => {
      if (isLoggedIn) {
        browser ? (window.location.href = "/dashboard") : null;
      }
    });
  });
</script>

<slot />
