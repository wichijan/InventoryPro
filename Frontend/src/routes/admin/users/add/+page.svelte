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

  let telephoneError: string = "";

  const validatePhonenumber = () => {
    const regex = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/;
    telephoneError = regex.test(phonenumber)
      ? ""
      : "False Telefonnummer! Die Telefonnummer muss in dem Format +1234567890 or 1234567890 sein.";
  };

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
          Swal.fire({
            icon: "success",
            title:
              "User wurde erstellt, der einladungscode lautet:" +
              code.RegistrationCode,
            footer: "Der Code wurde in die Zwischenablage kopiert",
          });
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
        text: "Bitte fülle alle Felder aus.",
      });
    }
  }

  export let data;
  let userTypes = data.userTypes;
</script>

<main class="p-6 bg-gray-100 min-h-screen">
  <h1 class="text-2xl font-bold mb-6 text-center">Erstelle einen neuen User</h1>

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
      on:input={validatePhonenumber}
      required
      bind:value={phonenumber}
      class="mt-1 block w-full px-3 py-2 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
    />
    <p class="text-red-500 text-sm mt-2">{telephoneError}</p>
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
    on:click={createUser}>Erstelle den User</button
  >
</main>
