<script lang="ts">
  import {
    CardList,
    Gear,
    GraphUp,
    People,
    GraphUpArrow,
    ListCheck,
    DoorOpen,
  } from "svelte-bootstrap-icons";

  import Sidebar from "./Sidebar.svelte";

  import NavbarItemMobile from "./NavbarItemMobile.svelte";
  import { onMount } from "svelte";
  import { afterNavigate } from "$app/navigation";
  import { page } from "$app/stores";
  import { isUserAdmin, isUserLoggedIn } from "$lib/_services/UserService";

  let active = "";

  let items: any[] = [];

  $: items = items;
  onMount(() => {
    syncItems();
  });

  function syncItems() {
    active = $page.route.id || "";
    items = [
      {
        name: "Items",
        link: "/items",
        icon: ListCheck,
        active: active.includes("/items") && active.length === 1,
      },
      {
        name: "Overview",
        link: "/overview",
        icon: CardList,
        active: active.includes("overview"),
      },
    ];

    getLoginDashboardItems();

    items = items;
  }

  async function getLoginDashboardItems() {
    const isLoggedIn = await isUserLoggedIn();
    if (isLoggedIn) {
      const isAdmin = await isUserAdmin();
      if (isAdmin) {
        if (!items.some((item) => item.name === "Admin"))
          items.push({
            name: "Admin",
            link: "/admin",
            icon: GraphUpArrow,
            active: active.includes("admin"),
          });
      } else {
        if (!items.some((item) => item.name === "Dashboard"))
          items.push({
            name: "Dashboard",
            link: "/dashboard",
            icon: GraphUp,
            active: active.includes("dashboard"),
          });
      }
      if (!items.some((item) => item.name === "Logout"))
        items.push({
          name: "Logout",
          link: "/auth/logout",
          icon: DoorOpen,
          active: active.includes("logout"),
        });
    } else {
      if (!items.some((item) => item.name === "Login"))
        items.push({
          name: "Login",
          link: "/auth/login",
          icon: People,
          active: active.includes("login"),
        });
    }
    items = items;
  }

  afterNavigate(() => {
    syncItems();
  });
</script>

<!-- Design for Mobile-->
<div class="sm:hidden">
  <div
    class="fixed bottom-0 w-full h-16 bg-[#e3d5ca] flex justify-around items-center shadow-lg"
  >
    {#key items}
      {#if items.length > 0}
        {#each items as item}
          <NavbarItemMobile {item} />
        {/each}
      {/if}
    {/key}
  </div>
</div>
<!-- Design for else -->
<div class="hidden sm:block">
  {#key items}
    {#if items.length > 0}
      <Sidebar {items} />
    {/if}
  {/key}
</div>
