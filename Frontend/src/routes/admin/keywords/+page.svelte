<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Spinner from "$lib/templates/Spinner.svelte";
  import Swal from "sweetalert2";

  export let data;

  let keywords: any = data.keywords;
  $: keywords = keywords;

  $: {
    if (keywords) {
      keywords = keywords.sort((a, b) => {
        if (a.Keyword < b.Keyword) {
          return -1;
        }
        if (a.Keyword > b.Keyword) {
          return 1;
        }
        return 0;
      });
    } else {
      keywords = [];
    }
  }
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6 mb-10">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">Keywords</div>
  </div>

  {#if keywords}
    {#if keywords.length === 0}
      <div class="text-center text-gray-700">Es wurden keine Keywords gefunden</div>
    {:else}
      <div
        class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
      >
        <table class="w-full text-sm text-left text-gray-700">
          <thead class="text-xs uppercase bg-gray-200 text-gray-700">
            <tr>
              <th scope="col" class="px-6 py-3">Keyword</th>
              <th scope="col" class="px-6 py-3 text-right">Action</th>
            </tr>
          </thead>
          <tbody>
            {#each keywords as keyword (keyword.ID)}
              <tr
                class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
                on:click={() => {
                  Swal.fire({
                    title: "Update keyword",
                    html: `
                      <input
                        id="Keyword"
                        class="swal2-input"
                        value="${keyword.Keyword}"
                      />
                    `,
                    showCancelButton: true,
                    confirmButtonText: `Update`,
                  }).then((result) => {
                    if (result.isConfirmed) {
                      const Keyword = document.getElementById("Keyword").value;

                      fetch(API_URL + "keywords", {
                        method: "PUT",
                        headers: {
                          "Content-Type": "application/json",
                        },
                        credentials: "include",
                        body: JSON.stringify({
                          ID: keyword.ID,
                          Keyword: Keyword,
                        }),
                      }).then(() => {
                        Swal.fire("Updated!", "", "success");
                        keywords = keywords.map((s) => {
                          if (s.ID === keyword.ID) {
                            return {
                              ...s,
                              Keyword: Keyword,
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
                  >{keyword.Keyword}</th
                >
                <td class="px-6 py-4 text-right">
                  <button
                    on:click|stopPropagation={() => {
                      Swal.fire({
                        title: "Delete keyword",
                        text: "Are you sure you want to delete this keyword?",
                        showCancelButton: true,
                        confirmButtonText: `Delete`,
                      }).then((result) => {
                        if (result.isConfirmed) {
                          fetch(API_URL + "keywords/" + keyword.ID, {
                            method: "DELETE",
                            credentials: "include",
                          }).then(() => {
                            Swal.fire("Deleted!", "", "success");
                            keywords = keywords.filter(
                              (s) => s.ID !== keyword.ID,
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
          title: "Create keyword",
          html: `
          <input id="Keyword" class="swal2-input" placeholder="Keyword" />
        `,
          showCancelButton: true,
          confirmButtonText: `Create`,
        }).then((result) => {
          if (result.isConfirmed) {
            const Keyword = document.getElementById("Keyword").value;

            fetch(API_URL + "keywords", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              credentials: "include",
              body: JSON.stringify({
                Name: Keyword,
              }),
            }).then((res) => {
              if (res.ok) {
                const newkeyword = res.json();
                Swal.fire("Created!", "", "success");
                keywords = [
                  ...keywords,
                  {
                    ID: newkeyword.ID,
                    Keyword: Keyword,
                  },
                ];
              }
            });
          }
        });
      }}
    >
      Create
    </button>
  {:else}
    <Spinner />
  {/if}
</div>
