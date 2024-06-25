<script lang="ts">
  import { browser } from "$app/environment";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  let email = "";
  let firstname = "";
  let jobtitle = "";
  let lastname = "";
  let phonenumber = "";
  let username = "";
  let usertypename = "";

  function createUser() {
    if (
      email &&
      firstname &&
      jobtitle &&
      lastname &&
      phonenumber &&
      username &&
      usertypename
    ) {
      fetch(`${API_URL}auth/generate-code`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          firstname,
          jobtitle,
          lastname,
          phonenumber,
          username,
          usertypename,
        }),
      }).then(async (response) => {
        if (response.ok) {
          const code = await response.json();
          //returns a code
          Swal.fire({
            icon: "success",
            title: "User created, the code is: " + code.RegistrationCode,
            footer: "The code has been copied to your clipboard",
          });
          //copy it to clipboard
          browser ? navigator.clipboard.writeText(code.RegistrationCode) : null;
        } else {
          Swal.fire({
            icon: "error",
            title: "Error",
            text: "Failed to create user",
          });
        }
      });
    } else {
      Swal.fire({
        icon: "error",
        title: "Error",
        text: "Please fill in all the fields",
      });
    }
  }

  export let data;
  let userTypes = data.userTypes;
</script>

<main class="p-6 bg-gray-100 min-h-screen">
  <h1 class="text-2xl font-bold mb-6 text-center">Create a New User</h1>

  <div class="mb-4">
    <label for="email" class="block text-sm font-medium text-gray-700"
      >Email:</label
    >
    <input
      id="email"
      type="email"
      required
      bind:value={email}
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
    />
  </div>

  <div class="mb-4">
    <label for="firstname" class="block text-sm font-medium text-gray-700"
      >First Name:</label
    >
    <input
      id="firstname"
      type="text"
      required
      bind:value={firstname}
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
    />
  </div>

  <div class="mb-4">
    <label for="lastname" class="block text-sm font-medium text-gray-700"
      >Last Name:</label
    >
    <input
      id="lastname"
      type="text"
      required
      bind:value={lastname}
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
    />
  </div>

  <div class="mb-4">
    <label for="username" class="block text-sm font-medium text-gray-700"
      >Username:</label
    >
    <input
      id="username"
      type="text"
      required
      bind:value={username}
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
    />
  </div>

  <div class="mb-4">
    <label for="phonenumber" class="block text-sm font-medium text-gray-700"
      >Phone Number:</label
    >
    <input
      id="phonenumber"
      type="text"
      required
      bind:value={phonenumber}
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
    />
  </div>

  <div class="mb-4">
    <label for="jobtitle" class="block text-sm font-medium text-gray-700"
      >Job Title:</label
    >
    <input
      id="jobtitle"
      type="text"
      required
      bind:value={jobtitle}
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
    />
  </div>

  <div class="mb-4">
    <label for="usertypename" class="block text-sm font-medium text-gray-700"
      >User Type:</label
    >
    <select
      bind:value={usertypename}
      class="block w-full p-2 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
      required
    >
      {#each userTypes as type}
        <option value={type.TypeName}>{type.TypeName}</option>
      {/each}
    </select>
  </div>

  <button
    class="w-full py-2 px-4 bg-indigo-600 text-white font-medium rounded-md shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
    on:click={createUser}>Create User</button
  >
</main>
