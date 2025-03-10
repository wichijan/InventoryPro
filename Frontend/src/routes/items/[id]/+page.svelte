<script lang="ts">
  import { API_URL } from "$lib/_services/ShelfService";
  import { goto } from "$app/navigation";
  import { getUser } from "$lib/_services/UserService";
  import Swal from "sweetalert2";
  import { onMount } from "svelte";
  import { browser } from "$app/environment";
  import { jsPDF } from "jspdf";

  import { Archive, Book, FilePdf } from "svelte-bootstrap-icons";

  export let data;

  let user: any;
  $: user = user;

  onMount(async () => {
    await getUser().then((res) => {
      user = res;
      if (item.UsersBorrowed) {
        userHasBorrowedItem = item.UsersBorrowed.some(
          (user) => user.BorrowedByUserID === res.ID
        );
      } else {
        userHasBorrowedItem = false;
      }
      userHasBorrowedItem = userHasBorrowedItem;
    });

    const mI = await getMoreInformation();
    item = { ...item, ...mI };
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

  let item: Item = data.item;
  let quickshelves = data.quickshelves;

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

  async function getMoreInformation() {
    let url = item.ItemTypes === "book" ? "book" : "set-of-objects";
    const response = await fetch(`${API_URL}items/${url}/${item.ID}`, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });
    let data = await response.json();
    return data;
  }
  async function makeReservation() {
    if (!user) {
      Swal.fire({
        title: "Du musst eingeloggt sein, um eine Reservierung vorzunehmen",
        icon: "error",
        confirmButtonText: "Ok",
      });
      return;
    }

    Swal.fire({
      title: "Für welches Datum möchtest du das Item reservieren?",
      html:
        '<input id="quantity" class="swal2-input" placeholder="Menge" type="number" min="1">' +
        '<input id="fromDate" class="swal2-input" placeholder="Start Datum" type="date">' +
        '<input id="toDate" class="swal2-input" placeholder="End Datum" type="date">',
      showCancelButton: true,
      confirmButtonText: "Reservieren",
      cancelButtonText: "Abbrechen",
      preConfirm: () => {
        const quantity = (
          document.getElementById("quantity") as HTMLInputElement
        ).value;
        const fromDate = (
          document.getElementById("fromDate") as HTMLInputElement
        ).value;
        const toDate = (document.getElementById("toDate") as HTMLInputElement)
          .value;

        if (!quantity || !fromDate || !toDate) {
          Swal.showValidationMessage("Alle Felder müssen ausgefüllt werden");
          return false;
        }

        if (parseInt(quantity) < 1) {
          Swal.showValidationMessage("Die Menge muss größer als 0 sein");
          return false;
        }
        if (new Date(fromDate) > new Date(toDate)) {
          Swal.showValidationMessage(
            "Das Startdatum muss vor dem Enddatum liegen"
          );
          return false;
        }
        if (new Date(fromDate) < new Date()) {
          Swal.showValidationMessage(
            "Das Startdatum muss in der Zukunft liegen"
          );
          return false;
        }
        if (parseInt(quantity) > item.QuantityInShelf) {
          Swal.showValidationMessage(
            "Die Menge ist größer als die verfügbare Menge"
          );
          return false;
        }

        return { quantity, fromDate, toDate };
      },
    }).then((result) => {
      if (result.isConfirmed) {
        const { quantity, fromDate, toDate } = result.value;
        fetch(`${API_URL}items/reserve`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
          body: JSON.stringify({
            itemID: item.ID,
            quantity: parseInt(quantity),
            timeFrom: fromDate,
            timeTo: toDate,
          }),
        })
          .then((res) => {
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
          })
          .catch((error) => {
            Swal.fire({
              title: "Fehler",
              text: "Ein Fehler ist aufgetreten",
              icon: "error",
              confirmButtonText: "Ok",
            });
          });
      }
    });
  }
  async function toPdf() {
    const doc = new jsPDF();
    doc.text("Item Information", 10, 10);
    doc.text("Name: " + item.Name, 10, 20);
    doc.text("Description: " + item.Description, 10, 30);
    doc.text("Quantity: " + item.QuantityInShelf, 10, 40);
    doc.text("HintText: " + item.HintText, 10, 50);
    doc.text("Damaged: " + item.Damaged, 10, 60);
    doc.text("ClassOne: " + item.ClassOne, 10, 70);
    doc.text("ClassTwo: " + item.ClassTwo, 10, 80);
    doc.text("ClassThree: " + item.ClassThree, 10, 90);
    doc.text("ClassFour: " + item.ClassFour, 10, 100);
    doc.text("DamagedDesc: " + item.DamagedDescription, 10, 110);
    doc.text("Keywords: " + item.Keywords, 10, 130);
    doc.text("Subject: " + item.Subject, 10, 140);
    doc.text("Reservations: " + item.Reservations, 10, 150);

    if (browser) {
      let output = doc.output("datauristring");
      let x = window.open();
      x.document.open();
      x.document.write(
        "<iframe src='" + output + "' width='100%' height='100%'></iframe>"
      );
      x.document.close();
    }
  }

  async function putItemInQuickShelf() {
    Swal.fire({
      title: "In welches Schnellregal soll das Item hinzugefügt werden?",
      html: `
          <!-- Select quickshelf -->
          <select id="quickshelf" class="swal2-select w-1/2">
            ${quickshelves.map(
              (qs) =>
                `<option value="${qs.QuickShelfID}">${getWarehouseAndRoomName(qs).warehouseName} - ${getWarehouseAndRoomName(qs).roomName} - ${qs.Name}</option>`
            )}
          </select>
        `,
      showCancelButton: true,
      confirmButtonText: "Hinzufügen",
      cancelButtonText: "Abbrechen",
      preConfirm: () => {
        const quickshelfID = document.getElementById("quickshelf").value;

        if (!quickshelfID) {
          Swal.showValidationMessage("Alle Felder müssen ausgefüllt werden");
          return false;
        }

        return { quickshelfID };
      },
    }).then((result) => {
      if (result.isConfirmed) {
        const { quickshelfID } = result.value;
        fetch(`${API_URL}items/add-item-to-quick-shelf`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
          body: JSON.stringify({
            QuickShelfId: quickshelfID,
            ItemID: item.ID,
            Quantity: Number(1),
          }),
        })
          .then((res) => {
            if (res.ok) {
              Swal.fire({
                title: "Hinzugefügt",
                text: "Das Item wurde erfolgreich dem Schnellregal hinzugefügt",
                icon: "success",
                confirmButtonText: "Ok",
              });
              setTimeout(() => {
                browser ? window.location.reload() : null;
              }, 2000);
            } else {
              Swal.fire({
                title: "Fehler",
                text: "Das Item konnte nicht hinzugefügt werden",
                icon: "error",
                confirmButtonText: "Ok",
              });
            }
          })
          .catch((error) => {
            Swal.fire({
              title: "Fehler",
              text: "Ein Fehler ist aufgetreten",
              icon: "error",
              confirmButtonText: "Ok",
            });
          });
      }
    });
  }
  function getWarehouseAndRoomName(quickshelf) {
    let warehouses = data.warehouses;
    for (let warehouse of warehouses) {
      if (warehouse.Rooms) {
        for (let room of warehouse.Rooms) {
          if (room.ID === quickshelf.RoomID) {
            return {
              warehouseName: warehouse.Name,
              roomName: room.Name,
            };
          }
        }
      }
    }
    return null;
  }
</script>

{#if item}
  <div class="flex flex-col items-center w-full my-10">
    <div class="bg-gray-50 p-8 rounded-lg shadow-lg w-full max-w-4xl">
      <h2 class="flex text-3xl font-semibold text-gray-800 mb-4">
        Generelle Information
        <button on:click={toPdf} class="ml-4">
          <FilePdf class="w-8 h-8 my-auto" />
        </button>
      </h2>
      <div class="flex flex-row justify-between">
        <div class="flex-shrink-0 w-1/3">
          {#if item.Picture}
            <img
              src="{API_URL}items-picture/{item.ID}"
              alt="{item.Name} Bild"
              class="rounded-lg shadow-md w-full object-cover"
            />
          {:else if item.ItemTypes === "book"}
            <Book class=" text-gray-400 w-full h-full px-5 py-5" />
          {:else}
            <Archive class=" text-gray-400 w-full h-full px-5 py-5" />
          {/if}
        </div>
        <div class="flex-grow ml-10">
          <div class="text-gray-700 text-lg mb-4">Name: {item.Name}</div>
          <div class="text-gray-700 text-lg mb-4">
            Beschreibung: {item.Description}
          </div>
          <div class="text-gray-700 text-lg mb-4">
            Anzahl: {item.QuantityInShelf}
          </div>
          {#if item.UsersBorrowed && item.UsersBorrowed.length > 0}
            <div class="text-gray-700 text-lg mb-4">
              Ausgeliehen von: {#each item.UsersBorrowed as user, index}
                {user.BorrowedByUserName}
                {#if index !== item.UsersBorrowed.length - 1},
                {/if}
              {/each}
            </div>
          {/if}
          {#if item.Subject}
            <div class="text-gray-700 text-lg mb-4">
              Fächer: {#each item.Subject as subject}
                {subject.Name}
              {/each}
            </div>
          {/if}
          {#if item.Keywords}
            <div class="text-gray-700 text-lg mb-4">
              Keywords: {#each item.Keywords as keyword}
                {keyword.Keyword}
              {/each}
            </div>
          {/if}
          {#if item.HintText}
            <div class="text-gray-700 text-lg mb-4">
              Hinweis: {item.HintText}
            </div>
          {/if}
          <div
            class="flex flex-col space-y-5 sm:space-y-0 sm:space-x-4 mt-6 sm:flex-row"
          >
            <button
              class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-md transition"
              on:click={borrow}
            >
              Ausleihen
            </button>
            {#if user && userHasBorrowedItem}
              <button
                class="bg-red-600 hover:bg-red-700 text-white font-semibold py-2 px-4 rounded-md transition"
                on:click={returnItem}
              >
                Zurückgeben
              </button>
              <button
                class="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-4 rounded-md transition"
                on:click={() => {
                  let allUsers = data.users;
                  if (!allUsers) {
                    Swal.fire({
                      title: "Fehler",
                      text: "Es gibt keine User",
                      icon: "error",
                      confirmButtonText: "Ok",
                    });
                    return;
                  }
                  let usersToSend = allUsers.filter((u) => u.ID !== user.ID);
                  if (usersToSend.length === 0) {
                    Swal.fire({
                      title: "Fehler",
                      text: "Es gibt keine anderen User",
                      icon: "error",
                      confirmButtonText: "Ok",
                    });
                    return;
                  }
                  Swal.fire({
                    title: "An anderen User senden",
                    html: `
                      <!-- Select user which isnt him-->
                      <select id="username" class="swal2-select w-1/2">
                        ${allUsers
                          .filter((u) => u.ID !== user.ID)
                          .map(
                            (u) =>
                              `<option value="${u.ID}">${u.Username}</option>`
                          )}
                      </select>
                    `,
                    showCancelButton: true,
                    confirmButtonText: "Senden",
                    cancelButtonText: "Abbrechen",
                    preConfirm: () => {
                      const userID = document.getElementById("username").value;

                      if (!userID) {
                        Swal.showValidationMessage(
                          "Alle Felder müssen ausgefüllt werden"
                        );
                        return false;
                      }

                      return { userID };
                    },
                  }).then((result) => {
                    if (result.isConfirmed) {
                      const { userID } = result.value;
                      fetch(`${API_URL}items/transfer-request`, {
                        method: "POST",
                        headers: {
                          "Content-Type": "application/json",
                        },
                        credentials: "include",
                        body: JSON.stringify({
                          ItemID: item.ID,
                          NewUserID: userID,
                        }),
                      })
                        .then((res) => {
                          if (res.ok) {
                            Swal.fire({
                              title: "Gesendet",
                              text: "Das Item wurde erfolgreich gesendet",
                              icon: "success",
                              confirmButtonText: "Ok",
                            });
                            setTimeout(() => {
                              browser ? window.location.reload() : null;
                            }, 2000);
                          } else {
                            Swal.fire({
                              title: "Fehler",
                              text: "Das Item konnte nicht gesendet werden",
                              icon: "error",
                              confirmButtonText: "Ok",
                            });
                          }
                        })
                        .catch((error) => {
                          Swal.fire({
                            title: "Fehler",
                            text: "Ein Fehler ist aufgetreten",
                            icon: "error",
                            confirmButtonText: "Ok",
                          });
                        });
                    }
                  });
                }}
              >
                Senden an User
              </button>
              <button
                class="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-4 rounded-md transition"
                on:click={() => {
                  putItemInQuickShelf();
                }}
              >
                Ins Schnellregal
              </button>
            {/if}
          </div>
        </div>
      </div>
    </div>
    {#if item.ItemTypes === "book"}
      <div class="bg-gray-50 p-8 rounded-lg shadow-lg w-full max-w-4xl mt-10">
        <h3 class="text-2xl font-semibold text-gray-800 mb-4">
          Buch Information
        </h3>
        <div class="text-gray-700 text-lg mb-4">
          Herausgeber: {item.Publisher}
        </div>
        <div class="text-gray-700 text-lg mb-4">Autor: {item.Author}</div>
        <div class="text-gray-700 text-lg mb-4">ISBN: {item.Isbn}</div>
        <div class="text-gray-700 text-lg mb-4">Ausgabe: {item.Edition}</div>
      </div>
    {/if}
    {#if item.ItemTypes === "set_of_objects"}
      <div class="bg-gray-50 p-8 rounded-lg shadow-lg w-full max-w-4xl mt-10">
        <h3 class="text-2xl font-semibold text-gray-800 mb-4">
          Normale Information
        </h3>
        <div class="text-gray-700 text-lg mb-4">
          Totale Objecte: {item.TotalObjects}
        </div>
        <div class="text-gray-700 text-lg mb-4">
          Nutzbare Objecte: {item.UsefulObjects}
        </div>
        <div class="text-gray-700 text-lg mb-4">
          Kaputte Objecte: {item.BrokenObjects}
        </div>
        <div class="text-gray-700 text-lg mb-4">
          Verlorene Objecte: {item.LostObjects}
        </div>
      </div>
    {/if}

    <div class="flex justify-between w-full max-w-4xl mt-10">
      <div class="bg-gray-50 p-8 rounded-lg shadow-lg w-1/2 mr-5">
        <h3 class="text-2xl font-semibold text-gray-800 mb-4">Kaputt</h3>
        <div class="text-gray-700 text-lg">{item.Damaged ? "Ja" : "Nein"}</div>
        {#if item.Damaged}
          <div class="text-gray-700 text-lg mt-2">
            {item.DamagedDescription}
          </div>
        {/if}
      </div>
      <div class="bg-gray-50 p-8 rounded-lg shadow-lg w-1/2 ml-5">
        <h3 class="text-2xl font-semibold text-gray-800 mb-4">Klassen</h3>
        <div class="flex flex-col text-gray-700 text-lg">
          <div>{item.ClassOne ? "Klasse 1" : ""}</div>
          <div>{item.ClassTwo ? "Klasse 2" : ""}</div>
          <div>{item.ClassThree ? "Klasse 3" : ""}</div>
          <div>{item.ClassFour ? "Klasse 4" : ""}</div>
        </div>
      </div>
    </div>
    <div class="flex flex-col w-full max-w-4xl mt-10" id="reservations">
      <div class="bg-gray-50 p-8 rounded-lg shadow-lg w-full mb-10">
        <h3 class="text-2xl font-semibold text-gray-800 mb-4">
          Reservierungen
        </h3>
        <p class="text-gray-700 text-lg mb-4">
          Wenn der Artikel nicht verfügbar ist, können Sie eine Reservierung
          vornehmen.
        </p>
        <div class="flex space-x-4">
          <button
            class="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-4 rounded-md transition"
            on:click={makeReservation}
          >
            Item reservieren
          </button>
        </div>
        {#if errorText}
          <div class="text-red-500 text-lg mt-4">{errorText}</div>
        {/if}
      </div>

      {#if item.Reservations && item.Reservations.length > 0}
        <div class="bg-gray-50 p-8 rounded-lg shadow-lg w-full">
          <h3 class="text-2xl font-semibold text-gray-800 mb-4">
            Bestehende Reservierungen
          </h3>
          <ul class="list-disc list-inside">
            {#each item.Reservations as reservation}
              <li class="text-gray-700 text-lg mb-2">
                <strong>Username:</strong>
                {reservation.Username},
                <strong>Anzahl:</strong>
                {reservation.Quantity},
                <strong>Von:</strong>
                {new Date(reservation.TimeFrom).toLocaleDateString()},
                <strong>Bis:</strong>
                {new Date(reservation.TimeTo).toLocaleDateString()}
                {#if user && reservation.UserID === user.ID}
                  <button
                    class="bg-red-600 hover:bg-red-700 text-white font-semibold py-2 px-4 rounded-md transition ml-4"
                    on:click={() => {
                      fetch(
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
                            title: "Reservierung gelöscht",
                            text: "Die Reservierung wurde erfolgreich gelöscht",
                            icon: "success",
                            confirmButtonText: "Ok",
                          });
                          setTimeout(() => {
                            browser ? window.location.reload() : null;
                          }, 2000);
                        } else {
                          Swal.fire({
                            title: "Fehler",
                            text: "Die Reservierung konnte nicht gelöscht werden",
                            icon: "error",
                            confirmButtonText: "Ok",
                          });
                        }
                      });
                    }}
                  >
                    Löschen
                  </button>
                {/if}
              </li>
            {/each}
          </ul>
        </div>
      {/if}
    </div>
  </div>
{/if}
