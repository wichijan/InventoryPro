<script lang="ts">
  import { writable } from "svelte/store";

  let editable = writable(false);
  export let data;

  let item: any | unknown[] = data.item;

  const toggleEditable = () => {
    editable.update((value) => !value);
  };
</script>

<div
  class="m-4 p-4 bg-gray-100 text-gray-800 transform shadow-lg rounded-xl transition ease-in-out duration-500 ml-16 mr-10"
>
  <h1 class="text-5xl font-extrabold mb-2 text-gray-700">
    Informationen über {item.Name}
  </h1>
  <div class="flex justify-between items-center mb-4">
    <button
      class="px-4 py-2 bg-indigo-500 text-white rounded-2xl
                  transform active:scale-90 focus:outline-none transition hover:bg-indigo-600 ease-in-out duration-500 shadow-md hover:shadow-lg"
      on:click|preventDefault={toggleEditable}
    >
      {#if $editable}
        Speichern
      {:else}
        Bearbeiten
      {/if}
    </button>
  </div>
  <div>
    {#each Object.keys(item) as key (key)}
      <div class="my-2 flex items-center space-x-4">
        {#if key !== "ID"}
          <label for={key} class="w-1/3 text-gray-600 capitalize">{key}</label>
          {#if $editable && key.includes("Class")}
            <select
              class="py-2 px-4 w-full border border-gray-300 rounded-md
                        focus:outline-none focus:ring-2 focus:ring-indigo-500 transition
                        bg-gray-200 text-gray-700 shadow-sm hover:shadow-md"
              bind:value={item[key]}
              id={key}
            >
              <option value="true">True</option>
              <option value="false">False</option>
            </select>
          {:else if $editable}
            <input
              class="py-2 px-4 w-full border border-gray-300 rounded-md
                        focus:outline-none focus:ring-2 focus:ring-indigo-500 transition
                        bg-gray-200 text-gray-700 shadow-sm hover:shadow-md"
              type="text"
              bind:value={item[key]}
              id={key}
            />
          {:else}
            <div
              class="py-2 px-4 w-full border border-gray-300 rounded-md shadow-sm"
            >
              {item[key]}
            </div>
          {/if}
        {/if}
      </div>
    {/each}
  </div>
</div>
