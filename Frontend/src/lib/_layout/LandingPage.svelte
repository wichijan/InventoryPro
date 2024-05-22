<script lang="ts">
  import { onMount } from "svelte";

  let isAdmin = Math.random() > 0.5;

  let navbarItems = [
    { color: "bg-blue-500", text: "Home", href: "/" },
    { color: "bg-red-500", text: "Overview", href: "/overview" },
    { color: "bg-yellow-500", text: "Detail", href: "/detail" },
  ];

  $: navbarItems = navbarItems;

  onMount(() => {
    if (isAdmin) {
      navbarItems.push({
        color: "bg-green-500",
        text: "Admin",
        href: "/admin",
      });
    } else {
      navbarItems.push({
        color: "bg-green-500",
        text: "Login",
        href: "/auth/login",
      });
    }
    navbarItems = navbarItems;
  });
</script>

<div
  class="min-h-screen bg-primary flex items-center justify-between px-8 py-6"
>
  <div class="space-x-4">
    <a
      href="/"
      class="text-2xl font-extrabold hover:text-black tracking-widest transition transform duration-500 ease-in-out hover:scale-105 text-blue-500"
    >
      InventoryPro
    </a>
  </div>

  <div class="space-x-8">
    {#each navbarItems as item}
      <a
        href={item.href}
        class="text-lg font-bold px-5 py-3 rounded-lg transform transition duration-500 ease-in-out hover:scale-105 text-black hover:text-white hover:bg-blue-500"
      >
        {item.text}
      </a>
    {/each}
  </div>
</div>
