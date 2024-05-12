<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  //return to main page
  onMount(() => {
    window.location.href = "/";
  });
</script>
