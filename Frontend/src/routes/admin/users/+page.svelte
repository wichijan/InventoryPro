<script>
  import Swal from "sweetalert2";
  import { API_URL } from "$lib/_services/ShelfService";

  export let data;

  const registrationRequests = data.registrationRequests;
</script>

<div class="flex flex-col items-center w-full">
  <div class="mt-10 mb-4">
    <h1 class="text-3xl font-bold text-black">User requests</h1>
  </div>
  {#if !registrationRequests}
    <div class="text-lg">No user requests</div>
  {/if}
  {#if registrationRequests}
    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 bg-tertiary rounded px-2 py-4"
    >
      {#each registrationRequests as user (user.UserID)}
        <button
          class="max-w-sm rounded overflow-hidden shadow-lg m-3 bg-white px-5 py-5 hover:shadow-xl duration-300 ease-in-out transform hover:scale-[102%]"
          on:click={() => {
            Swal.fire({
              title: "Möchtest du diesen User akzeptieren oder ablehnen?",
              showDenyButton: true,
              showCancelButton: true,
              confirmButtonText: `Accept`,
              denyButtonText: `Reject`,
            }).then((result) => {
              if (result.isConfirmed) {
                console.log(user);
                fetch(`${API_URL}auth/accept-registration/${user.UserID}`, {
                  method: "POST",
                  credentials: "include",
                }).then((response) => {
                  if (response.ok) {
                    Swal.fire("User angenommen!", "", "success");
                  } else {
                    Swal.fire("Error accepting user", "", "error");
                  }
                });
              } else if (result.isDenied) {
                Swal.fire("This feature is WIP!", "", "info");
              }
            });
          }}
        >
          <img
            class="mx-auto rounded w-12 h-12 object-cover"
            src="https://via.placeholder.com/150"
            alt="Image"
          />
          <div class="px-6 py-4">
            <div class="font-bold text-xl mb-2">
              {new Date(user.RequestTime).toLocaleString()}
            </div>
            <p class="text-gray-700 text-base">{user.UserID}</p>
          </div>
        </button>
      {/each}
    </div>
  {/if}
</div>
