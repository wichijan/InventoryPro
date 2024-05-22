<script>
  import { browser } from "$app/environment";
  import { API_URL } from "$lib/_services/ShelfService";

  import { onMount } from "svelte";

  onMount(() => {
    fetch(API_URL + "users/get-me", {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((response) => {
      if (response.ok) {
        response.json().then((data) => {
          if (data.UserType === "Admin") {
            browser ? (window.location.href = "/admin") : null;
          } else {
            browser ? (window.location.href = "/dashboard") : null;
          }
        });
      } else {
        browser ? (window.location.href = "/auth/login") : null;
      }
    });
  });
</script>

<slot />
