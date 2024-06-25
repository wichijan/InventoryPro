<script lang="ts">
  import { goto } from "$app/navigation";

  export let item: any;
</script>

<button
  class="nav-item flex-row w-full {!item.isLast
    ? 'border-r border-black'
    : ''} disabled:opacity-50 disabled:cursor-not-allowed"
  class:active={item.active}
  on:click={() => {
    goto(item.link);
  }}
>
  <div class="flex w-full">
    <svelte:component this={item.icon} class="mx-auto text-black " />
  </div>
  <span class="text-xs text-black">{item.name}</span>
</button>

<style>
  .nav-item {
    text-align: center;
    flex-grow: 1;
  }
  .active {
    color: #00aaff;
  }
</style>
