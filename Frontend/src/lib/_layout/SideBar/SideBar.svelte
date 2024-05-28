<script lang="ts">
  import { onMount } from "svelte";
  import SideBarIcon from "./SideBarIcon.svelte";
  import { writable, readable } from "svelte/store";

  import {
    HouseDoor,
    CardList,
    InfoCircle,
    PersonCircle,
    Gear,
    GraphUp,
    Safe,
  } from "svelte-bootstrap-icons";
  import { afterNavigate, beforeNavigate } from "$app/navigation";
  import { page } from "$app/stores";
  import { API_URL } from "$lib/_services/ShelfService";

  let isAdmin = false;

  let sidebarItems = [
    { icon: HouseDoor, text: "InventoryPro", href: "/", active: false },
    { icon: CardList, text: "Overview", href: "/overview", active: false },
    { icon: InfoCircle, text: "Detail", href: "/detail", active: false },
    { icon: Gear, text: "Settings", href: "/settings", active: false },
  ];

  $: sidebarItems = sidebarItems;

  let recentPages: any[] = [];

  $: recentPages = recentPages;

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
          isAdmin = data.UserType === "Admin";
          if (isAdmin) {
            sidebarItems.push({
              icon: Safe,
              text: "Admin",
              href: "/admin",
              active: false,
            });
          } else {
            sidebarItems.push({
              icon: GraphUp,
              text: "Dashboard",
              href: "/dashboard",
              active: false,
            });
          }
          sidebarItems = sidebarItems;
        });
      } else {
        sidebarItems.push({
          icon: PersonCircle,
          text: "Login oder Register",
          href: "/auth/login",
          active: false,
        });
        sidebarItems = sidebarItems;
      }
    });
  });

  afterNavigate(() => {
    addRecentPage($page);
    getActive();
  });
  beforeNavigate(() => {
    getActive();
  });

  function addRecentPage(page: any) {
    //check double if double move it to the top
    recentPages = recentPages.filter((item) => item.href !== page.url.pathname);

    if (recentPages.length > 5) {
      recentPages.pop();
    }
    let routeName = page.route.id
      .toString()
      .replaceAll("[id]", "-Subpage")
      .replaceAll("/", "");
    routeName = routeName.charAt(0).toUpperCase() + routeName.slice(1);

    let routeForFind = page.url.pathname.split("/").slice(0, 2).join("/");
    if (routeForFind === "/auth") {
      routeForFind = "/auth/login";
      routeName = "Login";
    }

    let recentPage = {
      icon: sidebarItems.find((item) => item.href === routeForFind)?.icon,
      text: routeName,
      href: page.url.pathname,
    };
    recentPages = [recentPage, ...recentPages];
  }

  function getActive() {
    sidebarItems = sidebarItems.map((item) => {
      item.active = $page.url.pathname.includes(item.href) && item.href !== "/";
      return item;
    });
    sidebarItems = sidebarItems;
  }
</script>

<div
  class="fixed top-0 left-0 h-screen w-16 m-0 flex flex-col bg-white text-white shadow-lg"
>
  {#each sidebarItems as item}
    <SideBarIcon
      icon={item.icon}
      text={item.text}
      href={item.href}
      active={item.active}
    />
  {/each}
  <!-- <hr />
  <p class="text-sm mx-auto mt-3">Recent:</p>
  {#each recentPages as item}
    <SideBarIcon
      icon={item.icon}
      text={item.text}
      href={item.href}
      active={false}
    />
  {/each} -->
</div>
