<script lang="ts">
  export let data;

  let shelves: any = data.shelves;
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">{shelves.Name}</div>
    <div class="text-gray-700 text-base mt-2">{shelves.Description}</div>
  </div>

  {#if shelves}
    <div
      class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
    >
      <table class="w-full text-sm text-left text-gray-700">
        <thead>
          <tr>
            <th class="font-semibold text-gray-600 uppercase px-6 py-4">
              Name
            </th>
            <th class="font-semibold text-gray-600 uppercase px-6 py-4">
              Items
            </th>
            <th class="font-semibold text-gray-600 uppercase px-6 py-4">
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          {#each shelves as shelf (shelf.ID)}
            <tr>
              <td class="px-6 py-4">{shelf.Name}</td>
              <td class="px-6 py-4"
                >{#if shelf.Items}
                  {shelf.Items.length}
                {:else}
                  0
                {/if}
              </td>
              <td class="px-6 py-4">
                <button
                  class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
                >
                  Edit
                </button>
                <button
                  class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
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
</div>
