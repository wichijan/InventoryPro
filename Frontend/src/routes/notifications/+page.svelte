<script lang="ts">
  import { goto } from "$app/navigation";
  import { isUserAdmin } from "$lib/_services/UserService.js";
  import { state } from "$lib/_services/WebSocket";
  import { onMount } from "svelte";

  export let data;

  let user = data.users;

  let notificationCount = 0;
  let notifications = [];
  $: notifications = notifications;
  $: notificationCount = notificationCount;

  onMount(() => {
    state.subscribe((value) => {
      notifications = value.requests;
      notificationCount = value.requests.length;
    });

    isUserAdmin().then((isAdmin) => {
      if (!isAdmin) {
        //remove all notifications where is not item move request
        if (notifications) {
          notifications.forEach((notification) => {
            const object = buildObject(notification);
            if (object.title !== "Item move request created") {
              notifications = notifications.filter(
                (n) => n.Message !== "Item move request created"
              );
            }
          });
        }
      }
    });
  });

  function buildObject(notificationType) {
    const notification = JSON.parse(notificationType);
    return {
      title: notification.Message,
      message:
        notification.Message === "Item move request created"
          ? "Der Benutzer " +
            getUserFromId(notification.Sender).Username +
            " hat dir eine Anfrage geschickt"
          : "Ein neuer Benutzer wurde erstellt",
      timestamp: getDate(notification.TimeStamp),
      url:
        notification.Message === "Item move request created"
          ? "/requests"
          : "/admin/users",
    };
  }
  function getDate(timestamp) {
    const date = new Date(timestamp);
    return date.toLocaleString();
  }

  function getUserFromId(userId) {
    return user.find((u) => u.ID === userId);
  }
</script>

<div class="p-6 max-w-screen-md mx-auto">
  <h1 class="text-2xl font-bold mb-6">Notifications</h1>

  <!-- Notification count badge -->
  <div class="mt-6">
    <p class="text-sm font-medium text-gray-900">
      Du hast {notificationCount} neue Benachrichtigungen
    </p>
  </div>
  <!-- Notification list -->
  {#if notifications.length === 0}
    <p class="text-gray-500">Keine Benachrichtigungen</p>
  {:else}
    <ul class="divide-y divide-gray-200">
      {#each notifications as notification}
        <li class="py-4">
          <button
            class="flex space-x-3 text-left py-3 px-4 w-full bg-gray-200 rounded-md hover:scale-[102%] transition-transform duration-300 ease-in-out"
            on:click={() => {
              goto(buildObject(notification)?.url);
            }}
          >
            <!-- Notification content -->
            <div class="flex-1 space-y-1">
              <p class="text-sm font-medium text-gray-900">
                {buildObject(notification).title}
              </p>
              <p class="text-sm text-gray-500">
                {buildObject(notification).message}
              </p>
              <p class="text-xs text-gray-400">
                {buildObject(notification).timestamp}
              </p>
            </div>
          </button>
        </li>
      {/each}
    </ul>
  {/if}
</div>
