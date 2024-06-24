<script lang="ts">
  import { goto } from "$app/navigation";
  import {
    PersonCircle,
    ListNested,
    BuildingFill,
    PeopleFill,
    ArchiveFill,
    Backpack2Fill,
  } from "svelte-bootstrap-icons";

  let dashboardItems = [
    {
      name: "Users",
      path: "/admin/users",
      icon: PersonCircle,
      description: "Manage users",
    },
    {
      name: "Warehouses",
      path: "/admin/warehouses",
      icon: BuildingFill,
      description: "Manage warehouses",
    },
    {
      name: "Keywords",
      path: "/admin/keywords",
      icon: ListNested,
      description: "Manage keywords",
    },
    {
      name: "Roles",
      path: "/admin/roles",
      icon: PeopleFill,
      description: "Manage roles",
    },
    {
      name: "Subjects",
      path: "/admin/subjects",
      icon: Backpack2Fill,
      description: "Manage Subjects",
    },
    {
      name: "Usertypes",
      path: "/admin/usertypes",
      icon: ArchiveFill,
      description: "Manage usertypes",
    },
  ];
</script>

<div class="flex flex-col items-center w-ful mt-5 mb-12">
  <div class="my-3 flex flex-col">
    <h1 class="text-4xl font-bold text-gray-900 mx-auto">Admin Dashboard</h1>
    <p class="text-gray-700 text-base mt-2">
      Welcome to the admin dashboard. Here you can manage users and warehouses.
    </p>
  </div>
  <div
    class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 gap-4 bg-tertiary rounded px-2 py-4"
  >
    {#each dashboardItems as item}
      <button
        class="max-w-sm rounded overflow-hidden shadow-lg m-3 bg-white px-5 py-5 hover:shadow-xl duration-300 ease-in-out transform hover:scale-[102%]"
        on:click={() => {
          goto(`${item.path}`);
        }}
      >
        <svelte:component
          this={item.icon}
          class="mx-auto rounded w-8 h-8 object-cover"
        />
        <div class="px-6 py-4">
          <div class="font-bold text-xl mb-2">{item.name}</div>
          <p class="text-gray-700 text-base">{item.description}</p>
        </div>
      </button>
    {/each}
  </div>
</div>
