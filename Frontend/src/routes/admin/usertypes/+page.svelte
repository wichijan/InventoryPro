<script lang="ts">
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Spinner from "$lib/templates/Spinner.svelte";
  import Swal from "sweetalert2";

  export let data;

  let usertypes: any = data.usertypes;
  $: usertypes = usertypes;

  $: {
    if (usertypes) {
      usertypes = usertypes.sort((a, b) => {
        if (a.typeName < b.typeName) {
          return -1;
        }
        if (a.typeName > b.typeName) {
          return 1;
        }
        return 0;
      });
    } else {
      usertypes = [];
    }
  }
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6 mb-10">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">Usertypes</div>
  </div>

  {#if usertypes}
    {#if usertypes.length === 0}
      <div class="text-center text-gray-700">Keine Usertypes gefunden</div>
    {:else}
      <div
        class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
      >
        <table class="w-full text-sm text-left text-gray-700">
          <thead class="text-xs uppercase bg-gray-200 text-gray-700">
            <tr>
              <th scope="col" class="px-6 py-3">Typnamen</th>
              <th scope="col" class="px-6 py-3 text-right">Action</th>
            </tr>
          </thead>
          <tbody>
            {#each usertypes as typeName (typeName.ID)}
              <tr
                class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
                on:click={() => {
                  Swal.fire({
                    title: "Update Typnamen",
                    html: `
                        <input
                          id="typeName"
                          class="swal2-input"
                          value="${typeName.TypeName}"
                        />
                      `,
                    showCancelButton: true,
                    confirmButtonText: `Update`,
                  }).then((result) => {
                    if (result.isConfirmed) {
                      const tn = document.getElementById("typeName").value;

                      fetch(API_URL + "user-types", {
                        method: "PUT",
                        headers: {
                          "Content-Type": "application/json",
                        },
                        credentials: "include",
                        body: JSON.stringify({
                          ID: typeName.ID,
                          typeName: tn,
                        }),
                      }).then(() => {
                        Swal.fire("Updated!", "", "success");
                        browser ? location.reload() : null;
                      });
                    }
                  });
                }}
              >
                <th
                  scope="row"
                  class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                  >{typeName.TypeName}</th
                >
                <td class="px-6 py-4 text-right">
                  <button
                    on:click|stopPropagation={() => {
                      Swal.fire({
                        title: "Löschen eines Typnamens",
                        text: "Bist du sicher, dass du diesen Typnamen löschen möchtest?",
                        showCancelButton: true,
                        confirmButtonText: `Delete`,
                      }).then((result) => {
                        if (result.isConfirmed) {
                          fetch(API_URL + "user-types/" + typeName.ID, {
                            method: "DELETE",
                            credentials: "include",
                          }).then(() => {
                            Swal.fire("Deleted!", "", "success");
                            usertypes = usertypes.filter(
                              (s) => s.ID !== typeName.ID
                            );
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
          title: "Typname erstellen",
          html: `
            <input id="typeName" class="swal2-input" placeholder="Typname" />
          `,
          showCancelButton: true,
          confirmButtonText: `Create`,
        }).then((result) => {
          if (result.isConfirmed) {
            const typeName = document.getElementById("typeName").value;

            fetch(API_URL + "user-types", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              credentials: "include",
              body: JSON.stringify({
                Name: typeName,
              }),
            }).then(async (res) => {
              if (res.ok) {
                const newtypeName = await res.json();
                Swal.fire("Created!", "", "success");
                usertypes = [
                  ...usertypes,
                  {
                    ID: newtypeName,
                    TypeName: typeName,
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
