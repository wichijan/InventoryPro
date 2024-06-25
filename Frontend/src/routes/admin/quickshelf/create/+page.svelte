<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";
  export let data;

  let warehouses = data.warehouses;
  let selectedWarehouse = null;
  let selectedRoom = null;
  let quickshelfName = "";

  //sort warehoues by room count
  $: warehouses = warehouses.sort((a, b) => {
    let aRooms = a.Rooms ? a.Rooms.length : 0;
    let bRooms = b.Rooms ? b.Rooms.length : 0;
    return bRooms - aRooms;
  });

  function handleWarehouseChange(event) {
    const warehouseId = event.target.value;
    selectedWarehouse = warehouses.find(
      (warehouse) => warehouse.ID === warehouseId
    );
    selectedRoom = selectedWarehouse?.Rooms ? selectedWarehouse.Rooms[0] : null;
  }

  function handleRoomChange(event) {
    const roomId = event.target.value;
    selectedRoom = selectedWarehouse.Rooms.find((room) => room.ID === roomId);
  }

  function createQuickshelf() {
    if (selectedWarehouse && selectedRoom && quickshelfName) {
      fetch(`${API_URL}quick-shelves`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          RoomId: selectedRoom.ID,
        }),
      }).then((response) => {
        if (response.ok) {
          Swal.fire({
            icon: "success",
            title: "Quickshelf created successfully!",
          });
        } else {
          Swal.fire({
            icon: "error",
            title: "Error",
            text: "Failed to create Quickshelf",
          });
        }
      });
    } else {
      Swal.fire({
        icon: "error",
        title: "Error",
        text: "Please fill in all the fields",
      });
    }
  }
</script>

<main class="p-6 bg-gray-100 min-h-screen">
  <h1 class="text-2xl font-bold mb-6 text-center">Create a New Quickshelf</h1>

  <div class="mb-4">
    <label for="warehouse" class="block text-sm font-medium text-gray-700"
      >Select Warehouse:</label
    >
    <select
      id="warehouse"
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
      on:change={handleWarehouseChange}
    >
      <option value="" disabled selected>Select a warehouse</option>
      {#each warehouses as warehouse}
        <option value={warehouse.ID}
          >{warehouse.Name} ({warehouse.Rooms
            ? warehouse.Rooms.length
            : 0})</option
        >
      {/each}
    </select>
  </div>

  {#if selectedWarehouse && selectedWarehouse.Rooms}
    <div class="mb-4">
      <label for="room" class="block text-sm font-medium text-gray-700"
        >Select Room:</label
      >
      <select
        id="room"
        required
        class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        on:change={handleRoomChange}
      >
        {#each selectedWarehouse.Rooms as room}
          <option value={room.ID}>{room.Name}</option>
        {/each}
      </select>
    </div>
  {/if}

  {#if selectedRoom}
    <div class="mb-4">
      <label
        for="quickshelfName"
        class="block text-sm font-medium text-gray-700">Quickshelf Name:</label
      >
      <input
        id="quickshelfName"
        type="text"
        required
        bind:value={quickshelfName}
        class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
      />
    </div>

    <button
      class="w-full py-2 px-4 bg-indigo-600 text-white font-medium rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      on:click={createQuickshelf}>Create Quickshelf</button
    >
  {/if}
</main>
