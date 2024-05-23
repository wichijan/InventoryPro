<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import { Button, Modal } from "flowbite-svelte";
  import { createEventDispatcher } from "svelte";
  let clickOutsideModal = false;

  export let warehouse;

  const dispatch = createEventDispatcher();

  let name = "";
</script>

<button
  class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full"
  on:click={() => (clickOutsideModal = true)}
>
  Neuen Raum erstellen
</button>

<Modal
  title="Raum erstellen für {warehouse.Name}"
  bind:open={clickOutsideModal}
  autoclose
  outsideclose
>
  <div class="flex flex-col">
    <label for="Name">Name</label>
    <input type="text" id="Name" bind:value={name} />
  </div>

  <svelte:fragment slot="footer">
    <button
      on:click={() => {
        fetch(API_URL + "rooms", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            Name: name,
            WarehouseID: warehouse.ID,
          }),
        }).then(() => {
          clickOutsideModal = false;
          dispatch("reload");
        });
      }}>Speichern</button
    >
    <Button
      color="alternative"
      on:click={() => {
        clickOutsideModal = false;
      }}>Abbrechen</Button
    >
  </svelte:fragment>
</Modal>
