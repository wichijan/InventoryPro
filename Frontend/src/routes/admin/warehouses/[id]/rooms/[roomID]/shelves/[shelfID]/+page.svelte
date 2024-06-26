<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";
  export let data;

  let shelf: any = data.shelf;
  const items = shelf.Items || [];
  let warehouseID = data.warehouseID;
  let roomID = data.roomID;
  let shelfID = data.shelfID;
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">{shelf.Name}</div>
  </div>

  {#if items.length > 0}
    <div
      class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
    >
      <table class="w-full text-sm text-left text-gray-700">
        <thead class="text-xs uppercase bg-gray-200 text-gray-700">
          <tr>
            <th scope="col" class="px-6 py-3">Name</th>
            <th scope="col" class="text-center">Anzahl</th>
            <th scope="col" class="px-6 py-3 text-right">Action</th>
          </tr>
        </thead>
        <tbody>
          {#each items as item (item.ID)}
            <tr
              class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
              on:click={() => {
                goto(
                  `/admin/warehouses/${warehouseID}/rooms/${roomID}/shelves/${shelfID}/items/${item.ID}`
                );
              }}
            >
              <th
                scope="row"
                class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                >{item.Name}</th
              >
              <td class="text-center">{item.Quantity}</td>
              <td class="px-6 py-4 text-right">
                <button
                  on:click|stopPropagation={() =>
                    Swal.fire({
                      title: "Item löschen",
                      text: "Möchten Sie dieses Item wirklich löschen?",
                      showCancelButton: true,
                      confirmButtonText: `Delete`,
                    }).then((result) => {
                      if (result.isConfirmed) {
                        fetch(`${API_URL}items/${item.ID}`, {
                          method: "DELETE",
                          credentials: "include",
                        }).then((response) => {
                          if (response.ok) {
                            Swal.fire("Deleted!", "", "success");
                            location.reload();
                          } else {
                            Swal.fire("Error!", "", "error");
                          }
                        });
                      }
                    })}
                >
                  Löschen
                </button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {:else}
    <div class="text-center text-gray-700 mt-4">
      Es sind keine Items vorhanden
    </div>
  {/if}
  <button
    class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
    on:click={() => {
      goto(
        `/admin/warehouses/${warehouseID}/rooms/${roomID}/shelves/${shelfID}/items/create`
      );
    }}
  >
    Item hinzufügen
  </button>
</div>
