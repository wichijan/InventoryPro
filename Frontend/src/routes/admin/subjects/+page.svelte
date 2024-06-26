<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Spinner from "$lib/templates/Spinner.svelte";
  import Swal from "sweetalert2";

  export let data;

  let subjects: any = data.subjects;
  $: subjects = subjects;

  $: {
    subjects = subjects.sort((a, b) => {
      if (a.Name < b.Name) {
        return -1;
      }
      if (a.Name > b.Name) {
        return 1;
      }
      return 0;
    });
  }
</script>

<div class=" min-h-screen text-gray-900 flex flex-col items-center p-6 mb-10">
  <div class="flex flex-col mt-10 mb-4 text-center">
    <div class="text-4xl font-bold text-gray-900">Fächer</div>
  </div>

  {#if subjects}
    <div
      class="relative overflow-x-auto shadow-lg rounded-lg w-full max-w-4xl bg-white"
    >
      <table class="w-full text-sm text-left text-gray-700">
        <thead class="text-xs uppercase bg-gray-200 text-gray-700">
          <tr>
            <th scope="col" class="px-6 py-3">Name</th>
            <th scope="col" class="text-center">Beschreibung</th>
            <th scope="col" class="px-6 py-3 text-right">Action</th>
          </tr>
        </thead>
        <tbody>
          {#each subjects as subject (subject.ID)}
            <tr
              class="odd:bg-gray-100 even:bg-gray-50 hover:bg-gray-300 transition-colors cursor-pointer"
              on:click={() => {
                //usign swal to update
                Swal.fire({
                  title: "Update Fach",
                  html: `
                    <input
                      id="name"
                      class="swal2-input"
                      value="${subject.Name}"
                    />
                    <input
                      id="description"
                      class="swal2-input"
                      value="${subject.Description}"
                    />
                  `,
                  showCancelButton: true,
                  confirmButtonText: `Update`,
                }).then((result) => {
                  if (result.isConfirmed) {
                    const name = document.getElementById("name").value;
                    const description =
                      document.getElementById("description").value;

                    fetch(API_URL + "subjects/" + subject.ID, {
                      method: "PUT",
                      headers: {
                        "Content-Type": "application/json",
                      },
                      credentials: "include",
                      body: JSON.stringify({
                        Name: name,
                        Description: description,
                      }),
                    }).then(() => {
                      Swal.fire("Updated!", "", "success");
                      subjects = subjects.map((s) => {
                        if (s.ID === subject.ID) {
                          return {
                            ...s,
                            Name: name,
                            Description: description,
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
                >{subject.Name}</th
              >
              <td class="text-center">{subject.Description}</td>
              <td class="px-6 py-4 text-right">
                <button
                  on:click|stopPropagation={() => {
                    Swal.fire({
                      title: "Delete Fach",
                      text: "Bist du sicher, dass du das Fach löschen möchtest?",
                      showCancelButton: true,
                      confirmButtonText: `Delete`,
                    }).then((result) => {
                      if (result.isConfirmed) {
                        fetch(API_URL + "subjects/" + subject.ID, {
                          method: "DELETE",
                          credentials: "include",
                        }).then(() => {
                          Swal.fire("Deleted!", "", "success");
                          subjects = subjects.filter(
                            (s) => s.ID !== subject.ID
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
    <button
      class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mt-4"
      on:click={() => {
        Swal.fire({
          title: "Create Fach",
          html: `
            <input id="name" class="swal2-input" placeholder="Name" />
            <input id="description" class="swal2-input" placeholder="Beschreibung" />
          `,
          showCancelButton: true,
          confirmButtonText: `Create`,
        }).then((result) => {
          if (result.isConfirmed) {
            const name = document.getElementById("name").value;
            const description = document.getElementById("description").value;

            fetch(API_URL + "subjects", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              credentials: "include",
              body: JSON.stringify({
                Name: name,
                Description: description,
              }),
            }).then((res) => {
              if (res.ok) {
                const newSubject = res.json();
                Swal.fire("Created!", "", "success");
                subjects = [
                  ...subjects,
                  {
                    ID: newSubject.ID,
                    Name: name,
                    Description: description,
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
