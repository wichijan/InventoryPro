<script lang="ts">
  import { onMount } from "svelte";
  export let data;
  import { ArrowRight } from "svelte-bootstrap-icons";
  import { isUserAdmin } from "$lib/_services/UserService";
  let warehouses: any[] = data.warehouses;

  let rooms: any[] = data.rooms;

  let warehosesWithRooms = warehouses.map((warehouse) => {
    const warehouseRooms = rooms.filter(
      (room) => room.WarehouseID === warehouse.ID
    );

    return {
      ...warehouse,
      rooms: warehouseRooms,
    };
  });
</script>

<div
  class="grid grid-cols-{warehouses.length} ml-10 space-x-5 mt-8 bg-tertiary p-4 rounded-md mr-5"
>
  {#key warehosesWithRooms}
    {#each warehosesWithRooms as warehouse, index}
      <div
        class="{index !== warehouses.length - 1
          ? 'border-r'
          : ''} border-gray-600 p-4"
      >
        <div class="bg-white rounded-md p-4 hover:shadow-lg duration-300">
          <div class="px-4 py-2 ring-2 ring-gray-400 text-black rounded-md">
            <h1 class="text-2xl font-bold">{warehouse.Name}</h1>
            <p>{warehouse.Description}</p>
          </div>

          <div
            class="mt-4 flex flex-col"
            class:hidden={warehouse.rooms.length === 0}
          >
            <h2 class="text-xl font-bold mx-auto underline">Räume</h2>
            {#each warehouse.rooms as room}
              <div class="mt-4">
                <h2 class="text-xl font-bold">{room.Name}</h2>
                <a
                  href="/overview/rooms/{room.ID}"
                  class="flex hover:text-blue-500 duration-300"
                  >Siehe Raum <ArrowRight
                    class="my-auto ml-2 animate-pulse"
                  /></a
                >
              </div>
            {/each}
          </div>
        </div>
      </div>
    {/each}
  {/key}
</div>
