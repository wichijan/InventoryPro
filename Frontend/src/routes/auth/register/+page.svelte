<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import { ArrowRight } from "svelte-bootstrap-icons";
  import Swal from "sweetalert2";
  import { fly } from "svelte/transition";

  export let data;

  /*
     "email": "string",
  "firstname": "string",
  "jobtitle": "string",
  "lastname": "string",
  "password": "string",
  "phonenumber": "string",
  "username": "string",
  "usertypename": "string"
   */

  let email: string = "";
  let password: string = "";
  let username: string = "";
  let firstname: string = "";
  let lastname: string = "";
  let jobtitle: string = "";
  let phonenumber: string = "";
  let usertypename: string = "";
  let passwordError: string = "";

  const validatePassword = () => {
    const regex =
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    passwordError = regex.test(password)
      ? ""
      : "Invalid password! Password should contain at least one uppercase letter, one lowercase letter, one number, one special character and should be at least 8 characters long.";
  };

  let step = 0;

  function preRegister(e: any) {
    e.preventDefault();

    fetch(API_URL + "auth/check-email", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
      }),
    }).then((response) => {
      if (response.ok) {
        register();
      } else {
        Swal.fire({
          position: "top-end",
          icon: "error",
          title: "E-Mail existiert bereits!",
          showConfirmButton: false,
          timer: 1500,
        });
      }
    });
  }

  function register() {
    fetch(API_URL + "auth/register", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
        firstname: firstname,
        jobtitle: jobtitle,
        lastname: lastname,
        password: password,
        phonenumber: phonenumber,
        username: username,
        usertypename: usertypename,
      }),
    }).then((response) => {
      if (response.ok) {
        Swal.fire({
          position: "top-end",
          icon: "info",
          title:
            "Deine Registrierung war erfolgreich! Ein Admin muss dich jetzt freischalten.",
          showConfirmButton: false,
          timer: 1500,
        });
      } else {
        Swal.fire({
          position: "top-end",
          icon: "error",
          title: "Registration failed!",
          showConfirmButton: false,
          timer: 1500,
        });
      }
    });
  }

  function firstCheck(e: any) {
    e.preventDefault();

    fetch(API_URL + "auth/check-username", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username,
      }),
    }).then((response) => {
      if (response.ok) {
        step = 1;
      } else {
        Swal.fire({
          position: "top-end",
          icon: "error",
          title: "Benutzername existiert bereits!",
          showConfirmButton: false,
          timer: 1500,
        });
      }
    });
  }
</script>

<div class="min-h-screen flex justify-center items-center">
  <div class="flex space-x-5 justify-between w-full" class:mr-5={step === 1}>
    <div
      class="max-w-xs w-full m-auto bg-tertiary rounded-xl shadow-md overflow-hidden md:max-w-2xl px-4 py-6 hover:shadow-lg duration-300"
    >
      <h2
        class="text-center text-4xl font-extrabold text-white tracking-widest hover:text-blue-300 mb-4 duration-300"
      >
        Register
      </h2>
      <form class="mt-8 space-y-6" on:submit={firstCheck}>
        <div>
          <label for="email" class="text-white text-lg font-bold">
            Benutzername
          </label>
          <div class="mt-1 rounded-md shadow-sm">
            <input
              bind:value={username}
              type="text"
              class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
              placeholder="Benutzername"
              required
            />
          </div>
        </div>
        <div>
          <label for="password" class="text-white text-lg font-bold">
            Password
          </label>
          <div class="mt-1 rounded-md shadow-sm">
            <input
              bind:value={password}
              on:input={validatePassword}
              type="password"
              class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
              placeholder="Password"
              required
            />
            <p class="text-red-500 text-sm mt-2">{passwordError}</p>
          </div>
        </div>
        <div class="flex justify-between">
          <div class="flex flex-row w-max gap-1">
            <input type="checkbox" class="mx-auto my-auto rounded" />
            <label
              for="remember"
              class="text-gray-700 text-sm font-bold my-auto"
              >Angemeldet bleiben?</label
            >
          </div>
          <a
            href="/auth/code"
            class="text-gray-700 text-sm font-bold hover:text-blue-400 duration-300"
            >Anmelden mit code</a
          >
          <a
            href="/auth/login"
            class="text-gray-700 text-sm font-bold hover:text-blue-400 duration-300"
            >Bereits ein Account?</a
          >
        </div>

        <button
          disabled={passwordError}
          type="submit"
          class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full"
          hidden={step === 1}
        >
          Next Step <ArrowRight
            class="inline-block w-5 h-5 animate-pulse my-auto"
          />
        </button>
      </form>
    </div>
    {#key step}
      <div
        class="max-w-xs w-full m-auto bg-tertiary rounded-xl shadow-md overflow-hidden md:max-w-2xl px-4 py-6 my-5 hover:shadow-lg duration-300"
        class:hidden={step !== 1}
        in:fly={{ x: 100, duration: 200, delay: 201 }}
        out:fly={{ x: -100, duration: 200 }}
      >
        <h2
          class="text-center text-4xl font-extrabold text-white tracking-widest hover:text-blue-300 mb-4 duration-300"
        >
          Zusätzliche Informationen
        </h2>
        <form class="mt-8 space-y-4" on:submit={preRegister}>
          <div>
            <label for="username" class="text-white text-lg font-bold">
              E-Mail
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                bind:value={email}
                type="email"
                class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
                placeholder="E-Mail"
                required
              />
            </div>
          </div>
          <div>
            <label for="firstname" class="text-white text-lg font-bold">
              Vorname
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                bind:value={firstname}
                type="text"
                class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
                placeholder="Vorname"
                required
              />
            </div>
          </div>
          <div>
            <label for="lastname" class="text-white text-lg font-bold">
              Nachname
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                bind:value={lastname}
                type="text"
                class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
                placeholder="Nachname"
                required
              />
            </div>
          </div>
          <div>
            <label for="phonenumber" class="text-white text-lg font-bold">
              Telefonnummer
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                bind:value={phonenumber}
                type="text"
                class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
                placeholder="Telefonnummer"
                required
              />
            </div>
          </div>
          <div>
            <label for="jobtitle" class="text-white text-lg font-bold">
              Jobtitel
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                bind:value={jobtitle}
                type="text"
                class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
                placeholder="Jobtitel"
                required
              />
            </div>
          </div>
          <div>
            <label for="usertypename" class="text-white text-lg font-bold">
              Usertyp
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <!-- <input
                bind:value={usertypename}
                type="text"
                class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
                placeholder="Usertyp"
                required
              /> -->
              <!--Dropdown menu for user-types-->
              <select
                bind:value={usertypename}
                class="block w-full p-3 border-gray-300 rounded-md placeholder-gray-500 text-gray-900 focus:border-blue-500"
                required
              >
                {#await data.userTypes then types}
                  {#each types as type}
                    <option value={type.TypeName}>{type.TypeName}</option>
                  {/each}
                {/await}
              </select>
            </div>
          </div>

          <button
            disabled={passwordError}
            type="submit"
            class="bg-[#d5bdaf] hover:bg-d6ccc2 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 mt-5 w-full"
          >
            Register
          </button>
        </form>
      </div>
    {/key}
  </div>
</div>
