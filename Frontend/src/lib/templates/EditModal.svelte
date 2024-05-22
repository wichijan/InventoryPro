<script lang="ts">
  import { Button, Modal } from "flowbite-svelte";
  import { createEventDispatcher } from "svelte";
  let clickOutsideModal = false;

  export let buttonText;
  export let valuesToEdit;

  const dispatch = createEventDispatcher();
</script>

<button
  class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full"
  on:click={() => (clickOutsideModal = true)}
>
  {buttonText}
</button>

<Modal
  title="Terms of Service"
  bind:open={clickOutsideModal}
  autoclose
  outsideclose
>
  {#each Object.keys(valuesToEdit) as key}
    <div class="flex flex-col">
      <label for={key}>{key}</label>
      <input type="text" id={key} bind:value={valuesToEdit[key]} />
    </div>
  {/each}

  <svelte:fragment slot="footer">
    <button
      on:click={() => {
        dispatch("save", valuesToEdit);
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
