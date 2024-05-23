<script>
  import { browser } from "$app/environment";
  import { API_URL } from "$lib/_services/ShelfService";

  import { onMount } from "svelte";

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
          if (data.loggedIn) {
            browser ? (window.location.href = "/dashboard") : null;
          }
        });
      }
    });
  });
</script>

<slot />
