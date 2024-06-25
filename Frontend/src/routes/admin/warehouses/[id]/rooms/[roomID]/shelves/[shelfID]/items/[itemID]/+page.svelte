<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  export let data;

  let item = data.item;
  const copyItem = JSON.parse(JSON.stringify(item));
  $: item = item;

  let itemType = item.ItemTypes;
  function handleSave() {
    let type = item.ItemTypes.replaceAll("_", "-");

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
      itemTypes: item.ItemTypes,
    };

    if (itemType === "book") {
      body.Publisher = item.Publisher;
      body.Isbn = item.Isbn;
      body.Edition = item.Edition;
      body.Author = item.Author;
    } else {
      body.BrokenObjects = item.BrokenObjects;
      body.LostObjects = item.LostObjects;
      body.TotalObjects = item.TotalObjects;
      body.UsefulObjects = item.UsefulObjects;
    }

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
          }).then(() => {});
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

        <div class="flex gap-5">
          <div class="mb-4">
            <label class="block text-gray-700" for="classOne">Class One</label>
            <input
              type="checkbox"
              id="classOne"
              class="mr-2"
              bind:checked={item.ClassOne}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classTwo">Class Two</label>
            <input
              type="checkbox"
              id="classTwo"
              class="mr-2"
              bind:checked={item.ClassTwo}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classThree"
              >Class Three</label
            >
            <input
              type="checkbox"
              id="classThree"
              class="mr-2"
              bind:checked={item.ClassThree}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classFour">Class Four</label
            >
            <input
              type="checkbox"
              id="classFour"
              class="mr-2"
              bind:checked={item.ClassFour}
            />
          </div>
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

        <div class="mb-4 w-full ring-2 ring-gray-500 rounded-md py-2 px-2">
          {#if itemType === "book"}
            <div class="mb-4">
              <label class="block text-gray-700" for="publisher"
                >Publisher</label
              >
              <input
                type="text"
                id="publisher"
                bind:value={item.Publisher}
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="isbn">ISBN</label>
              <input
                type="text"
                id="isbn"
                bind:value={item.Isbn}
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="edition">Edition</label>
              <input
                type="text"
                bind:value={item.Edition}
                id="edition"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="author">Author</label>
              <input
                type="text"
                id="author"
                bind:value={item.Author}
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
          {:else}
            <div class="mb-4">
              <label class="block text-gray-700" for="brokenObjects"
                >Broken Objects</label
              >
              <input
                type="number"
                bind:value={item.BrokenObjects}
                id="brokenObjects"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="lostObjects"
                >Lost Objects</label
              >
              <input
                type="number"
                bind:value={item.LostObjects}
                id="lostObjects"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="totalObjects"
                >Total Objects</label
              >
              <input
                type="number"
                id="totalObjects"
                bind:value={item.TotalObjects}
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="usefulObjects"
                >Useful Objects</label
              >
              <input
                type="number"
                bind:value={item.UsefulObjects}
                id="usefulObjects"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
          {/if}
        </div>

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
