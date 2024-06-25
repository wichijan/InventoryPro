<script>
  import Swal from "sweetalert2";
  import { API_URL } from "$lib/_services/ShelfService";

  export let data;

  const registrationRequests = data.registrationRequests;
  const roles = data.roles;
  let users = data.users;

  let vUsers = users.filter((user) => user.IsActive);
</script>

<div class="flex flex-col items-center w-full">
  <div class="mt-10 mb-4">
    <h1 class="text-3xl font-bold text-black">Users</h1>
  </div>
  <div
    class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 bg-tertiary rounded px-2 py-4 mb-10"
  >
    {#each vUsers as user}
      <button
        class="max-w-sm rounded overflow-hidden shadow-lg m-3 bg-white px-5 py-5 hover:shadow-xl duration-300 ease-in-out transform hover:scale-[102%] flex flex-col"
        on:click={() => {
          Swal.fire({
            title:
              "Möchtest du diesen User einem Role hinzufügen oder entfernen?",
            showDenyButton: true,
            showCancelButton: true,
            confirmButtonText: `Add`,
            denyButtonText: `Remove`,
          }).then((result) => {
            if (result.isConfirmed) {
              let inputOptions = {};
              roles.forEach((role) => {
                if (!user.Roles.find((r) => r.ID === role.ID))
                  inputOptions[role.ID] = role.RoleName;
              });

              Swal.fire({
                title: "Wähle eine Rolle aus",
                input: "select",
                inputOptions: inputOptions,
                showCancelButton: true,
                confirmButtonText: `Add`,
              }).then((result) => {
                if (result.isConfirmed) {
                  if (
                    result.value &&
                    !user.Roles.find((role) => role.ID === result.value)
                  ) {
                    fetch(`${API_URL}user-roles/add-role`, {
                      method: "POST",
                      credentials: "include",
                      headers: {
                        "Content-Type": "application/json",
                      },
                      body: JSON.stringify({
                        UserID: user.ID,
                        RoleID: result.value,
                      }),
                    }).then((response) => {
                      if (response.ok) {
                        Swal.fire("Role hinzugefügt!", "", "success");
                      } else {
                        Swal.fire("Error adding role", "", "error");
                      }
                    });
                  }
                }
              });
            } else if (result.isDenied) {
              if (user.Roles.length === 1) {
                Swal.fire(
                  "User muss mindestens eine Rolle haben",
                  "Dieser User hat nur eine Rolle",
                  "error"
                );
                return;
              }

              let inputOptions = {};
              user.Roles.forEach((role) => {
                inputOptions[role.ID] = role.RoleName;
              });

              Swal.fire({
                title: "Wähle eine Rolle aus",
                input: "select",
                inputOptions: inputOptions,
                showCancelButton: true,
                confirmButtonText: `Remove`,
              }).then((result) => {
                if (result.isConfirmed) {
                  if (
                    result.value &&
                    user.Roles.find((role) => role.ID === result.value)
                  ) {
                    fetch(`${API_URL}user-roles/remove-role`, {
                      method: "DELETE",
                      credentials: "include",
                      headers: {
                        "Content-Type": "application/json",
                      },
                      body: JSON.stringify({
                        UserID: user.ID,
                        RoleID: result.value,
                      }),
                    }).then((response) => {
                      if (response.ok) {
                        Swal.fire("Role entfernt!", "", "success");
                      } else {
                        Swal.fire("Error removing role", "", "error");
                      }
                    });
                  }
                }
              });
            }
          });
        }}
      >
        <img
          class="mx-auto rounded w-12 h-12 object-cover mb-5"
          src={user.ProfilePicture}
          alt={user.FirstName + " " + user.LastName}
        />
        <div
          class="flex flex-col px-6 py-4 w-full ring-1 rounded-md ring-gray-500 items-start my-auto"
        >
          <div class="font-bold text-xl mb-2">
            {user.FirstName}
            {user.LastName}
          </div>
          <p class="text-gray-700 text-base">
            {user.Email}
          </p>
          <p class="text-gray-700 text-base">
            {user.JobTitle}
          </p>
          <p class="text-gray-700 text-base">
            {user.PhoneNumber}
          </p>
          <p class="text-gray-700 text-base">
            {user.UserTypeName}
          </p>
          <p class="text-gray-700 text-base">
            {user.Username}
          </p>
        </div>
      </button>
    {/each}
  </div>

  <div class="mt-10 mb-4">
    <h1 class="text-3xl font-bold text-black">User Requests</h1>
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
                fetch(`${API_URL}auth/decline-registration/${user.UserID}`, {
                  method: "DELETE",
                  credentials: "include",
                }).then((response) => {
                  if (response.ok) {
                    Swal.fire("User abgelehnt!", "", "success");
                  } else {
                    Swal.fire("Error denying user", "", "error");
                  }
                });
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
            <p class="text-gray-700 text-base">
              {users.find((u) => u.ID === user.UserID).Username}
            </p>
          </div>
        </button>
      {/each}
    </div>
  {/if}
</div>
