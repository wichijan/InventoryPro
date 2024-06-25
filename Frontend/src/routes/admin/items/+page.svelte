<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import { SortDown, SortUp, Filter } from "svelte-bootstrap-icons";

  export let data;

  let items = data.items;
  const defaultItems = JSON.parse(JSON.stringify(items));

  let showItems = items;
  $: showItems = showItems;

  let cutOffDescription = 40;

  async function getMoreInformation(item: any) {
    let url = item.ItemTypes === "book" ? "book" : "set-of-objects";
    const response = await fetch(`${API_URL}items/${url}/${item.ID}`);
    const data = await response.json();
    return data;
  }

  let sort = 0;
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6 mb-10">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">Items</div>
  </div>

  {#if items}
    {#if items.length === 0}
      <div class="text-center text-gray-700">No items found</div>
    {:else}
      <div
        class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
      >
        <!-- filter icon and sort icon should be here-->
        <div class="flex justify-between items-center p-4">
          <div class="flex items-center">
            <input
              type="text"
              class="border border-gray-300 p-2 rounded-lg"
              placeholder="Search..."
              on:input={(e) => {
                if (
                  e.target.value === "" ||
                  e.target.value === undefined ||
                  e.target.value === null
                ) {
                  showItems = defaultItems;
                } else {
                  //check if input string is in any of the fields
                  showItems = showItems.filter((item) => {
                    let damaged = item.Damaged ? "Kaputt" : "Unversehen";
                    return (
                      item.Name.toLowerCase().includes(
                        e.target.value.toLowerCase()
                      ) ||
                      item.Description.toLowerCase().includes(
                        e.target.value.toLowerCase()
                      ) ||
                      item.QuantityInShelf.toString().includes(
                        e.target.value.toLowerCase()
                      ) ||
                      damaged
                        .toLowerCase()
                        .includes(e.target.value.toLowerCase())
                    );
                  });
                }
              }}
            />
            <!-- <button class="ml-2">
              <Filter class="h-6 w-6 text-gray-500" />
            </button> -->
          </div>
          <div class="flex items-center">
            <select
              class="border border-gray-300 p-2 rounded-lg"
              name="sort"
              id="sort"
              on:change={(e) => {
                items = items.sort((a, b) => {
                  if (e.target.value === "name") {
                    if (sort === 0) {
                      return a.Name < b.Name ? -1 : 1;
                    } else {
                      return a.Name > b.Name ? -1 : 1;
                    }
                  } else if (e.target.value === "quantity") {
                    if (sort === 0) {
                      return a.QuantityInShelf < b.QuantityInShelf ? -1 : 1;
                    } else {
                      return a.QuantityInShelf > b.QuantityInShelf ? -1 : 1;
                    }
                  } else {
                    if (sort === 0) {
                      return a.Damaged < b.Damaged ? -1 : 1;
                    } else {
                      return a.Damaged > b.Damaged ? -1 : 1;
                    }
                  }
                });
              }}
            >
              <option value="name">Name</option>
              <option value="quantity">Quantity</option>
              <option value="status">Status</option>
            </select>
            <button
              class="ml-2"
              on:click={() => {
                sort = sort === 0 ? 1 : 0;
                let e = document.getElementById("sort");
                //sort the items in the right way
                items = items.sort((a, b) => {
                  if (e.value === "name") {
                    if (sort === 0) {
                      return a.Name < b.Name ? -1 : 1;
                    } else {
                      return a.Name > b.Name ? -1 : 1;
                    }
                  } else if (e.value === "quantity") {
                    if (sort === 0) {
                      return a.QuantityInShelf < b.QuantityInShelf ? -1 : 1;
                    } else {
                      return a.QuantityInShelf > b.QuantityInShelf ? -1 : 1;
                    }
                  } else {
                    if (sort === 0) {
                      return a.Damaged < b.Damaged ? -1 : 1;
                    } else {
                      return a.Damaged > b.Damaged ? -1 : 1;
                    }
                  }
                });
              }}
            >
              {#if sort === 0}
                <SortUp class="h-6 w-6 text-gray-500" />
              {:else}
                <SortDown class="h-6 w-6 text-gray-500" />
              {/if}
            </button>
          </div>
        </div>
        <table class="w-full text-sm text-left text-gray-700">
          <thead class="text-xs uppercase bg-gray-200 text-gray-700">
            <tr>
              <th scope="col" class="px-6 py-3">Name</th>
              <th scope="col" class="px-6 py-3">Description</th>
              <th scope="col" class="px-6 py-3">QuantityInShelf</th>
              <th scope="col" class="px-6 py-3">Status</th>
              <th scope="col" class="px-6 py-3 text-right">Action</th>
            </tr>
          </thead>
          <tbody>
            {#each showItems as item (item.ID)}
              <tr
                class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
                on:click={() => {
                  getMoreInformation(item).then((data) => {
                    console.log(data);
                  });
                }}
              >
                <td class="px-6 py-4">{item.Name}</td>
                <td class="px-6 py-4"
                  >{item.Description.length > cutOffDescription
                    ? item.Description.substring(0, cutOffDescription) + "..."
                    : item.Description}</td
                >
                <td class="px-6 py-4">{item.QuantityInShelf}</td>
                <td class="px-6 py-4"
                  >{item.Damaged ? "Kaputt" : "Unversehen"}</td
                >
                <td class="px-6 py-4 text-right">
                  <button
                    class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                  >
                    Edit
                  </button>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  {/if}
</div>
