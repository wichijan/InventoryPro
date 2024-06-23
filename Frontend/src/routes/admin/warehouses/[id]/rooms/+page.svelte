<script>
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";

  browser ? history.back() : goto("/admin");
</script>
