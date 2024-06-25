<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  export let data;

  let shelves: any = data.shelves;
  let room = data.room;
  let warehouseID = data.warehouseID;
  let roomID = data.roomID;
  console.log(shelves);
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">{room.Name}</div>
  </div>

  {#if shelves}
    <div
      class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
    >
      <table class="w-full text-sm text-left text-gray-700">
        <thead class="text-xs uppercase bg-gray-200 text-gray-700">
          <tr>
            <th scope="col" class="px-6 py-3">Name</th>
            <th scope="col" class="text-center">Items</th>
            <th scope="col" class="px-6 py-3 text-right">Action</th>
          </tr>
        </thead>
        <tbody>
          {#each shelves as shelf (shelf.ID)}
            <tr
              class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
              on:click={() => {
                goto(
                  `/admin/warehouses/${warehouseID}/rooms/${roomID}/shelves/${shelf.ID}`
                );
              }}
            >
              <th
                scope="row"
                class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                >{shelf.Name}</th
              >
              {#if shelf.Items}
                <td class="text-center">{shelf.Items.length}</td>
              {:else}
                <td class="text-center">0</td>
              {/if}
              <td class="px-6 py-4 text-right">
                <button
                  on:click|stopPropagation={() =>
                    Swal.fire({
                      title: "Delete room",
                      text: "Are you sure you want to delete this room?",
                      showCancelButton: true,
                      confirmButtonText: `Delete`,
                    }).then((result) => {
                      if (result.isConfirmed) {
                        fetch(API_URL + "shelves/" + shelf.ID, {
                          method: "DELETE",
                          credentials: "include",
                          mode: "cors",
                          headers: {
                            "Content-Type": "application/json",
                          },
                        }).then(async (response) => {
                          if (response.ok) {
                            Swal.fire("Shelves deleted!", "", "success");
                            shelves = shelves.filter((s) => s.ID !== shelf.ID);
                          } else {
                            Swal.fire("Error deleting shelf", "", "error");
                          }
                        });
                      }
                    })}
                  class="text-blue-500 hover:text-blue-600 font-semibold transition-colors"
                >
                  Delete
                </button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {:else}
    <div class="text-lg">No shelves</div>
  {/if}
  <!-- Create new shelf-->
  <button
    class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
    on:click={() => {
      Swal.fire({
        title: "Create Shelf",
        html: `
          <input id="name" class="swal2-input" placeholder="Name" />
        `,
        showCancelButton: true,
        confirmButtonText: "Create",
      }).then((result) => {
        if (result.isConfirmed) {
          const name = document.getElementById("name").value;
          fetch(API_URL + "shelves", {
            method: "POST",
            credentials: "include",
            mode: "cors",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              Name: name,
              RoomID: roomID,
            }),
          }).then(async (response) => {
            if (response.ok) {
              Swal.fire("Shelf created!", "", "success").then(() => {
                location.reload();
              });
            } else {
              Swal.fire("Error creating shelf", "", "error");
            }
          });
        }
      });
    }}
  >
    Create Shelf
  </button>
</div>
