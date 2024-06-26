<script lang="ts">
  import { goto } from "$app/navigation";
  import {
    PersonCircle,
    ListNested,
    BuildingFill,
    PeopleFill,
    ArchiveFill,
    Backpack2Fill,
    FilesAlt,
    Speedometer2,
  } from "svelte-bootstrap-icons";

  let dashboardItems = [
    {
      name: "Benutzer",
      path: "/admin/users",
      icon: PersonCircle,
      description: "Benutzer verwalten",
    },
    {
      name: "Warehouses",
      path: "/admin/warehouses",
      icon: BuildingFill,
      description: "Warehouses verwalten",
    },
    {
      name: "Keywords",
      path: "/admin/keywords",
      icon: ListNested,
      description: "Keywords verwalten",
    },
    {
      name: "Rollen",
      path: "/admin/roles",
      icon: PeopleFill,
      description: "Rollen verwalten",
    },
    {
      name: "Fächer",
      path: "/admin/subjects",
      icon: Backpack2Fill,
      description: "Fächer verwalten",
    },
    {
      name: "Benutzertypen",
      path: "/admin/usertypes",
      icon: ArchiveFill,
      description: "Benutzertypen verwalten",
    },
    {
      name: "Items",
      path: "/admin/items",
      icon: FilesAlt,
      description: "Elemente verwalten",
    },
    {
      name: "Schnellregale",
      path: "/admin/quickshelf",
      icon: Speedometer2,
      description: "Schnellregale verwalten",
    },
  ];
</script>

<div class="flex flex-col items-center w-ful mt-5 mb-12">
  <div class="my-3 flex flex-col">
    <h1 class="text-4xl font-bold text-gray-900 mx-auto">Admin Dashboard</h1>
    <p class="text-gray-700 text-base mt-2">
      Willkommen auf dem Admin-Dashboard. Hier können Sie Benutzer und
      Lagerhäuser verwalten.
    </p>
  </div>
  <div
    class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 bg-tertiary rounded px-2 py-4"
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
