<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  export let data;

  let transferRequests = data.transferRequests;

  let users = data.users;

  function getUserFromId(userId) {
    return users.find((u) => u.ID === userId);
  }

  function acceptRequest(requestId) {
    //using swal to tell user if successful
    fetch(API_URL + "items/transfer-accept/" + requestId, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }).then((response) => {
      if (response.ok) {
        Swal.fire({
          icon: "success",
          title: "Anfrage angenommen",
          showConfirmButton: false,
          timer: 1500,
        });
      } else {
        Swal.fire({
          icon: "error",
          title: "Fehler beim Annehmen der Anfrage",
          showConfirmButton: false,
          timer: 1500,
        });
      }
    });
  }
</script>

<h1 class="container mx-auto text-2xl font-bold mb-6 mt-5">
  Transfer Requests
</h1>

<div class="container mx-auto p-6">
  {#if transferRequests}
    {#if transferRequests.length === 0}
      <p class="text-gray-500">Zurzeit keine Transferanfragen</p>
    {:else}
      {#each transferRequests as request}
        <div class="card">
          <div class="flex justify-between items-center mb-4">
            <div>
              <div class="text-lg font-semibold text-gray-900">
                Item ID: {request.ItemID}
              </div>
              <div class="text-sm text-gray-500">
                Datum der Anfrage: {request.RequestDate}
              </div>
              <div class="text-sm text-gray-500">
                Von: {getUserFromId(request.UserID)?.Username}
              </div>
              <div class="text-sm text-gray-500">
                Zu: {getUserFromId(request.TargetUserID)?.Username}
              </div>
            </div>
            <div>
              {#if !request.isAccepted}
                <button
                  class="button accept-button mr-2"
                  on:click={() => acceptRequest(request.TransferRequestID)}
                  >Annehmen</button
                >
              {:else}
                <span class="text-green-500 font-semibold">Angenommen</span>
              {/if}
            </div>
          </div>
        </div>
      {/each}
    {/if}
  {:else}
    <p class="text-gray-500">Keine Transferanfragen gefunden</p>
  {/if}
</div>

<style>
  .card {
    @apply bg-white shadow-xl rounded-2xl p-6 mb-6 transition duration-200 ease-in-out transform hover:scale-[1.01];
  }
  .button {
    @apply px-4 py-2 rounded-full text-white font-medium;
  }
  .accept-button {
    @apply bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50;
  }
</style>
