<script lang="ts">
  import { afterNavigate, goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { state } from "$lib/_services/WebSocket";

  export let icon: any;
  export let text: any;
  export let href: any;
  export let active: any;
  export let nCount: number = -1;

  let notificationDisplay = false;
  let notificationStyle = "";

  onMount(() => {
    // Check nCount and set notificationDisplay and notificationStyle accordingly
    if (nCount !== -1 && nCount !== 0) {
      notificationDisplay = true;
      notificationStyle = `bg-red-500 text-white`;
    }
  });

  function resetNotification() {
    nCount = 0;
  }
</script>

<button
  class="sidebar-icon relative group {active
    ? 'bg-[#a3b18a] text-black'
    : ''} disabled:opacity-50 disabled:cursor-not-allowed"
  on:click={() => {
    if (nCount > 0) {
      resetNotification();
    }
    goto(href);
  }}
>
  <svelte:component this={icon} />

  <span class="sidebar-tooltip group-hover:scale-100">{text}</span>

  {#if notificationDisplay}
    <span
      class="absolute top-0 right-0 inline-flex items-center justify-center px-2 py-1 mr-2 mt-2 text-xs font-bold leading-none rounded-full {notificationStyle}"
    >
      {nCount}
    </span>
  {/if}
</button>
