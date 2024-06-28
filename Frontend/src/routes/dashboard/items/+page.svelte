<script lang="ts">
  import { goto } from "$app/navigation";

  export let data;

  let userInfo = data.items[0] ?? {};

  let userItems = userInfo.BorrowedItems ?? [];
  let userReservations = userInfo.Reservations ?? [];
  console.log(userInfo);
  //filter out where Reservation -> ItemID is null
  userReservations = userReservations.filter(
    (r) => r.Reservation.ItemID !== null
  );
</script>

<div class="container mx-auto p-6">
  <h1 class="text-3xl font-bold mb-8 text-center">User Dashboard</h1>

  <div class="mb-12">
    <h2 class="text-2xl font-semibold mb-4 text-center">
      Ausgeliehene Gegenstände
    </h2>
    {#if userItems.length > 0}
      <ul class="space-y-6">
        {#each userItems as item}
          <li>
            <button
              class="bg-gray-100 shadow-lg rounded-lg p-6 flex items-start w-full text-left"
              on:click={() => {
                goto(`/items/${item.ID}`);
              }}
            >
              <img
                src={item.Picture ?? ""}
                alt={item.Name}
                class="w-24 h-24 mr-6 rounded-lg object-cover"
              />
              <div>
                <h3 class="text-xl font-medium">{item.Name ?? "No Name"}</h3>
                <p class="text-gray-600">
                  {item.Description ?? "No Description"}
                </p>
                <p class="text-gray-600">
                  Ausgeliehen am: {new Date(
                    item.BorrowedAt ?? Date.now()
                  ).toLocaleString()}
                </p>
                {#if item.Damaged}
                  <p class="text-red-500">
                    Damaged: {item.DamagedDescription ?? "Keine Beschreibung"}
                  </p>
                {/if}
              </div>
            </button>
          </li>
        {/each}
      </ul>
    {:else}
      <p class="text-center text-gray-500">Keine ausgeliehenen Gegenstände.</p>
    {/if}
  </div>

  <div>
    <h2 class="text-2xl font-semibold mb-4 text-center">Reservierungen</h2>
    {#if userReservations.length > 0}
      <ul class="space-y-6">
        {#each userReservations as { Item, Reservation }}
          <li>
            <button
              class="bg-gray-100 shadow-lg rounded-lg p-6 flex items-start w-full text-left"
              on:click={() => {
                goto(`/items/${Item.ID}`);
              }}
            >
              <img
                src={Item.Picture ?? ""}
                alt={Item.Name}
                class="w-24 h-24 mr-6 rounded-lg object-cover"
              />
              <div>
                <h3 class="text-xl font-medium">{Item.Name || "Kein Name"}</h3>
                <p class="text-gray-600">
                  {Item.Description ?? "Keine Beschreibung"}
                </p>
                <p class="text-gray-600">
                  Reserviert am: {new Date(
                    Reservation.ReservedAt ?? Date.now()
                  ).toLocaleString()}
                </p>
                <p class="text-gray-600">
                  Von: {new Date(
                    Reservation.TimeFrom ?? Date.now()
                  ).toLocaleString()} Bis: {new Date(
                    Reservation.TimeTo ?? Date.now()
                  ).toLocaleString()}
                </p>
                {#if Reservation.IsCancelled}
                  <p class="text-red-500">Cancelled</p>
                {/if}
              </div>
            </button>
          </li>
        {/each}
      </ul>
    {:else}
      <p class="text-center text-gray-500">No reservations available.</p>
    {/if}
  </div>
</div>

<style>
  .container {
    max-width: 800px;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
      "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji",
      "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
  }
</style>
