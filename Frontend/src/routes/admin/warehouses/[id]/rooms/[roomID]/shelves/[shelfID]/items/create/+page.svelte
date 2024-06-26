<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  let types = ["book", "normalobject"];

  export let data;

  let keywords = data.keywords;
  let subjects = data.subjects;

  let selectedType = "none";

  $: selectedType = selectedType;

  let newItem = {
    BaseQuantityInShelf: 0,
    Name: "",
    ClassOne: false,
    ClassTwo: false,
    ClassThree: false,
    ClassFour: false,
    Damaged: false,
    DamagedDesc: "",
    Description: "",
    HintText: "",
    ShelfID: data.shelfID,
    subjects: [],
    keywords: [],
  };

  function handleCreate() {
    let body = {};
    if (selectedType === "book") {
      body = {
        Name: newItem.Name,
        Description: newItem.Description,
        BaseQuantityInShelf: newItem.BaseQuantityInShelf,
        Damaged: newItem.Damaged,
        DamagedDesc: newItem.DamagedDesc,
        Publisher: document.getElementById("publisher").value,
        ISBN: document.getElementById("isbn").value,
        Edition: document.getElementById("edition").value,
        ClassOne: newItem.ClassOne,
        ClassTwo: newItem.ClassTwo,
        ClassThree: newItem.ClassThree,
        ClassFour: newItem.ClassFour,
        HintText: newItem.HintText,
        Author: document.getElementById("author").value,
        itemTypeName: "book",
        RegularShelfId: data.shelfID,
      };
    } else {
      body = {
        Name: newItem.Name,
        Description: newItem.Description,
        BaseQuantityInShelf: newItem.BaseQuantityInShelf,
        Damaged: newItem.Damaged,
        DamagedDesc: newItem.DamagedDesc,
        BrokenObjects: Number(document.getElementById("brokenObjects").value),
        LostObjects: Number(document.getElementById("lostObjects").value),
        TotalObjects: Number(document.getElementById("totalObjects").value),
        UsefulObjects: Number(document.getElementById("usefulObjects").value),
        ClassOne: newItem.ClassOne,
        ClassTwo: newItem.ClassTwo,
        ClassThree: newItem.ClassThree,
        ClassFour: newItem.ClassFour,
        HintText: newItem.HintText,
        itemTypeName: "set_of_objects",
        RegularShelfId: data.shelfID,
      };
    }
    let url = `${API_URL}items/${selectedType === "book" ? "book" : "set-of-objects"}`;
    fetch(`${url}`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    })
      .then(async (response) => {
        if (response.ok) {
          const data = await response.json();
          uploadRest(newItem.keywords, newItem.subjects, data);

          Swal.fire({
            title: "Success",
            text: "Item wurde erfolgreich erstellt",
            icon: "success",
          }).then(() => {
            //ask if user wants to upload a picture, if not redirect to shelf
            Swal.fire({
              title: "Möchtest du ein Bild hochladen?",
              showDenyButton: true,
              confirmButtonText: `Ja`,
              denyButtonText: `Nein`,
            }).then((result) => {
              if (result.isConfirmed) {
                uploadPic(data);
              } else if (result.isDenied) {
                goto(
                  `/admin/warehouses/${data.warehouseID}/rooms/${data.roomID}/shelves/${data.shelfID}`
                );
              }
            });
          });
        } else {
          Swal.fire({
            title: "Error",
            text: "An error occurred while creating the item",
            icon: "error",
          });
        }
      })
      .catch((error) => {
        Swal.fire({
          title: "Error",
          text: "An error occurred while creating the item",
          icon: "error",
        });
      });
  }

  async function uploadRest(keywords, subjects, data) {
    const aKeyWordPromises = keywords.map((keyword) => {
      return fetch(`${API_URL}items/add-keyword`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ItemID: data,
          KeywordName: keyword,
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
          ItemID: data,
          SubjectName: subject,
        }),
      });
    });
    Promise.all([...aKeyWordPromises, ...aSubjectPromises]);
  }

  function uploadPic(itemID) {
    //Using swal, to ask for pic via form
    Swal.fire({
      title: "Bild hochladen",
      html: `
        <form id="uploadForm" enctype="multipart/form-data">
          <input type="hidden" id="itemID" name="id" value="${itemID}" />
          <input type="file" id="file" name="file" accept="image/*" required>
        </form>
      `,
      showCancelButton: true,
      confirmButtonText: `Hochladen`,
    }).then((result) => {
      if (result.isConfirmed) {
        const formData = new FormData();
        formData.append("file", document.getElementById("file").files[0]);
        formData.append("id", document.getElementById("itemID").value);
        fetch(`${API_URL}items-picture`, {
          method: "POST",
          credentials: "include",
          body: formData,
        }).then((response) => {
          if (response.ok) {
            Swal.fire({
              title: "Success",
              text: "Bild wurde erfolgreich hochgeladen",
              icon: "success",
            }).then(() => {
              goto(
                `/admin/warehouses/${data.warehouseID}/rooms/${data.roomID}/shelves/${data.shelfID}`
              );
            });
          } else {
            Swal.fire({
              title: "Error",
              text: "An error occurred while uploading the picture",
              icon: "error",
            });
          }
        });
      }
    });
  }
</script>

<main class="p-6 bg-gray-100 min-h-screen">
  <div class="container mx-auto">
    <h1 class="text-3xl font-bold mb-6">Create Item</h1>
    <div class="bg-white p-6 rounded shadow-md">
      <form on:submit|preventDefault={handleCreate}>
        <div class="mb-4">
          <label class="block text-gray-700" for="name">Name</label>
          <input
            type="text"
            id="name"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={newItem.Name}
          />
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="description"
            >Beschreibung</label
          >
          <textarea
            id="description"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={newItem.Description}
          ></textarea>
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="quantity">Basis Anzahl</label>
          <input
            id="quantity"
            type="number"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={newItem.BaseQuantityInShelf}
          />
        </div>
        <div class="flex gap-5">
          <div class="mb-4">
            <label class="block text-gray-700" for="classOne">Klasse 1</label>
            <input
              type="checkbox"
              id="classOne"
              class="mr-2"
              bind:checked={newItem.ClassOne}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classTwo">Klasse 2</label>
            <input
              type="checkbox"
              id="classTwo"
              class="mr-2"
              bind:checked={newItem.ClassTwo}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classThree">Klasse 3</label>
            <input
              type="checkbox"
              id="classThree"
              class="mr-2"
              bind:checked={newItem.ClassThree}
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700" for="classFour">Klasse 4</label>
            <input
              type="checkbox"
              id="classFour"
              class="mr-2"
              bind:checked={newItem.ClassFour}
            />
          </div>
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="damaged">Kaputt</label>
          <input
            type="checkbox"
            id="damaged"
            class="mr-2"
            bind:checked={newItem.Damaged}
          />
        </div>

        {#if newItem.Damaged}
          <div class="mb-4">
            <label class="block text-gray-700" for="dd"
              >Beschreibung des Schadens</label
            >
            <textarea
              id="dd"
              class="w-full p-2 border border-gray-300 rounded mt-1"
              bind:value={newItem.DamagedDesc}
            ></textarea>
          </div>
        {/if}

        <div class="mb-4">
          <label class="block text-gray-700" for="itemTypes">Item Type</label>
          <select
            id="itemTypes"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={selectedType}
          >
            <option value="none">Auswählen</option>
            {#each types as type}
              <option value={type}>{type}</option>
            {/each}
          </select>
        </div>

        <div class="mb-4 w-full ring-2 ring-gray-500 rounded-md py-2 px-2">
          {#if selectedType === "none"}
            <p class="text-gray-400">Warte auf Auswahl</p>
          {:else if selectedType === "book"}
            <div class="mb-4">
              <label class="block text-gray-700" for="publisher"
                >Publisher</label
              >
              <input
                type="text"
                id="publisher"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="isbn">ISBN</label>
              <input
                type="text"
                id="isbn"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="edition">Edition</label>
              <input
                type="text"
                id="edition"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="author">Author</label>
              <input
                type="text"
                id="author"
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
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="usefulObjects"
                >Nutzbare Objekte</label
              >
              <input
                type="number"
                id="usefulObjects"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
          {/if}
        </div>
        <div class="mb-4">
          <label class="block text-gray-700" for="subjects">Fächer</label>
          <select
            id="subjects"
            multiple
            bind:value={newItem.subjects}
            class="w-full p-2 border border-gray-300 rounded mt-1"
          >
            {#each subjects as subject}
              <option value={subject.Name}>{subject.Name}</option>
            {/each}
          </select>
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="keywords">Keywords</label>
          <select
            multiple
            id="subjects"
            bind:value={newItem.keywords}
            class="w-full p-2 border border-gray-300 rounded mt-1"
          >
            {#each keywords as keyword}
              <option value={keyword.Keyword}>{keyword.Keyword}</option>
            {/each}
          </select>
        </div>

        <div class="flex justify-end gap-5">
          <button
            type="submit"
            class="bg-blue-500 text-white px-4 py-2 rounded shadow-md"
          >
            Erstellen
          </button>
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
