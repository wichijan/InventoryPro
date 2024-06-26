<script lang="ts">
  import { connect, state } from "$lib/_services/WebSocket";
  import {
    CardList,
    Speedometer2,
    GraphUp,
    People,
    GraphUpArrow,
    ListCheck,
    DoorOpen,
    Bell,
  } from "svelte-bootstrap-icons";

  import Sidebar from "./Sidebar.svelte";

  import NavbarItemMobile from "./NavbarItemMobile.svelte";
  import { onMount } from "svelte";
  import { afterNavigate } from "$app/navigation";
  import { page } from "$app/stores";
  import { isUserAdmin, isUserLoggedIn } from "$lib/_services/UserService";

  let active = "";

  let items: any[] = [];

  let notificationCount = 0;
  $: notificationCount = notificationCount;

  state.subscribe((value) => {
    console.log(value);
    notificationCount = value.requests.length;
    items = items.map((item) => {
      if (item.name === "Notifications") {
        item.nCount = notificationCount;
      }
      return item;
    });
  });

  $: items = items;
  onMount(() => {
    syncItems();
    connect();
  });

  function syncItems() {
    active = $page.route.id || "";
    items = [
      {
        name: "Items",
        link: "/items",
        icon: ListCheck,
        active: active.includes("items"),
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
      if (!items.some((item) => item.name === "Schnellregal"))
        items.push({
          name: "Schnellregal",
          link: "/quickshelves",
          icon: Speedometer2,
          active: active.includes("quickshelves"),
        });

      if (!items.some((item) => item.name === "Notifications"))
        items.push({
          name: "Notifications",
          link: "/notifications",
          icon: Bell,
          active: active.includes("notifications"),
          nCount: notificationCount,
        });

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
