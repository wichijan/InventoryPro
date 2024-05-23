<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import EditModal from "$lib/templates/EditModal.svelte";

  export let data;

  type Item = {
    ID: string;
    Name: string;
    Description: string;
    ClassOne: boolean;
    ClassTwo: boolean;
    ClassThree: boolean;
    ClassFour: boolean;
    Damaged: boolean;
    DamagedDesc: string;
    Quantity: number;
    Status: string;
    Keywords: string[];
    Subject: string;
    Pictures: string[];
  };

  const item: Item = data.item;

  function save() {
    fetch(`${API_URL}items`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    }).then((res) => {
      if (res.ok) {
        console.log("Item updated");
      } else {
        console.log("Item not updated");
      }
    });
  }
</script>

<div class="my-5">
  <div
    class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 rounded-md"
    id="generalInformation"
  >
    <div class="mx-auto font-semibold text-2xl" id="header">
      General information
    </div>
    <div class="flex flex-row justify-between mt-3">
      <div class="ml-5 w-1/2" id="imgItem">
        <img
          src="https://via.placeholder.com/250"
          alt="item"
          class="rounded-md w-fit h-fit mx-auto object-cover"
        />
      </div>
      <div class="flex flex-col text-[#344e41] ml-12 font-semibold">
        <div class="">Name: {item.Name}</div>
        <div class="">Beschreibung: {item.Description}</div>
        <div class="">Menge: {item.Quantity}</div>
        <div class="">Status: {item.Status}</div>
        <EditModal
          buttonText="Bearbeiten"
          valuesToEdit={{
            Name: item.Name,
            Description: item.Description,
            Quantity: item.Quantity,
            Status: item.Status,
          }}
          on:save={(e) => {
            item.Name = e.detail.Name;
            item.Description = e.detail.Description;
            item.Quantity = e.detail.Quantity;
            item.Status = e.detail.Status;
            save();
          }}
        />
      </div>
    </div>
  </div>
  <div class="flex justify-between">
    <div
      class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 w-full rounded-md"
      id="damaged"
    >
      <div class="mx-auto font-semibold text-xl" id="header">Beschädigt</div>
      <div class="w-full">
        <div class="flex space-x-3">
          {#if item.Damaged}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6 rounded-md bg-red-400 text-white"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M6 18 18 6M6 6l12 12"
              />
            </svg>
            <div class="text-red-600">Beschädigt</div>
          {:else}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6 rounded-md bg-green-400 text-white"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M5 13l4 4L19 7"
              />
            </svg>
            <div class="text-green-600">In Ordnung</div>
          {/if}
        </div>
        {#if item.DamagedDesc}
          <br />
          <hr />
          <div class="mt-5">Beschreibung: <br /> {item.DamagedDesc}</div>
        {/if}

        <hr class="mt-5" />
        <EditModal
          buttonText="Bearbeiten"
          valuesToEdit={{
            Damaged: item.Damaged,
            DamagedDesc: item.DamagedDesc,
          }}
          on:save={(e) => {
            item.Damaged = e.detail.Damaged;
            item.DamagedDesc = e.detail.DamagedDesc;
            save();
          }}
        />
      </div>
    </div>
    <div
      class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 w-full rounded-md"
      id="classes"
    >
      <div class="mx-auto font-semibold text-xl" id="header">Klassen</div>
      <div class="w-full">
        <div class="flex-row">
          <div class={item.ClassOne ? "text-green-500" : "text-red-500"}>
            Klasse 1
          </div>
          <div class={item.ClassTwo ? "text-green-500" : "text-red-500"}>
            Klasse 2
          </div>
          <div class={item.ClassThree ? "text-green-500" : "text-red-500"}>
            Klasse 3
          </div>
          <div class={item.ClassFour ? "text-green-500" : "text-red-500"}>
            Klasse 4
          </div>
          <hr class="mt-5" />

          <EditModal
            buttonText="Bearbeiten"
            valuesToEdit={{
              ClassOne: item.ClassOne,
              ClassTwo: item.ClassTwo,
              ClassThree: item.ClassThree,
              ClassFour: item.ClassFour,
            }}
            on:save={(e) => {
              item.ClassOne = e.detail.ClassOne;
              item.ClassTwo = e.detail.ClassTwo;
              item.ClassThree = e.detail.ClassThree;
              item.ClassFour = e.detail.ClassFour;
              save();
            }}
          />
        </div>
      </div>
    </div>
  </div>
</div>
