<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import { onMount } from "svelte";
  import { isUserLoggedIn } from "$lib/_services/UserService";

  onMount(async () => {
    const isLoggedIn = await isUserLoggedIn();
    if (!isLoggedIn) {
      goto("/auth/login");
    }
    fetch(`${API_URL}auth/logout`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    }).then((response) => {
      if (response.ok) {
        goto("/auth/login");
      }
    });
  });
</script>
