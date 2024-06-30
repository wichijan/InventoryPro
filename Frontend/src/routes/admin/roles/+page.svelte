<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Spinner from "$lib/templates/Spinner.svelte";
  import Swal from "sweetalert2";

  export let data;

  let roles: any = data.roles;
  $: roles = roles;

  $: {
    if (roles) {
      roles = roles.sort((a, b) => {
        if (a.RoleName < b.RoleName) {
          return -1;
        }
        if (a.RoleName > b.RoleName) {
          return 1;
        }
        return 0;
      });
    } else {
      roles = [];
    }
  }
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6 mb-10">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">Rollen</div>
  </div>

  {#if roles}
    {#if roles.length === 0}
      <div class="text-center text-gray-700">Keine Rollen vorhanden</div>
    {:else}
      <div
        class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
      >
        <table class="w-full text-sm text-left text-gray-700">
          <thead class="text-xs uppercase bg-gray-200 text-gray-700">
            <tr>
              <th scope="col" class="px-6 py-3">Rollenname</th>
              <th scope="col" class="px-6 py-3 text-right">Action</th>
            </tr>
          </thead>
          <tbody>
            {#each roles as role}
              <tr
                class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
                on:click={() => {
                  Swal.fire({
                    title: "Update Rollenname",
                    html: `
                        <input
                          id="RoleName"
                          class="swal2-input"
                          value="${role.RoleName}"
                        />
                      `,
                    showCancelButton: true,
                    confirmButtonText: `Update`,
                  }).then((result) => {
                    if (result.isConfirmed) {
                      const roleName =
                        document.getElementById("RoleName").value;

                      fetch(API_URL + "roles", {
                        method: "PUT",
                        headers: {
                          "Content-Type": "application/json",
                        },
                        credentials: "include",
                        body: JSON.stringify({
                          ID: role.ID,
                          RoleName: roleName,
                        }),
                      }).then(() => {
                        Swal.fire("Updated!", "", "success");
                        roles = roles.map((s) => {
                          if (s.ID === role.ID) {
                            return {
                              ...s,
                              RoleName: roleName,
                            };
                          }
                          return s;
                        });
                      });
                    }
                  });
                }}
              >
                <th
                  scope="row"
                  class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                  >{role.RoleName}</th
                >
                <td class="px-6 py-4 text-right">
                  <button
                    on:click|stopPropagation={() => {
                      Swal.fire({
                        title: "Delete Rollenname",
                        text: "Bist du sicher?",
                        showCancelButton: true,
                        confirmButtonText: `Delete`,
                      }).then((result) => {
                        if (result.isConfirmed) {
                          fetch(API_URL + "roles/" + role.ID, {
                            method: "DELETE",
                            credentials: "include",
                          }).then(() => {
                            Swal.fire("Deleted!", "", "success");
                            roles = roles.filter((s) => s.ID !== role.ID);
                          });
                        }
                      });
                    }}
                    class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
                  >
                    Löschen
                  </button>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
    <button
      class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mt-4"
      on:click={() => {
        Swal.fire({
          title: "Rollenname erstellen",
          html: `
            <input id="RoleName" class="swal2-input" placeholder="Rollenname" />
          `,
          showCancelButton: true,
          confirmButtonText: `Create`,
        }).then((result) => {
          if (result.isConfirmed) {
            const roleName = document.getElementById("RoleName").value;

            fetch(API_URL + "roles", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              credentials: "include",
              body: JSON.stringify({
                roleName,
              }),
            }).then(async (res) => {
              if (res.ok) {
                const newRoleName = await res.json();
                Swal.fire("Created!", "", "success");
                roles = [
                  ...roles,
                  {
                    ID: newRoleName,
                    RoleName: roleName,
                  },
                ];
              }
            });
          }
        });
      }}
    >
      Erstellen
    </button>
  {:else}
    <Spinner />
  {/if}
</div>
