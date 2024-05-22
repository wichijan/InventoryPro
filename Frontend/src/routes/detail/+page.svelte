<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  export let data;

  const shelvesItems: any | unknown[] = data.shelvesItems;

  let allItems: any[] = [];
  let itemsCopy = allItems;
  $: allItems = allItems;
  let latestThreeItems: any[] = [];
  $: latestThreeItems = latestThreeItems;

  onMount(async () => {
    for (const shelf of shelvesItems) {
      if (!shelf.Items) continue;
      for (const item of shelf.Items) {
        allItems = [
          ...allItems,
          {
            ID: item.ID,
            Name: item.Name,
            Description: item.Description,
            Quantity: item.Quantity,
            Regal: shelf.ShelveTypeName,
            Room: await getRoomName(shelf.RoomID),
            Damaged: item.Damaged,
          },
        ];
      }
    }
    itemsCopy = allItems;
    allItems.forEach((item, index) => {
      if (index < 3) {
        latestThreeItems = [...latestThreeItems, item];
      }
    });
  });

  async function getRoomName(roomID: number) {
    return new Promise((resolve, reject) => {
      fetch(API_URL + "rooms/" + roomID, {
        method: "GET",
        credentials: "include",
        mode: "cors",
        headers: {
          "Content-Type": "application/json",
        },
      }).then(async (response) => {
        if (response.ok) {
          await response.json().then((data) => {
            resolve(data.Name);
          });
        } else {
          reject(response.statusText);
        }
      });
    });
  }

  function search(event) {
    allItems = [...itemsCopy];
    const query = event.target.value.toLowerCase();
    allItems = allItems.filter((item) => {
      return item.Name.toLowerCase().includes(query);
    });
  }
</script>

<div class="flex min-h-screen items-center flex-col">
  <div class="mt-10 mb-4">
    <h1 class="text-3xl font-bold text-black">Items</h1>
  </div>
  <div class="grid grid-cols-3 bg-tertiary rounded">
    {#each latestThreeItems as item (item.ID)}
      <button
        class="max-w-sm rounded overflow-hidden shadow-lg m-3 bg-white px-5 py-5 hover:shadow-2xl duration-300 ease-in-out transform hover:scale-[1.02]"
        on:click={() => {
          goto(`/detail/${item.ID}`);
        }}
      >
        <img
          class="mx-auto rounded w-12 h-12 object-cover"
          src="https://via.placeholder.com/150"
          alt="Sunset in the mountains"
        />
        <div class="px-6 py-4">
          <div class="font-bold text-xl mb-2">{item.Name}</div>
          <p class="text-gray-700 text-base">{item.Description}</p>
        </div>
        <div class="pb-4 space-y-5">
          <span
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
          >
            Anzahl: {item.Quantity}
          </span>
          <span
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
          >
            Raum: {item.Room}
          </span>
          <span
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
          >
            Regal: {item.Regal}
          </span>
        </div>
      </button>
    {/each}
  </div>
  <div class="container mx-auto">
    <div class="my-2 flex sm:flex-row flex-col">
      <div class="sm:flex sm:flex-row-reverse mt-2">
        <div class="flex items-center my-2 sm:mb-0">
          <div class="relative">
            <input
              class="h-10 pl-2 pr-8 rounded-full border-2 border-gray-300 focus:ring-2 focus:ring-blue-500 duration-300 focus:outline-none"
              type="text"
              on:input|preventDefault={search}
              placeholder="Search by name..."
            />
          </div>
        </div>
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
              <th class="px-4 py-3">Raum</th>
              <th class="px-4 py-3">Regal</th>
              <th class="px-4 py-3">Kaputt?</th>
            </tr>
          </thead>
          <tbody class="bg-white">
            {#each allItems as item (item.ID)}
              <tr
                class="text-gray-700 hover:bg-tertiary duration-300 cursor-pointer"
                on:click={() => {
                  goto(`/detail/${item.ID}`);
                }}
              >
                <td class="px-4 py-3 border">{item.Name}</td>
                <td class="px-4 py-3 border">{item.Description}</td>
                <td class="px-4 py-3 border">{item.Quantity}</td>
                <td class="px-4 py-3 border">{item.Room}</td>
                <td class="px-4 py-3 border">{item.Regal}</td>
                <td class="px-4 py-3 border">{item.Damaged ? "Ja" : "Nein"}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
