<script lang="ts">
  import { browser } from "$app/environment";
  import { API_URL } from "$lib/_services/ShelfService";
  import Swal from "sweetalert2";

  import { PersonCircle } from "svelte-bootstrap-icons";

  export let data;

  let userInfo = data.userData ?? {};

  let passwordError: string = "";
  let telephoneError: string = "";

  const validatePassword = () => {
    const regex =
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    passwordError = regex.test(password)
      ? ""
      : "Falsches Passwort! Das Passwort muss mindestens einen Großbuchstaben, einen Kleinbuchstaben, eine Zahl, ein Sonderzeichen enthalten und mindestens 8 Zeichen lang sein.";
  };

  const validatePhonenumber = () => {
    const regex = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/;
    telephoneError = regex.test(editableUserInfo.PhoneNumber)
      ? ""
      : "False Telefonnummer! Die Telefonnummer muss in dem Format +1234567890 or 1234567890 sein.";
  };

  let editableUserInfo = {
    Email: userInfo.Email ?? "",
    FirstName: userInfo.FirstName ?? "",
    LastName: userInfo.LastName ?? "",
    JobTitle: userInfo.JobTitle ?? "",
    PhoneNumber: userInfo.PhoneNumber ?? "",
  };

  let password = "";

  const updateUser = () => {
    let aPromisses = [];
    if (password !== "") {
      aPromisses.push(
        fetch(`${API_URL}auth/reset-password`, {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ Password: password }),
        })
      );
    }

    aPromisses.push(
      fetch(`${API_URL}users`, {
        method: "PUT",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(editableUserInfo),
      })
    );

    Promise.all(aPromisses)
      .then((responses) => {
        if (responses.some((response) => !response.ok)) {
          throw new Error("Failed to update user information");
        }
        return Promise.all(responses.map((response) => response.json()));
      })
      .then((data) => {
        Swal.fire({
          title: "Success",
          text: "User information updated successfully",
          icon: "success",
          confirmButtonText: "Ok",
        });
      })
      .catch((error) => {
        Swal.fire({
          title: "Error",
          text: error.message,
          icon: "error",
          confirmButtonText: "Ok",
        });
      });
  };
</script>

<div class="container mx-auto p-6">
  <h1 class="text-3xl font-bold mb-8 text-center">User Dashboard</h1>

  <div class="mb-12">
    <h2 class="text-2xl font-semibold mb-4 text-center">User Information</h2>
    <div class="bg-gray-100 shadow-lg rounded-lg p-6 mb-8">
      <div class="flex items-center mb-6">
        {#if userInfo.ProfilePicture}
          <img
            src={userInfo.ProfilePicture ?? ""}
            alt="Profile of {userInfo.Username}"
            class="w-24 h-24 rounded-full mr-6 object-cover"
          />
        {:else}
          <PersonCircle
            class="w-20 h-20 rounded-full mr-6 object-cover text-black"
          />
        {/if}
        <div>
          <h3 class="text-xl font-medium">{userInfo.Username ?? "Username"}</h3>
          <p class="text-gray-600">{userInfo.UserTypeName ?? "User Type"}</p>
        </div>
      </div>
      <form on:submit|preventDefault={updateUser} class="space-y-4">
        <div>
          <label for="email" class="block text-gray-700">Email:</label>
          <input
            type="email"
            id="email"
            bind:value={editableUserInfo.Email}
            class="w-full px-4 py-2 border rounded-md"
          />
        </div>
        <div>
          <label for="firstName" class="block text-gray-700">First Name:</label>
          <input
            type="text"
            id="firstName"
            bind:value={editableUserInfo.FirstName}
            class="w-full px-4 py-2 border rounded-md"
          />
        </div>
        <div>
          <label for="lastName" class="block text-gray-700">Last Name:</label>
          <input
            type="text"
            id="lastName"
            bind:value={editableUserInfo.LastName}
            class="w-full px-4 py-2 border rounded-md"
          />
        </div>
        <div>
          <label for="jobTitle" class="block text-gray-700">Job Title:</label>
          <input
            type="text"
            id="jobTitle"
            bind:value={editableUserInfo.JobTitle}
            class="w-full px-4 py-2 border rounded-md"
          />
        </div>
        <div>
          <label for="phoneNumber" class="block text-gray-700"
            >Phone Number:</label
          >
          <input
            type="tel"
            id="phoneNumber"
            on:input={validatePhonenumber}
            bind:value={editableUserInfo.PhoneNumber}
            class="w-full px-4 py-2 border rounded-md"
          />
          <p class="text-red-500 text-sm mt-2">{telephoneError}</p>
        </div>
        <div>
          <label for="password" class="block text-gray-700">Password:</label>
          <input
            type="password"
            id="password"
            on:input={validatePassword}
            bind:value={password}
            class="w-full px-4 py-2 border rounded-md"
          />
          <p class="text-red-500 text-sm mt-2">{passwordError}</p>
        </div>
        <button
          type="submit"
          class="w-full bg-blue-600 text-white py-2 rounded-md"
          >Update Information</button
        >
      </form>
    </div>
  </div>
</div>

<style>
  .container {
    max-width: 800px;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
      "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji",
      "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
  }
</style>
