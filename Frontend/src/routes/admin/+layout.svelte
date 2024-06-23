<script>
  import { goto } from "$app/navigation";
  import { isUserAdmin } from "$lib/_services/UserService";

  import { onMount } from "svelte";

  onMount(async () => {
    await isUserAdmin().then((res) => {
      if (!res) goto("/dashboard");
    });
  });
</script>

<slot />
