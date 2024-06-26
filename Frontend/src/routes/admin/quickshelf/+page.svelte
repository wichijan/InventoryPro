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

  async function deleteQuickShelf(quickShelfID) {
    Swal.fire({
      title: "Bist du sicher?",
      text: "Diese Aktion ist irreversibel!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Ja, lösche es!",
    }).then((result) => {
      if (result.isConfirmed) {
        fetch(`${API_URL}quick-shelves/${quickShelfID}`, {
          method: "DELETE",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
        }).then((response) => {
          if (response.ok) {
            Swal.fire(
              "Deleted!",
              "Dein Schnellregal wurde gelöscht.",
              "success"
            );
          } else {
            Swal.fire("Error", "Failed to delete quickshelf", "error");
          }
        });
      }
    });
  }

  async function clearQuickShelf(quickshelf) {
    if (!quickshelf.Items || quickshelf.Items.length === 0) {
      Swal.fire("Error", "Schnellregal ist bereits leer", "error");
      return;
    }
    Swal.fire({
      title: "Bist du sicher?",
      text: "Diese Aktion ist irreversibel!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Ja, leere es!",
    }).then((result) => {
      if (result.isConfirmed) {
        fetch(`${API_URL}clear-quick-shelves/${quickshelf.QuickShelfID}`, {
          method: "DELETE",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
        }).then((response) => {
          if (response.ok) {
            Swal.fire(
              "Cleared!",
              "Dein Schnellregal wurde geleert.",
              "success"
            );
          } else {
            Swal.fire("Error", "Failed to clear quickshelf", "error");
          }
        });
      }
    });
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
          <div class="bg-white rounded-lg shadow-lg w-full text-left">
            <div class="p-4 space-y-2">
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
            <div class="p-4">
              <button
                class="bg-red-500 text-white rounded-lg py-2 px-2 shadow-lg hover:bg-red-400 duration-300"
                on:click={() => {
                  clearQuickShelf(quickshelf);
                }}
              >
                Leeren
              </button>
              <button
                class="bg-red-500 text-white rounded-lg py-2 px-2 shadow-lg hover:bg-red-400 duration-300"
                on:click={() => {
                  deleteQuickShelf(quickshelf.QuickShelfID);
                }}
              >
                Löschen
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
{/if}
<div class="container mx-auto py-8">
  <div class="flex flex-row flex-wrap">
    <div class="w-full p-4">
      <button
        class="bg-white rounded-lg shadow-lg w-full hover:bg-green-500 duration-300"
        on:click={() => {
          goto("/admin/quickshelf/create");
        }}
      >
        <div class="p-4">
          <h3 class="font-semibold text-xl">Erstellen</h3>
          <p>Erstelle ein neues Schnellregal.</p>
        </div>
      </button>
    </div>
  </div>
</div>
