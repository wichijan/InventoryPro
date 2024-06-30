<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import { onMount } from "svelte";
  import {
    SortDown,
    SortUp,
    XCircle,
    CheckCircle,
  } from "svelte-bootstrap-icons";
  import Swal from "sweetalert2";

  export let data;

  const quickShelf = data.quickShelf;
  let items = quickShelf.Items;
  const defaultItems = JSON.parse(JSON.stringify(items));

  let showItems = items;
  $: showItems = showItems;

  let cutOffDescription = 40;

  function buildClassString(item: any) {
    let classes = "";
    if (item.ClassOne) {
      classes += "1";
    }
    if (item.ClassTwo) {
      classes += classes.length > 0 ? ", 2" : "2";
    }
    if (item.ClassThree) {
      classes += classes.length > 0 ? ", 3" : "3";
    }
    if (item.ClassFour) {
      classes += classes.length > 0 ? ", 4" : "4";
    }
    return classes;
  }

  let sort = 0;
</script>

<div
  class=" min-h-screen text-gray-900 flex flex-col items-center p-7 mb-10 w-full h-full"
>
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">Items</div>
  </div>

  {#if items}
    {#if items.length === 0}
      <div class="text-center text-gray-700">Keine Items gefunden</div>
    {:else}
      <div
        class="relative overflow-x-auto shadow-lg rounded-lg w-full mt-5 ml-1 bg-white"
      >
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
                        .includes(e.target.value.toLowerCase()) ||
                      buildClassString(item)
                        .toLowerCase()
                        .includes(e.target.value.toLowerCase())
                    );
                  });
                }
              }}
            />
          </div>
          <div class="flex items-center">
            <select
              class="border border-gray-300 p-2 rounded-lg"
              name="sort"
              id="sort"
              on:change={(e) => {
                showItems = showItems.sort((a, b) => {
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
              <option value="quantity">Anzahl</option>
              <option value="status">Status</option>
            </select>
            <button
              class="ml-2"
              on:click={() => {
                sort = sort === 0 ? 1 : 0;
                let e = document.getElementById("sort");
                showItems = showItems.sort((a, b) => {
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
              <th scope="col" class="px-6 py-3">Beschreibung</th>
              <th scope="col" class="px-6 py-3">Anzahl im Regal</th>
              <th scope="col" class="px-6 py-3">Status</th>
              <th scope="col" class="px-6 py-3">Verfügbar</th>
              <th scope="col" class="px-6 py-3 text-right">Klassen</th>
              <th scope="col" class="px-6 py-3 text-right">Action</th>
            </tr>
          </thead>
          <tbody>
            {#each showItems as item (item.ID)}
              <tr
                class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
                on:click={() => {
                  goto(`/items/${item.ID}`);
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
                <td class="px-6 py-4 text-center"
                  >{#if item.UsersBorrowed && item.QuantityInShelf === 0}
                    <XCircle class="h-6 w-6 text-red-500 mx-auto" />
                  {:else}
                    <CheckCircle class="h-6 w-6 text-green-500 mx-auto" />
                  {/if}
                </td>
                <td class="px-6 py-4 text-right">
                  {buildClassString(item)}
                </td>
                <td class="px-6 py-4 text-right">
                  <button
                    class="text-red-500"
                    on:click|stopPropagation={(e) => {
                      e.stopPropagation();
                      fetch(`${API_URL}items/remove-item-from-quickshelf`, {
                        credentials: "include",
                        method: "DELETE",
                        headers: {
                          "Content-Type": "application/json",
                        },
                        body: JSON.stringify({
                          QuickShelfID: quickShelf.QuickShelfID,
                          ItemID: item.ID,
                        }),
                      }).then((res) => {
                        if (res.status !== 200) {
                          Swal.fire({
                            title: "Fehler",
                            text: "Item konnte nicht zurück konvertiert werden",
                            icon: "error",
                          });
                        } else {
                          Swal.fire({
                            title: "Erfolg",
                            text: "Item wurde zurück konvertiert",
                            icon: "success",
                          });
                          items = items.filter((i) => i.ID !== item.ID);
                          showItems = showItems.filter((i) => i.ID !== item.ID);
                        }
                      });
                    }}
                  >
                    Löschen
                  </button>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  {:else}
    <div class="text-center text-gray-700">Keine Items gefunden</div>
  {/if}
</div>
