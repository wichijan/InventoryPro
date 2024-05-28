<script>
  import { goto } from "$app/navigation";

  export let data;

  const rooms = data.rooms;
</script>

<div class="grid grid-cols-3 gap-4 ml-10 mt-4">
  {#each rooms as room}
    <button
      class="bg-tertiary p-4 rounded-md hover:shadow-lg duration-300 hover:bg-[#d5bdaf] hover:text-white"
      on:click={() => {
        goto(`/overview/rooms/${room.ID}`);
      }}
    >
      <h1 class="text-2xl font-bold">{room.Name}</h1>
    </button>
  {/each}
</div>
