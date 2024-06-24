<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  export let data;

  let item = data.item;
  const copyItem = JSON.parse(JSON.stringify(item));
  $: item = item;

  function handleSave() {
    let type = item.ItemTypes.replace("_", "-");

    let body = {
      ID: item.ID,
      ItemID: item.ID,
      HintText: item.HintText,
      Name: item.Name,
      Description: item.Description,
      QuantityInShelf: item.QuantityInShelf,
      Damaged: item.Damaged,
      DamagedDesc: item.DamagedDescription,
      regularShelfID: item.RegularShelfID,
      ClassOne: item.ClassOne,
      ClassTwo: item.ClassTwo,
      ClassThree: item.ClassThree,
      ClassFour: item.ClassFour,
      ItemTypes: item.ItemTypes,
    };
    fetch(`${API_URL}items/${type}`, {
      method: "PUT",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    })
      .then((response) => {
        if (response.ok) {
          Swal.fire({
            title: "Success",
            text: "Item has been updated",
            icon: "success",
          }).then(() => {
            goto(
              `/admin/warehouses/${item.WarehouseID}/rooms/${item.RoomID}/shelves/${item.ShelfID}`
            );
          });
        } else {
          Swal.fire({
            title: "Error",
            text: "An error occurred while updating the item",
            icon: "error",
          });
        }
      })
      .catch((error) => {
        Swal.fire({
          title: "Error",
          text: "An error occurred while updating the item",
          icon: "error",
        });
      });
  }
</script>

<main class="p-6 bg-gray-100 min-h-screen">
  <div class="container mx-auto">
    <h1 class="text-3xl font-bold mb-6">Edit Item</h1>
    <div class="bg-white p-6 rounded shadow-md">
      <form on:submit|preventDefault={handleSave}>
        <div class="mb-4">
          <label class="block text-gray-700" for="name">Name</label>
          <input
            type="text"
            id="name"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={item.Name}
          />
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="description"
            >Description</label
          >
          <textarea
            id="description"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={item.Description}
          ></textarea>
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="quantity"
            >Quantity In Shelf</label
          >
          <input
            id="quantity"
            type="number"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={item.QuantityInShelf}
          />
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="damaged">Damaged</label>
          <input
            type="checkbox"
            id="damaged"
            class="mr-2"
            bind:checked={item.Damaged}
          />
        </div>

        {#if item.Damaged}
          <div class="mb-4">
            <label class="block text-gray-700" for="dd"
              >Damaged Description</label
            >
            <textarea
              id="dd"
              class="w-full p-2 border border-gray-300 rounded mt-1"
              bind:value={item.DamagedDescription}
            ></textarea>
          </div>
        {/if}

        <div class="mb-4">
          <label class="block text-gray-700">Reservation Details</label>
          <ul class="list-disc ml-6">
            {#if item.Reservations}
              {#each item.Reservations as reservation}
                <li class="mt-2">
                  <div>
                    <strong>Username:</strong>
                    {reservation.Username}
                  </div>
                  <div>
                    <strong>Quantity:</strong>
                    {reservation.Quantity}
                  </div>
                  <div>
                    <strong>Reservation Date:</strong>
                    {new Date(reservation.ReservationDate).toLocaleString()}
                  </div>
                  <div>
                    <strong>Time From:</strong>
                    {new Date(reservation.TimeFrom).toLocaleString()}
                  </div>
                  <div>
                    <strong>Time To:</strong>
                    {new Date(reservation.TimeTo).toLocaleString()}
                  </div>
                  <div>
                    <strong>Is Cancelled:</strong>
                    {reservation.IsCancelled ? "Yes" : "No"}
                  </div>
                </li>
              {/each}
            {:else}
              <li>No reservations</li>
            {/if}
          </ul>
        </div>

        <div class="flex justify-end gap-5">
          <button
            type="button"
            on:click={() => {
              Swal.fire({
                title: "Are you sure?",
                text: "Do you want to cancel editing this item?",
                icon: "warning",
                showCancelButton: true,
                confirmButtonText: "Yes",
                cancelButtonText: "No",
              }).then((result) => {
                if (result.isConfirmed) {
                  item = copyItem;
                }
              });
            }}
            class="bg-red-500 text-white px-4 py-2 rounded shadow-md"
            >Abbrechen</button
          >
          <button
            type="submit"
            class="bg-blue-500 text-white px-4 py-2 rounded shadow-md"
            >Save</button
          >
        </div>
      </form>
    </div>
  </div>
</main>

<style>
  main {
    font-family:
      system-ui,
      -apple-system,
      "Segoe UI",
      Roboto,
      "Helvetica Neue",
      Arial,
      "Noto Sans",
      sans-serif,
      "Apple Color Emoji",
      "Segoe UI Emoji",
      "Segoe UI Symbol",
      "Noto Color Emoji";
  }
</style>
