<script lang="ts">
  import "../app.css";
  import Footer from "$lib/_layout/Footer.svelte";
  import { crossfade } from "svelte/transition";
  import { quadInOut } from "svelte/easing";
  import { page } from "$app/stores";
  import SideBar from "$lib/_layout/SideBar/SideBar.svelte";

  let url: string;

  $: url = $page.url.pathname;

  const [send, receive] = crossfade({
    duration: (d) => Math.sqrt(d * 800),

    fallback(node, params) {
      const style = getComputedStyle(node);
      const transform = style.transform === "none" ? "" : style.transform;

      return {
        duration: 600,
        easing: quadInOut,
        css: (t) => `
          transform: ${transform} scale((${1.0 - t} * 0.5) + ${t});
          opacity: ${t}
        `,
      };
    },
  });
</script>

<div
  class="relative flex flex-col justify-between min-w-screen min-h-screen bg-primary"
>
  <div>
    {#if url !== "/"}
      <div class="sticky top-0 z-50">
        <SideBar />
      </div>
    {/if}
    <div class="flex-grow flex-1 overflow-hidden h-full pl-11">
      <main>
        <slot />
      </main>
    </div>
  </div>
  <div class="flex min-w-screen h-full"><Footer /></div>
</div>
