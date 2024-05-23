<script lang="ts">
  import { browser } from "$app/environment";
  import UserSettings from "$lib/templates/UserSettings.svelte";
  import { onMount } from "svelte";
  import Swal from "sweetalert2";

  export let data;

  type Warehouse = {
    Description: string;
    Id: string;
    Name: string;
  };

  const allWarehouses: Warehouse[] = data.warehouses;

  let warehouse: string = "Keins ausgewählt";
  let oldWarehouse: string = "Keins ausgewählt";

  onMount(() => {
    warehouse = browser
      ? localStorage.getItem("warehouse") === null
        ? "Keins ausgewählt"
        : localStorage.getItem("warehouse")?.toString() || ""
      : "Keins ausgewählt";
    oldWarehouse = warehouse;
  });

  function selectWarehouse() {
    if (oldWarehouse === warehouse) {
      return;
    }
    browser && localStorage.setItem("warehouse", warehouse);
    Swal.fire({
      position: "top-end",
      icon: "success",
      title: "Warehouse wurde erfolgreich geändert!",
      showConfirmButton: false,
      timer: 1500,
    });
    oldWarehouse = warehouse;
  }
</script>

<div class="my-5">
  <div
    class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 rounded-md"
    id="userSettings"
  >
    <div class="mx-auto font-semibold text-2xl" id="header">User Settings</div>
    <div class="flex flex-row justify-between mt-3">
      <div class="flex flex-col">
        <div class="font-semibold text-lg">Warehouse</div>
        <select
          class="mt-2 w-60 h-10 border border-gray-300 rounded-md"
          bind:value={warehouse}
        >
          {#each allWarehouses as warehouse}
            <option value={warehouse.Name}>{warehouse.Name}</option>
          {/each}
        </select>
        {#key warehouse}
          <button
            class="bg-[#d5bdaf] hover:bg-d6ccc2 enabled:hover:text-black enabled:hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full
          disabled:opacity-50 disabled:cursor-not-allowed"
            class:disabled={oldWarehouse === warehouse}
            on:click={() => selectWarehouse()}
          >
            Speichern
          </button>
        {/key}
        {#if oldWarehouse !== warehouse}
          <div class="text-sm mt-2 text-red-500 mx-auto">
            Änderungen speichern!
          </div>
        {/if}
      </div>
      <div class="flex flex-col">
        <UserSettings />
      </div>
    </div>
  </div>
</div>
