<script lang="ts">
  import { browser } from "$app/environment";
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import { onMount } from "svelte";
  import Swal from "sweetalert2";

  export let data;

  let item = data.item;
  const copyItem = JSON.parse(JSON.stringify(item));

  $: item = item;

  let subjects = data.subjects ? data.subjects : [];
  let keywords = data.keywords ? data.keywords : [];

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
      .then(async (response) => {
        if (response.ok) {
          setRest(item.Keywords, item.Subject);
          Swal.fire({
            title: "Success",
            text: "Item wurde erfolgreich aktualisiert",
            icon: "success",
          }).then(() => {
            browser ? location.reload() : null;
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

  if (!item.Subject) item.Subject = [];
  if (!item.Keywords) item.Keywords = [];
  let addableSubjects = [];
  let addableKeywords = [];

  onMount(async () => {
    addableSubjects = subjects.filter(
      (subject) => !item.Subject.find((s) => s.ID === subject.ID)
    );
    addableKeywords = keywords.filter(
      (keyword) => !item.Keywords.find((k) => k.ID === keyword.ID)
    );
    addableSubjects.sort((a, b) => a.Name.localeCompare(b.Name));
    addableKeywords.sort((a, b) => a.Keyword.localeCompare(b.Keyword));
  });

  function setRest(keywords, subjects) {
    const aKeyWordPromises = keywords.map((keyword) => {
      return fetch(`${API_URL}items/add-keyword`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ItemID: item.ID,
          KeywordName: keyword.Keyword,
        }),
      });
    });
    const aSubjectPromises = subjects.map((subject) => {
      return fetch(`${API_URL}items/add-subject`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ItemID: item.ID,
          SubjectName: subject.Name,
        }),
      });
    });

    const aRemoveKeyWordPromises = addableKeywords.map((keyword) => {
      return fetch(`${API_URL}items/remove-keyword`, {
        method: "DELETE",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ItemID: item.ID,
          KeywordName: keyword.Keyword,
        }),
      });
    });
    const aRemoveSubjectPromises = addableSubjects.map((subject) => {
      return fetch(`${API_URL}items/remove-subject`, {
        method: "DELETE",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ItemID: item.ID,
          SubjectName: subject.Name,
        }),
      });
    });

    Promise.all([
      ...aKeyWordPromises,
      ...aSubjectPromises,
      ...aRemoveKeyWordPromises,
      ...aRemoveSubjectPromises,
    ]);
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
            >Beschreibung</label
          >
          <textarea
            id="description"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={item.Description}
          ></textarea>
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="quantity"
            >Anzahl im Regal</label
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
            <label class="block text-gray-700" for="classOne">Klasse 1</label>
            <input
              type="checkbox"
              id="classOne"
              class="mr-2"
              bind:checked={item.ClassOne}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classTwo">Klasse 2</label>
            <input
              type="checkbox"
              id="classTwo"
              class="mr-2"
              bind:checked={item.ClassTwo}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classThree">Klasse 3</label>
            <input
              type="checkbox"
              id="classThree"
              class="mr-2"
              bind:checked={item.ClassThree}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classFour">Klasse 4</label>
            <input
              type="checkbox"
              id="classFour"
              class="mr-2"
              bind:checked={item.ClassFour}
            />
          </div>
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="damaged">Kaputt</label>
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
              >Beschreibung des Schadens</label
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
                >Kaputte Objekte</label
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
                >Verlorene Objekte</label
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
                >Totale Objekte</label
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
                >Nutzbare Objekte</label
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
          <label class="block text-gray-700" for="subjects"
            >Fächer (hinzugefügt)</label
          >
          <div class="w-full p-2 border border-gray-300 rounded mt-1">
            {#each item.Subject as subject}
              <button
                type="button"
                on:click={() => {
                  item.Subject = item.Subject.filter(
                    (s) => s.ID !== subject.ID
                  );
                  addableSubjects.push(subject);
                }}
                class="bg-blue-500 text-white px-2 py-1 rounded mr-2"
              >
                {subject.Name}
              </button>
            {/each}
          </div>
          <label class="block text-gray-700" for="subjects"
            >Fächer (können hinzugefügt werden)</label
          >
          <div class="w-full p-2 border border-gray-300 rounded mt-1">
            {#each addableSubjects as subject}
              <button
                type="button"
                on:click={() => {
                  item.Subject.push(subject);
                  addableSubjects = addableSubjects.filter(
                    (s) => s.ID !== subject.ID
                  );
                }}
                class="bg-blue-500 text-white px-2 py-1 rounded mr-2"
              >
                {subject.Name}
              </button>
            {/each}
          </div>

          <div class="mb-4">
            <label class="block text-gray-700" for="keywords"
              >Keywords (hinzugefügt)</label
            >
            <div class="w-full p-2 border border-gray-300 rounded mt-1">
              {#each item.Keywords as keyword}
                <button
                  type="button"
                  on:click={() => {
                    item.Keywords = item.Keywords.filter(
                      (k) => k.ID !== keyword.ID
                    );
                    addableKeywords.push(keyword);
                  }}
                  class="bg-blue-500 text-white px-2 py-1 rounded mr-2"
                >
                  {keyword.Keyword}
                </button>
              {/each}
            </div>
            <label class="block text-gray-700" for="subjects"
              >Keywords (können hinzugefügt werden)</label
            >
            <div class="w-full p-2 border border-gray-300 rounded mt-1">
              {#each addableKeywords as keyword}
                <button
                  type="button"
                  on:click={() => {
                    item.Keywords.push(keyword);
                    addableKeywords = addableKeywords.filter(
                      (k) => k.ID !== keyword.ID
                    );
                  }}
                  class="bg-blue-500 text-white px-2 py-1 rounded mr-2"
                >
                  {keyword.Keyword}
                </button>
              {/each}
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-gray-700">Details zur Reservierung</label>
            <ul class="list-disc ml-6">
              {#if item.Reservations}
                {#each item.Reservations as reservation}
                  <li class="mt-2">
                    <div>
                      <strong>Username:</strong>
                      {reservation.Username}
                    </div>
                    <div>
                      <strong>Anzahl:</strong>
                      {reservation.Quantity}
                    </div>
                    <div>
                      <strong>Datum der Reservierung:</strong>
                      {new Date(reservation.ReservationDate).toLocaleString()}
                    </div>
                    <div>
                      <strong>Von:</strong>
                      {new Date(reservation.TimeFrom).toLocaleString()}
                    </div>
                    <div>
                      <strong>Bis:</strong>
                      {new Date(reservation.TimeTo).toLocaleString()}
                    </div>
                    <div>
                      <strong>Wurde abgesagt:</strong>
                      {reservation.IsCancelled ? "Yes" : "No"}
                    </div>
                  </li>
                {/each}
              {:else}
                <li>Keine Reservierungen</li>
              {/if}
            </ul>
          </div>

          <div class="flex justify-end gap-5">
            <button
              type="button"
              on:click={() => {
                Swal.fire({
                  title: "Bitte bestätigen",
                  text: "Willst du die Änderungen verwerfen?",
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
              >Speichern</button
            >
          </div>
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
