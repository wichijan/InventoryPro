<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import { goto } from "$app/navigation";
  import { getUser } from "$lib/_services/UserService";
  import Spinner from "$lib/templates/Spinner.svelte";
  import { fade } from "svelte/transition";

  import Swal from "sweetalert2";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";

  export let data;

  let user: any;
  $: user = user;

  onMount(async () => {
    await getUser().then((res) => {
      user = res;
      userHasBorrowedItem = item.UsersBorrowed
        ? item.UsersBorrowed.find(
            (user_) => user_.BorrowedByUserID === user.ID
          ) !== null
        : false;
      userHasBorrowedItem = userHasBorrowedItem;
    });
  });

  type Item = {
    ID: string;
    ItemTypes: string;
    QuantityInShelf: number;
    Name: string;
    Description: string;
    RegularShelfID: string;
    ClassOne: boolean;
    ClassTwo: boolean;
    ClassThree: boolean;
    ClassFour: boolean;
    Damaged: boolean;
    DamagedDesc: string;
    Picture: string;
    HintText: string;
    Quantity: number;
    UsersBorrowed: any[];
    Keywords: any[];
    Subject: any[];
    Reservations: any[];
  };

  const item: Item = data.item;

  let errorText = "";
  $: errorText = errorText;

  function checkIfItemIsReserved(): boolean {
    let rValue = false;
    if (item.Reservations) {
      item.Reservations.forEach((reservation) => {
        const curentDate = new Date();
        const timeFrom = new Date(reservation.TimeFrom);
        const timeTo = new Date(reservation.TimeTo);
        if (
          curentDate.getTime() >= timeFrom.getTime() &&
          curentDate.getTime() <= timeTo.getTime()
        ) {
          console.log(Math.abs(reservation.Quantity - item.QuantityInShelf));
          rValue = Math.abs(reservation.Quantity - item.QuantityInShelf) === 0;
        } else {
          rValue = false;
        }
      });
    } else {
      rValue = false;
    }
    return rValue;
  }

  let userHasBorrowedItem = false;
  $: userHasBorrowedItem = userHasBorrowedItem;

  async function borrow() {
    if (!user) {
      Swal.fire({
        title: "Du musst eingeloggt sein, um ein Item auszuleihen",
        icon: "error",
        confirmButtonText: "Ok",
      });
      return;
    }
    if (userHasBorrowedItem) {
      Swal.fire({
        title: "Du hast das Item bereits ausgeliehen",
        icon: "error",
        confirmButtonText: "Ok",
      });
      return;
    }
    if (item.QuantityInShelf === 0) {
      errorText =
        "Es sind keine Items mehr verfügbar, bitte reservieren Sie das Item";
      goto(`#reservations`);
      return;
    }

    if (checkIfItemIsReserved()) {
      errorText =
        "Das Item ist bereits reserviert und kann nicht ausgeliehen werden. Bitte reservieren Sie das Item";
      goto(`#reservations`);
      return;
    }

    Swal.fire({
      title: "Wie viele Items möchtest du ausleihen?",
      input: "number",
      showCancelButton: true,
      confirmButtonText: "Ausleihen",
      cancelButtonText: "Abbrechen",
      showLoaderOnConfirm: true,
      inputValidator: (value) => {
        if (!value) {
          return "Du musst eine Anzahl eingeben";
        }
        if (Number(value) > item.QuantityInShelf) {
          return (
            "Es sind nicht genügend Items vorhanden, maximal verfügbar: " +
            item.QuantityInShelf +
            " Items"
          );
        }
        if (Number(value) <= 0) {
          return "Die Anzahl muss größer als 0 sein";
        }
      },
      preConfirm: (quantity) => {
        return fetch(`${API_URL}items/borrow`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
          body: JSON.stringify({
            ItemID: item.ID,
            Quantity: Number(quantity),
          }),
        }).then((res) => {
          if (res.ok) {
            Swal.fire({
              title: "Ausgeliehen",
              text: "Das Item wurde erfolgreich ausgeliehen",
              icon: "success",
              confirmButtonText: "Ok",
            });
            setTimeout(() => {
              browser ? window.location.reload() : null;
            }, 2000);
          } else {
            Swal.fire({
              title: "Fehler",
              text: "Das Item konnte nicht ausgeliehen werden",
              icon: "error",
              confirmButtonText: "Ok",
            });
          }
        });
      },
    });
  }
  async function returnItem() {
    Swal.fire({
      title: "Möchtest du das Item wirklich zurückgeben?",
      showCancelButton: true,
      confirmButtonText: "Ja",
      cancelButtonText: "Nein",
      showLoaderOnConfirm: true,
      preConfirm: () => {
        return fetch(`${API_URL}items/return/${item.ID}`, {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        }).then((res) => {
          if (res.ok) {
            Swal.fire({
              title: "Zurückgegeben",
              text: "Das Item wurde erfolgreich zurückgegeben",
              icon: "success",
              confirmButtonText: "Ok",
            });
            setTimeout(() => {
              browser ? window.location.reload() : null;
            }, 2000);
          } else {
            Swal.fire({
              title: "Fehler",
              text: "Das Item konnte nicht zurückgegeben werden",
              icon: "error",
              confirmButtonText: "Ok",
            });
          }
        });
      },
    });
  }
</script>

{#if item}
  <div class="my-5 flex flex-col w-full">
    <div
      class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 rounded-md"
      id="generalInformation"
    >
      <div class="mx-auto font-semibold text-2xl" id="header">
        General information
      </div>
      <div class="flex flex-row justify-between mt-3">
        <div class="ml-5 w-1/2" id="imgItem">
          <img
            src="https://via.placeholder.com/250"
            alt="item"
            class="rounded-md w-fit h-fit mx-auto object-cover"
          />
        </div>
        <div class="flex flex-col text-[#344e41] ml-12 font-semibold">
          <div class="">Name: {item.Name}</div>
          <div class="">Beschreibung: {item.Description}</div>
          <div class="">Menge: {item.QuantityInShelf}</div>
          {#if userHasBorrowedItem}
            <div class="">
              Ausgeliehen von: {#each item.UsersBorrowed as user}
                {user.BorrowedByUserName},
              {/each}
            </div>
          {/if}
          {#if item.Subject}
            <div class="">Fächer</div>
            {#each item.Subject as subject}
              -
              <div class="text-[#344e41]">{subject.Name}</div>
            {/each}
          {/if}

          <hr class="my-5" />
          <div class="flex justify-between w-full space-x-5">
            <button
              class="bg-[#d5bdaf] hover:bg-green-500 hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 w-full"
              on:click={() => {
                borrow();
              }}
            >
              Ausleihen
            </button>
            {#if user}
              {#if userHasBorrowedItem}
                <button
                  class="bg-[#d5bdaf] hover:bg-red-500 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 w-full"
                  on:click={() => {
                    returnItem();
                  }}
                >
                  Zurückgeben
                </button>
              {/if}
            {/if}
          </div>
        </div>
      </div>
    </div>
    <div class="flex justify-between">
      <div
        class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 w-full rounded-md"
        id="damaged"
      >
        <div class="mx-auto font-semibold text-xl" id="header">Beschädigt</div>
        <div class="w-full">
          <div class="flex space-x-3">
            {#if item.Damaged}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6 rounded-md bg-red-400 text-white"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M6 18 18 6M6 6l12 12"
                />
              </svg>
              <div class="text-red-600">Beschädigt</div>
            {:else}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6 rounded-md bg-green-400 text-white"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <div class="text-green-600">In Ordnung</div>
            {/if}
          </div>
          {#if item.DamagedDesc}
            <br />
            <hr />
            <div class="mt-5">Beschreibung: <br /> {item.DamagedDesc}</div>
          {/if}
        </div>
      </div>
      <div
        class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 w-full rounded-md"
        id="classes"
      >
        <div class="mx-auto font-semibold text-xl" id="header">Klassen</div>
        <div class="w-full">
          <div class="flex-row">
            <div class={item.ClassOne ? "text-green-500" : "text-red-500"}>
              Klasse 1
            </div>
            <div class={item.ClassTwo ? "text-green-500" : "text-red-500"}>
              Klasse 2
            </div>
            <div class={item.ClassThree ? "text-green-500" : "text-red-500"}>
              Klasse 3
            </div>
            <div class={item.ClassFour ? "text-green-500" : "text-red-500"}>
              Klasse 4
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="flex flex-col bg-tertiary px-5 py-5 mt-5 ml-10 mr-5 rounded-md"
      id="reservations"
    >
      <div class="mx-auto font-semibold text-xl" id="header">
        Reservierungen
      </div>
      <div class="w-full">
        <div class="flex flex-col">
          {#key errorText}
            <div
              class="flex space-x-1 w-1/2 mx-auto mb-5"
              class:hidden={errorText === ""}
              transition:fade={{ delay: 0, duration: 300 }}
            >
              <button
                id="errorText"
                class="w-max text-red-500 hover:bg-white duration-300 flex space-x-2 px-2 py-2 rounded-md bg-red-100"
                on:click={() => {
                  errorText = "";
                }}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                  class="w-6 h-6 my-auto text-red-500"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z"
                  />
                </svg>{errorText}</button
              >
            </div>
          {/key}
          {#if !item.Reservations}
            <div class="text-[#344e41]">Keine Reservierungen</div>
          {:else}
            <div class="grid grid-cols-2 w-full gap-5">
              {#each item.Reservations as reservation}
                <div
                  class="flex flex-col bg-white py-3 px-3 rounded-md w-max mx-auto space-y-2"
                >
                  <div class="text-[#344e41] mx-5">
                    Reserviert von: {reservation.UserID}
                  </div>
                  <div class="text-[#344e41] mx-5">
                    Anzahl: {reservation.Quantity}
                  </div>
                  <div class="text-[#344e41] mx-5">
                    Von: {new Date(reservation.TimeFrom).toLocaleDateString()} Bis:
                    {new Date(reservation.TimeTo).toLocaleDateString()}
                  </div>

                  {#if user}
                    {#if user.ID === reservation.UserID}
                      <button
                        class="bg-[#d5bdaf] hover:bg-red-500 hover:text-black hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 w-full"
                        on:click={() => {
                          Swal.fire({
                            title:
                              "Möchtest du die Reservierung wirklich stornieren?",
                            showCancelButton: true,
                            confirmButtonText: "Ja",
                            cancelButtonText: "Nein",
                            showLoaderOnConfirm: true,
                            preConfirm: () => {
                              return fetch(
                                `${API_URL}items/reserve-cancel/${reservation.ReservationID}`,
                                {
                                  method: "DELETE",
                                  headers: {
                                    "Content-Type": "application/json",
                                  },
                                  credentials: "include",
                                }
                              ).then((res) => {
                                if (res.ok) {
                                  Swal.fire({
                                    title: "Storniert",
                                    text: "Die Reservierung wurde erfolgreich storniert",
                                    icon: "success",
                                    confirmButtonText: "Ok",
                                  });
                                  setTimeout(() => {
                                    browser ? window.location.reload() : null;
                                  }, 2000);
                                } else {
                                  Swal.fire({
                                    title: "Fehler",
                                    text: "Die Reservierung konnte nicht storniert werden",
                                    icon: "error",
                                    confirmButtonText: "Ok",
                                  });
                                }
                              });
                            },
                          });
                        }}
                      >
                        Stornieren
                      </button>
                    {/if}
                  {/if}
                </div>
                <hr class="mt-5" />
              {/each}
            </div>
          {/if}
          <form
            class="flex flex-col mt-5"
            on:submit|preventDefault={(event) => {
              fetch(`${API_URL}items/reserve`, {
                method: "POST",
                headers: {
                  "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({
                  ItemID: item.ID,
                  TimeFrom: event.target.timeFrom.value,
                  TimeTo: event.target.timeTo.value,
                  Quantity: Number(event.target.quantity.value),
                }),
              }).then((res) => {
                if (res.ok) {
                  Swal.fire({
                    title: "Reserviert",
                    text: "Das Item wurde erfolgreich reserviert",
                    icon: "success",
                    confirmButtonText: "Ok",
                  });
                  setTimeout(() => {
                    browser ? window.location.reload() : null;
                  }, 2000);
                } else {
                  Swal.fire({
                    title: "Fehler",
                    text: "Das Item konnte nicht reserviert werden",
                    icon: "error",
                    confirmButtonText: "Ok",
                  });
                }
              });
            }}
          >
            <div class="flex flex-col">
              <label for="quantity">Anzahl</label>
              <input
                type="number"
                id="quantity"
                class="rounded-md border-2 border-gray-300 focus:border-green-500"
                name="quantity"
                min="1"
                max={item.QuantityInShelf}
                required
              />
            </div>
            <div class="flex flex-col">
              <label for="timeFrom">Von</label>
              <input
                type="date"
                class="rounded-md border-2 border-gray-300 focus:border-green-500"
                id="timeFrom"
                name="timeFrom"
                required
              />
            </div>
            <div class="flex flex-col">
              <label for="timeTo">Bis</label>
              <input
                type="date"
                id="timeTo"
                name="timeTo"
                required
                class="rounded-md border-2 border-gray-300 focus:border-green-500"
              />
            </div>
            <button
              class="bg-[#d5bdaf] hover:bg-green-500 hover:shadow-lg duration-500 text-white rounded-md px-3 py-1 w-full mt-5"
              type="submit"
            >
              Reservieren
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
{:else}
  <Spinner />
{/if}
