<script lang="ts">
  import SideBarIcon from "./SideBarIcon.svelte";

  export let items: any[];
</script>

<div
  class="fixed top-0 left-0 h-screen w-16 m-0 flex flex-col bg-white text-white shadow-lg z-50"
>
  {#each items as item}
    <SideBarIcon
      icon={item.icon}
      text={item.name}
      href={item.link}
      active={item.active}
      nCount={item.nCount}
    />
  {/each}
</div>
