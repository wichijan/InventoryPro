<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import { onMount } from "svelte";
  import Swal from "sweetalert2";

  export let data;

  let warehouse = data.warehouse;

  let rooms: any = [];
  $: rooms = rooms;
  $: {
    rooms = rooms.sort((a, b) => {
      return a.Name.localeCompare(b.Name);
    });
  }

  onMount(async () => {
    warehouse.Rooms.forEach((room: any) => {
      getRoomDetails(room.ID).then((data) => {
        rooms.push(data);
        rooms = rooms;
      });
    });
  });

  async function getRoomDetails(id: string) {
    return new Promise((resolve, reject) => {
      fetch(API_URL + "rooms-with-shelves/" + id, {
        method: "GET",
        credentials: "include",
        mode: "cors",
        headers: {
          "Content-Type": "application/json",
        },
      }).then(async (response) => {
        if (response.ok) {
          await response.json().then((data) => {
            resolve(data);
          });
        } else {
          reject(response.statusText);
        }
      });
    });
  }
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">{warehouse.Name}</div>
    <div class="text-gray-700 text-base mt-2">{warehouse.Description}</div>
  </div>

  {#if warehouse.Rooms}
    <div
      class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
    >
      <table class="w-full text-sm text-left text-gray-700">
        <thead class="text-xs uppercase bg-gray-200 text-gray-700">
          <tr>
            <th scope="col" class="px-6 py-3">Name</th>
            <th scope="col" class="px-6 py-3">Shelves</th>
            <th scope="col" class="px-6 py-3">Items</th>
            <th scope="col" class="px-6 py-3"
              ><span class="sr-only">Action</span></th
            >
          </tr>
        </thead>
        <tbody>
          {#each rooms as room (room.ID)}
            <tr
              class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
              on:click|stopPropagation={() =>
                goto(`/admin/warehouses/${warehouse.ID}/rooms/${room.ID}`)}
            >
              <th
                scope="row"
                class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                >{room.Name}</th
              >
              {#if room.Shelves}
                <td class="px-6 py-4">{room.Shelves.length}</td>
              {:else}
                <td class="px-6 py-4">0</td>
              {/if}
              <td class="px-6 py-4">0</td>
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
                        fetch(API_URL + "rooms/" + room.ID, {
                          method: "DELETE",
                          credentials: "include",
                          mode: "cors",
                          headers: {
                            "Content-Type": "application/json",
                          },
                        }).then(async (response) => {
                          if (response.ok) {
                            Swal.fire("Room deleted!", "", "success");
                            rooms = rooms.filter((r) => r.ID !== room.ID);
                          } else {
                            Swal.fire("Error deleting room", "", "error");
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
    <div class="text-lg mt-10">No rooms</div>
  {/if}
  <button
    class="mt-5 bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition-colors"
    on:click={() =>
      Swal.fire({
        title: "Create new room",
        html: `
                    <input id="name" class="swal2-input" placeholder="Name">
                `,
        showCancelButton: true,
        confirmButtonText: `Create`,
      }).then((result) => {
        if (result.isConfirmed) {
          const name = document.getElementById("name").value;
          if (rooms.some((r) => r.Name === name)) {
            Swal.fire("Room already exists", "", "error");
            return;
          }
          fetch(API_URL + "rooms", {
            method: "POST",
            credentials: "include",
            mode: "cors",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              Name: name,
              WarehouseID: warehouse.ID,
            }),
          }).then(async (response) => {
            if (response.ok) {
              Swal.fire("Room created!", "", "success");
              await response.json().then(async (room) => {
                await getRoomDetails(room).then((data) => {
                  rooms.push(data);
                  rooms = rooms;
                });
              });
            } else {
              Swal.fire("Error creating room", "", "error");
            }
          });
        }
      })}
  >
    Create new room
  </button>
</div>
