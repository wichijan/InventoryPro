<script lang="ts">
  import "../app.css";
  import { onMount } from "svelte";
  import { crossfade } from "svelte/transition";
  import { fade } from "svelte/transition";
  import { page } from "$app/stores";
  import Footer from "$lib/_layout/Footer.svelte";
  import Navbar from "$lib/_layout/Navbar.svelte";
  import { afterNavigate, beforeNavigate } from "$app/navigation";
  import { quintOut } from "svelte/easing";

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

  const [send, receive] = crossfade({
    duration: (d) => Math.sqrt(d * 200),

    fallback(node, params) {
      const style = getComputedStyle(node);
      const transform = style.transform === "none" ? "" : style.transform;

      return {
        duration: 600,
        easing: quintOut,
        css: (t) => `
          transform: ${transform} scale(${t});
          opacity: ${t}
        `,
      };
    },
  });
</script>

<div
  class="relative flex flex-col justify-between min-w-screen min-h-screen bg-gray-800"
>
  {#key url}
    <div out:send={{ key: url }}>
      {#if url !== "/"}
        <div class="sticky top-0 z-50">
          <Navbar />
        </div>
      {/if}
      <div class="flex-grow flex-1 overflow-hidden w-screen h-full">
        <main in:receive={{ key: url }}>
          <slot />
        </main>
      </div>
    </div>
  {/key}
  <div class="flex min-w-screen h-full"><Footer /></div>
</div>
