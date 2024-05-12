<script lang="ts">
  import "../app.css";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import { fade } from "svelte/transition";
  import { page } from "$app/stores";
  import Footer from "$lib/_layout/Footer.svelte";
  import Navbar from "$lib/_layout/Navbar.svelte";
  import { afterNavigate, beforeNavigate } from "$app/navigation";

  let url: string;

  onMount(() => {
    url = $page.url.pathname;
  });
  $: {
    url = $page.url.pathname;
  }

  beforeNavigate(() => {
    url = $page.url.pathname;
  });

  afterNavigate(() => {
    url = $page.url.pathname;
  });
</script>

<div
  class="relative flex flex-col justify-between min-w-screen min-h-screen bg-gray-800"
>
  {#key url}
    <div transition:fade={{ delay: 100 }}>
      {#if url !== "/"}
        <div class="sticky top-0 z-50">
          <Navbar />
        </div>
      {/if}
      <div class="flex-grow flex-1 overflow-hidden w-screen h-full">
        <slot />
      </div>
    </div>
  {/key}
  <div class="flex min-w-screen h-full"><Footer /></div>
</div>
