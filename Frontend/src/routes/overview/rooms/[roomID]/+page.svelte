<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  export let data;

  let shelves: any = data.shelves;
  let room = data.room;
  let roomID = data.roomID;
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">{room.Name}</div>
  </div>

  {#if shelves}
    {#if shelves.length > 0}
      <div
        class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
      >
        <table class="w-full text-sm text-left text-gray-700">
          <thead class="text-xs uppercase bg-gray-200 text-gray-700">
            <tr>
              <th scope="col" class="px-6 py-3">Name</th>
              <th scope="col" class="text-right py-3 px-6">Items</th>
            </tr>
          </thead>
          <tbody>
            {#each shelves as shelf (shelf.ID)}
              <tr
                class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
                on:click={() => {
                  goto(`/overview/rooms/${roomID}/shelves/${shelf.ID}`);
                }}
              >
                <th
                  scope="row"
                  class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                  >{shelf.Name}</th
                >
                {#if shelf.Items}
                  <td class="text-right py-3 px-6">{shelf.Items.length}</td>
                {:else}
                  <td class="text-right py-3 px-6">0</td>
                {/if}
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {:else}
      <div class="text-lg">No shelves</div>
    {/if}
  {:else}
    <div class="text-lg">No shelves</div>
  {/if}
</div>
