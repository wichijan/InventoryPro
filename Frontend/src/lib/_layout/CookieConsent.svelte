<script>
  import { onMount } from "svelte";

  let isVisible = false;
  let preferences = {
    necessary: true,
    analytics: true,
    marketing: true,
  };

  const acceptAll = () => {
    preferences.analytics = true;
    preferences.marketing = true;
    savePreferences();
    isVisible = false;
  };

  const savePreferences = () => {
    localStorage.setItem("cookiePreferences", JSON.stringify(preferences));
    isVisible = false;
  };

  const loadPreferences = () => {
    const savedPreferences = localStorage.getItem("cookiePreferences");
    if (savedPreferences) {
      preferences = JSON.parse(savedPreferences);
      isVisible = false;
    } else {
      isVisible = true;
    }
  };

  onMount(() => {
    loadPreferences();
  });
</script>

{#if isVisible}
  <div
    class="fixed bottom-4 right-4 bg-white shadow-lg rounded-lg p-6 max-w-sm z-50"
  >
    <h2 class="text-lg font-bold mb-4">Cookie Preferences</h2>
    <div class="mb-4">
      <label class="flex items-center">
        <input
          type="checkbox"
          bind:checked={preferences.necessary}
          disabled
          class="mr-2"
        />
        <span>Necessary Cookies</span>
      </label>
      <p class="text-sm text-gray-600 ml-6">
        These cookies are essential for the website to function.
      </p>
    </div>
    <div class="mb-4">
      <label class="flex items-center">
        <input
          type="checkbox"
          bind:checked={preferences.analytics}
          class="mr-2"
        />
        <span>Analytics Cookies</span>
      </label>
      <p class="text-sm text-gray-600 ml-6">
        These cookies help us understand how visitors interact with our website.
      </p>
    </div>
    <div class="mb-4">
      <label class="flex items-center">
        <input
          type="checkbox"
          bind:checked={preferences.marketing}
          class="mr-2"
        />
        <span>Marketing Cookies</span>
      </label>
      <p class="text-sm text-gray-600 ml-6">
        These cookies are used to deliver advertising that is more relevant to
        you.
      </p>
    </div>
    <div class="flex justify-end">
      <button
        on:click={acceptAll}
        class="bg-blue-500 text-white px-4 py-2 rounded mr-2">Accept All</button
      >
      <button
        on:click={savePreferences}
        class="bg-gray-500 text-white px-4 py-2 rounded"
        >Save Preferences</button
      >
    </div>
  </div>
{/if}
