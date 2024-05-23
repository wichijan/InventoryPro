<script>
  export let data;

  const rooms = data.rooms;
</script>

<div class="grid grid-cols-3 gap-4">
  {#each rooms as room}
    <div class="bg-tertiary p-4 rounded-md">
      <h1 class="text-2xl font-bold">{room.Name}</h1>
    </div>
  {/each}
</div>
