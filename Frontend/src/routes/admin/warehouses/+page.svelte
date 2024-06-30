<script lang="ts">
  import Swal from "sweetalert2";
  import { API_URL } from "$lib/_services/ShelfService";
  import { BuildingFill } from "svelte-bootstrap-icons";
  import { goto } from "$app/navigation";

  export let data;

  let warehouse = data.allWarehouses;
  if (!warehouse) {
    warehouse = [];
  }
  $: warehouse = warehouse;
  $: {
    warehouse = warehouse.sort((a, b) => {
      return a.Name.localeCompare(b.Name);
    });
  }
</script>

<div class="flex flex-col items-center w-full">
  <div class="mt-10 mb-4">
    <h1 class="text-3xl font-bold text-black">Warehouses</h1>
  </div>
  {#if !warehouse}
    <div class="text-lg">Es sind keine Warehouses vorhanden</div>
  {/if}
  {#if warehouse}
    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 bg-tertiary rounded px-2 py-4"
    >
      {#each warehouse as warehouse (warehouse.ID)}
        <button
          class="max-w-sm rounded overflow-hidden shadow-lg m-3 hover:bg-white px-5 py-5 hover:shadow-xl duration-300 ease-in-out transform hover:scale-[102%]
            bg-[#a3b18a]"
          on:click={() => {
            goto(`/admin/warehouses/${warehouse.ID}`);
          }}
        >
          <BuildingFill class="mx-auto rounded w-8 h-8 object-cover" />
          <div class="px-6 py-4">
            <div class="font-bold text-xl mb-2">{warehouse.Name}</div>
            <p class="text-gray-700 text-base">{warehouse.Description}</p>
          </div>
        </button>
      {/each}
      <button
        class="max-w-sm rounded overflow-hidden shadow-lg m-3 bg-white hover:bg-green-500 hover:text-white px-5 py-5 hover:shadow-xl duration-300 ease-in-out transform hover:scale-[102%]"
        on:click={() => {
          Swal.fire({
            title: "Erstelle ein neues warehouse",
            html: `
              <input id="name" class="swal2-input" placeholder="Name">
              <input id="description" class="swal2-input" placeholder="Beschreibung">
            `,
            showCancelButton: true,
            confirmButtonText: `Create`,
          }).then((result) => {
            if (result.isConfirmed) {
              const name = document.getElementById("name").value;
              const description = document.getElementById("description").value;
              if (warehouse.some((w) => w.Name === name)) {
                Swal.fire("Warehouse existiert bereits", "", "error");
                return;
              }
              fetch(`${API_URL}warehouses`, {
                method: "POST",
                credentials: "include",
                headers: {
                  "Content-Type": "application/json",
                },
                body: JSON.stringify({
                  Name: name,
                  Description: description,
                }),
              }).then(async (response) => {
                if (response.ok) {
                  const warehouseID = await response.json();
                  Swal.fire("Warehouse created!", "", "success");
                  warehouse.push({
                    ID: warehouseID,
                    Name: name,
                    Description: description,
                    Rooms: [],
                  });
                  warehouse = warehouse;
                } else {
                  Swal.fire("Error creating warehouse", "", "error");
                }
              });
            }
          });
        }}
      >
        <BuildingFill class="mx-auto rounded w-8 h-8 object-cover" />
        <div class="px-6 py-4">
          <div class="font-bold text-xl mb-2">Erstelle ein neues warehouse</div>
          <p class="text-gray-700 text-base">
            Hier kannst du ein neues warehouse erstellen
          </p>
        </div>
      </button>
    </div>
  {/if}
</div>
