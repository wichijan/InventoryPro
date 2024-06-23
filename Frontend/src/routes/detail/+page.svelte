<script lang="ts">
  import Spinner from "$lib/templates/Spinner.svelte";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";

  export let data;

  const shelvesItems = data; // Assuming the entire dataset is passed as `data`
  console.log(shelvesItems);

  let warehouseName: string = "";

  let allItems: any[] = [];
  let itemsCopy = allItems;
  $: allItems = allItems;
  let latestThreeItems: any[] = [];
  $: latestThreeItems = latestThreeItems;

  onMount(() => {
    allItems = shelvesItems.shelvesItems;
    itemsCopy = allItems;
    latestThreeItems = allItems.slice(0, 3); // Get the latest three items
    if (browser) {
      warehouseName = localStorage.getItem("warehouse") || "";
    }
  });

  function search(event) {
    const query = event.target.value.toLowerCase();
    allItems = itemsCopy.filter((item) => {
      return item.Name.toLowerCase().includes(query);
    });
  }
</script>

<div class="w-full bg-tertiary rounded ml-10 mt-5 px-2 py-2 flex">
  <div class="text-lg">
    Zeigt items für das Warehouse: <b>{warehouseName}</b>
  </div>
  <button
    class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-white hover:shadow-lg duration-500 text-black rounded-md ml-3 px-1"
    on:click={() => {
      goto("/settings");
    }}
  >
    Warehouse ändern
  </button>
</div>

{#if allItems.length > 0}
  <div class="flex flex-col items-center w-full">
    <div class="mt-10 mb-4">
      <h1 class="text-3xl font-bold text-black">Items</h1>
    </div>
    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 bg-tertiary rounded px-2 py-4"
    >
      {#each latestThreeItems as item (item.ID)}
        <button
          class="max-w-sm rounded overflow-hidden shadow-lg m-3 bg-white px-5 py-5 hover:shadow-xl duration-300 ease-in-out transform hover:scale-[102%]"
          on:click={() => {
            goto(`/detail/${item.ID}`);
          }}
        >
          <img
            class="mx-auto rounded w-12 h-12 object-cover"
            src="https://via.placeholder.com/150"
            alt="Image"
          />
          <div class="px-6 py-4">
            <div class="font-bold text-xl mb-2">{item.Name}</div>
            <p class="text-gray-700 text-base">{item.Description}</p>
          </div>
          <div class="pb-4 space-y-5">
            <span
              class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
            >
              Anzahl: {item.QuantityInShelf}
            </span>
            <span
              class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
            >
              Class One: {item.ClassOne ? "Yes" : "No"}
            </span>
            <span
              class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
            >
              Class Two: {item.ClassTwo ? "Yes" : "No"}
            </span>
          </div>
        </button>
      {/each}
    </div>
    <div class="container mx-auto px-4">
      <div class="my-2 flex flex-col sm:flex-row justify-between items-center">
        <div class="relative w-full sm:w-auto mb-2 sm:mb-0">
          <input
            class="w-full sm:w-64 h-10 pl-2 pr-8 rounded-full border-2 border-gray-300 focus:ring-2 focus:ring-blue-500 duration-300 focus:outline-none"
            type="text"
            on:input|preventDefault={search}
            placeholder="Search by name..."
          />
        </div>
      </div>
      <div class="w-full mb-8 overflow-hidden rounded-lg shadow-lg">
        <div class="w-full overflow-x-auto">
          <table class="w-full">
            <thead>
              <tr
                class="text-md font-bold tracking-wide text-left text-gray-900 bg-gray-100 capitalize border-b border-gray-600"
              >
                <th class="px-4 py-3">Name</th>
                <th class="px-4 py-3">Beschreibung</th>
                <th class="px-4 py-3">Anzahl</th>
                <th class="px-4 py-3">Reservierungen</th>
                <th class="px-4 py-3">Verfügbar?</th>
                <th class="px-4 py-3">Kaputt?</th>
              </tr>
            </thead>
            <tbody class="bg-white">
              {#each allItems as item (item.ID)}
                <tr
                  class="text-gray-700 hover:bg-tertiary duration-300 cursor-pointer"
                  on:click={() => goto(`/detail/${item.ID}`)}
                >
                  <td class="px-4 py-3 border">{item.Name}</td>
                  <td class="px-4 py-3 border">{item.Description}</td>
                  <td class="px-4 py-3 border">{item.QuantityInShelf}</td>
                  <td class="px-4 py-3 border"
                    >{item.Reservations ? item.Reservations.length : 0}</td
                  >
                  <td class="px-4 py-3 border"
                    >{item.UsersBorrowed ? "Nein" : "Ja"}</td
                  >
                  <td class="px-4 py-3 border"
                    >{item.Damaged ? "Ja" : "Nein"}</td
                  >
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
{:else}
  <Spinner />
{/if}
