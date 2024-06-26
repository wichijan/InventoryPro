<!-- Notifications.svelte -->
<script lang="ts">
  import { goto } from "$app/navigation";
  import { state } from "$lib/_services/WebSocket";
  import { onMount } from "svelte";

  let notificationCount = 0;
  let notifications = [];
  $: notifications = notifications;
  $: notificationCount = notificationCount;

  onMount(() => {
    state.subscribe((value) => {
      notifications = value.requests;
      notificationCount = value.requests.length;
    });
  });

  function getObject(notificationType) {
    if (notificationType === "Registration Request for Admins!") {
      return {
        title: "Registrierungsanfrage!",
        message:
          "Ein neuer Benutzer hat sich registriert und wartet auf die Freischaltung.",
        timestamp: "2 hours ago",
        iconUrl: "https://cdn-icons-png.flaticon.com/512/1828/1828665.png",
        url: "/admin/users",
      };
    }
  }
</script>

<div class="p-6 max-w-screen-md mx-auto">
  <h1 class="text-2xl font-bold mb-6">Notifications</h1>

  <!-- Notification count badge -->
  <div class="mt-6">
    <p class="text-sm font-medium text-gray-900">
      You have {notificationCount} notifications
    </p>
  </div>
  <!-- Notification list -->
  {#if notifications.length === 0}
    <p class="text-gray-500">No notifications</p>
  {:else}
    <ul class="divide-y divide-gray-200">
      {#each notifications as notification}
        <li class="py-4">
          <button
            class="flex space-x-3 text-left py-3 px-4 w-full bg-gray-200 rounded-md hover:scale-[102%] transition-transform duration-300 ease-in-out"
            on:click={() => {
              goto(getObject(notification)?.url);
            }}
          >
            <!-- Notification content -->
            <div class="flex-1 space-y-1">
              <p class="text-sm font-medium text-gray-900">
                {getObject(notification).title}
              </p>
              <p class="text-sm text-gray-500">
                {getObject(notification).message}
              </p>
              <p class="text-xs text-gray-400">
                {getObject(notification).timestamp}
              </p>
            </div>
          </button>
        </li>
      {/each}
    </ul>
  {/if}
</div>
