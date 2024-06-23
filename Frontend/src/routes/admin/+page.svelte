<script lang="ts">
  import { goto } from "$app/navigation";

  let dashboardItems = [
    {
      name: "Items",
      path: "/admin/items",
      description: "Manage items",
    },
    {
      name: "Users",
      path: "/admin/users",
      description: "Manage users",
    },
    {
      name: "Warehouses",
      path: "/admin/warehouses",
      description: "Manage warehouses",
    },
  ];
</script>

<div class="flex flex-col items-center w-ful mt-5">
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
        <img
          class="mx-auto rounded w-12 h-12 object-cover"
          src="https://via.placeholder.com/150"
          alt="Image"
        />
        <div class="px-6 py-4">
          <div class="font-bold text-xl mb-2">{item.name}</div>
          <p class="text-gray-700 text-base">{item.description}</p>
        </div>
      </button>
    {/each}
  </div>
</div>
