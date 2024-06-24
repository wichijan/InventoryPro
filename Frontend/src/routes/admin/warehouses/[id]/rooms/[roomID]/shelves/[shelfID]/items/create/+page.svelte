<script lang="ts">
  import { goto } from "$app/navigation";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  let types = ["book", "normalobject"];

  export let data;

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
        ItemTypes: "book",
        RegularShelfId: data.shelfID,
      };
    } else {
      body = {
        Name: newItem.Name,
        Description: newItem.Description,
        BaseQuantityInShelf: newItem.BaseQuantityInShelf,
        Damaged: newItem.Damaged,
        DamagedDesc: newItem.DamagedDesc,
        BrokenObjects: document.getElementById("brokenObjects").value,
        LostObjects: document.getElementById("lostObjects").value,
        TotalObjects: document.getElementById("totalObjects").value,
        UsefulObjects: document.getElementById("usefulObjects").value,
        ClassOne: newItem.ClassOne,
        ClassTwo: newItem.ClassTwo,
        ClassThree: newItem.ClassThree,
        ClassFour: newItem.ClassFour,
        HintText: newItem.HintText,
        ItemTypes: "set-of-objects",
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
      .then((response) => {
        if (response.ok) {
          Swal.fire({
            title: "Success",
            text: "Item has been created",
            icon: "success",
          }).then(() => {
            Swal.fire({
              title: "Success",
              text: "Item has been created",
              icon: "success",
            }).then(() => {});
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

  function handleSubjectChange(event) {
    newItem.subjects = event.target.value
      .split(",")
      .map((subject) => subject.trim());
  }

  function handleKeywordChange(event) {
    newItem.keywords = event.target.value
      .split(",")
      .map((keyword) => keyword.trim());
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
            >Description</label
          >
          <textarea
            id="description"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={newItem.Description}
          ></textarea>
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="quantity">Base Quantity</label
          >
          <input
            id="quantity"
            type="number"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            bind:value={newItem.BaseQuantityInShelf}
          />
        </div>

        <div class="mb-4">
          <label class="block text-gray-700" for="damaged">Damaged</label>
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
              >Damaged Description</label
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
            <option value="none">Select an item type</option>
            {#each types as type}
              <option value={type}>{type}</option>
            {/each}
          </select>
        </div>

        <div class="mb-4 w-full ring-2 ring-gray-500 rounded-md py-2 px-2">
          {#if selectedType === "none"}
            <p class="text-gray-400">Waiting for item type selection...</p>
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
                >Broken Objects</label
              >
              <input
                type="number"
                id="brokenObjects"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="lostObjects"
                >Lost Objects</label
              >
              <input
                type="number"
                id="lostObjects"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="totalObjects"
                >Total Objects</label
              >
              <input
                type="number"
                id="totalObjects"
                class="w-full p-2 border border-gray-300 rounded mt-1"
              />
            </div>
            <div class="mb-4">
              <label class="block text-gray-700" for="usefulObjects"
                >Useful Objects</label
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
          <label class="block text-gray-700" for="subjects">Subjects</label>
          <input
            type="text"
            id="subjects"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            placeholder="Enter subjects separated by commas"
            on:change={handleSubjectChange}
          />
        </div>

        {#if newItem.subjects.length > 0}
          <div class="mb-4">
            <label class="block text-gray-700">Selected Subjects:</label>
            <ul>
              {#each newItem.subjects as subject, index}
                <li key={index} class="text-gray-700">{subject}</li>
              {/each}
            </ul>
          </div>
        {/if}

        <div class="mb-4">
          <label class="block text-gray-700" for="keywords">Keywords</label>
          <input
            type="text"
            id="keywords"
            class="w-full p-2 border border-gray-300 rounded mt-1"
            placeholder="Enter keywords separated by commas"
            on:change={handleKeywordChange}
          />
        </div>

        {#if newItem.keywords.length > 0}
          <div class="mb-4">
            <label class="block text-gray-700">Selected Keywords:</label>
            <ul>
              {#each newItem.keywords as keyword, index}
                <li key={index} class="text-gray-700">{keyword}</li>
              {/each}
            </ul>
          </div>
        {/if}

        <div class="flex justify-end gap-5">
          <button
            type="submit"
            class="bg-blue-500 text-white px-4 py-2 rounded shadow-md"
          >
            Create
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
