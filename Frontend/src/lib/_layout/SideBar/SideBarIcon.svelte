<script lang="ts">
  import { goto } from "$app/navigation";

  export let icon: any;
  export let text: any;
  export let href: any;
  export let active: any;
</script>

<button
  class="sidebar-icon group {active ? 'bg-green-500 text-black' : ''}"
  on:click={() => {
    goto(href);
  }}
>
  <svelte:component this={icon} />

  <span class="sidebar-tooltip group-hover:scale-100">{text}</span>
</button>
