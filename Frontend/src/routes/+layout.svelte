<script lang="ts">
  import "../app.css";
  import Footer from "$lib/_layout/Footer.svelte";
  import { crossfade } from "svelte/transition";
  import { quadInOut } from "svelte/easing";
  import { page } from "$app/stores";
  import SideBar from "$lib/_layout/SideBar/SideBar.svelte";
  import { afterNavigate, beforeNavigate } from "$app/navigation";
  import { onMount } from "svelte";

  let url: string;

  $: url = $page.url.pathname;

  let breadcrumbs: any[] = [];
  $: breadcrumbs = breadcrumbs;

  afterNavigate(() => {
    url = $page.url.pathname;
    setBreadcrumbs();
  });
  beforeNavigate(() => {
    url = $page.url.pathname;
    setBreadcrumbs();
  });

  onMount(() => {
    setBreadcrumbs();
  });

  function setBreadcrumbs() {
    breadcrumbs = [];
    let urlArray = url.split("/");
    urlArray.forEach((part) => {
      if (part === "auth") return;
      if (part === "login" || part === "register") {
        breadcrumbs.push({
          text:
            part.toString().charAt(0).toUpperCase() + part.toString().slice(1),
          href: `/auth/${part}`,
        });
        breadcrumbs = breadcrumbs;
        return;
      }
      if (part !== "") {
        breadcrumbs.push({
          text:
            part.toString().charAt(0).toUpperCase() + part.toString().slice(1),
          href: `/${part}`,
        });
      } else {
        breadcrumbs.push({ text: "Home", href: "/" });
      }
      breadcrumbs = breadcrumbs;
    });
  }
</script>

<div
  class="relative flex flex-col justify-between min-w-screen min-h-screen bg-primary"
>
  <div class="">
    {#if url !== "/"}
      <div class="sticky top-0">
        <SideBar />
      </div>
      <div
        class=" mt-3 bg-tertiary px-5 py-1 rounded-md ml-[5.25rem] mr-5"
        id="breadcrumbs"
      >
        {#key breadcrumbs}
          {#each breadcrumbs as breadcrumb, index}
            <a
              href={breadcrumb.href}
              class="text-[#344e41] hover:text-blue-500 duration-300"
              >{breadcrumb.text}</a
            >
            {index < breadcrumbs.length - 1 ? "/ " : " "}
          {/each}
        {/key}
        <!-- <a href="/" class="text-[#344e41] hover:text-blue-500 duration-300"
          >Home</a
        >
        /{" "}
        <a
          href="/detail"
          class="text-[#344e41] hover:text-blue-500 duration-300">Detail</a
        >
        /{" "} -->
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
