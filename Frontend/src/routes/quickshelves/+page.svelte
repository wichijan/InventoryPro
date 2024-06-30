<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  export let data;

  let quickshelves = data.quickshelves;
  let warehouses = data.warehouses;

  function getWarehouseAndRoomName(quickshelf) {
    for (let warehouse of warehouses) {
      if (warehouse.Rooms) {
        for (let room of warehouse.Rooms) {
          if (room.ID === quickshelf.RoomID) {
            return {
              warehouseName: warehouse.Name,
              roomName: room.Name,
            };
          }
        }
      }
    }
    return null;
  }
</script>

{#if !quickshelves}
  <div class="container mx-auto py-8">
    <div class="flex flex-row flex-wrap">
      <div class="w-full p-4">
        <div class="bg-white rounded-lg shadow-lg">
          <div class="p-4">
            <h3 class="font-semibold text-xl">Kein Schnellregal</h3>
            <p>Es gibt keine Schnellregale.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
{:else}
  <div class="container mx-auto py-8">
    <div class="flex flex-row flex-wrap w-full">
      {#each quickshelves as quickshelf}
        <div class="w-full p-4">
          <button
            class="bg-white rounded-lg shadow-lg w-full text-left"
            on:click={() => {
              goto(`/quickshelves/${quickshelf.QuickShelfID}`);
            }}
          >
            <div class="p-4 space-y-2">
              <h3 class="font-semibold text-xl">
                Name: {quickshelf.Name || "Unbekannt"}
              </h3>
              <h3 class="font-semibold text-xl">
                Warehouse: {getWarehouseAndRoomName(quickshelf)
                  ?.warehouseName || "Unbekannt"}
              </h3>
              <h2 class="font-semibold text-xl">
                Raum: {getWarehouseAndRoomName(quickshelf)?.roomName ||
                  "Unbekannt"}
              </h2>
              <p>Items: {quickshelf.Items ? quickshelf.Items.length : 0}</p>
            </div>
          </button>
        </div>
      {/each}
    </div>
  </div>
{/if}
