<script lang="ts">
  import CookieConsent from "$lib/_layout/CookieConsent.svelte";
  import Navbar from "$lib/_layout/SideBar/Navbar.svelte";
  import "../app.css";

  import { page } from "$app/stores";
  import { afterNavigate, beforeNavigate, goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { API_URL } from "$lib/_services/ShelfService";
  import Footer from "$lib/_layout/Footer.svelte";
  import { isUserLoggedIn } from "$lib/_services/UserService";

  let url: string;

  $: url = $page.url.pathname;

  let breadcrumbs: any[] = [];
  $: breadcrumbs = breadcrumbs;

  afterNavigate(async () => {
    url = $page.url.pathname;
    await setBreadcrumbs().then((data) => {
      breadcrumbs = data;
    });
  });

  onMount(async () => {
    await setBreadcrumbs().then((data) => {
      breadcrumbs = data;
    });
    checkLog();
  });

  afterNavigate(async (event) => {
    checkLog();
  });

  async function checkLog() {
    const logIn = await isUserLoggedIn();
    if (
      !logIn &&
      url !== "/auth/login" &&
      url !== "/auth/register" &&
      url !== "/auth/code"
    ) {
      goto("/auth/login");
    }
  }

  async function setBreadcrumbs(): Promise<any> {
    return new Promise(async (resolve, reject) => {
      const breadcrumbs = [];
      const urlArray = url.split("/");

      const fetchData = async (type, id) => {
        const response = await fetch(`${API_URL}${type}/${id}`, {
          method: "GET",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
        });

        if (!response.ok) {
          throw new Error(`Failed to fetch ${type} with id ${id}`);
        }

        const data = await response.json();
        return data.Name;
      };

      for (let i = 0; i < urlArray.length; i++) {
        const part = urlArray[i];

        if (part === "auth") continue;

        if (part === "login" || part === "register") {
          breadcrumbs.push({
            text: capitalize(part),
            href: `/auth/${part}`,
          });
          continue;
        }

        if (isUUID(part)) {
          const partBefore = urlArray[i - 1];

          if (partBefore === "detail") {
            try {
              const name = await fetchData("items", part);
              if (name && breadcrumbs[breadcrumbs.length - 1]?.text !== name) {
                breadcrumbs.push({
                  text: name,
                  href: `/items/${part}`,
                });
              }
            } catch (error) {
              console.error("Error fetching item details:", error);
            }
          } else if (partBefore === "rooms") {
            try {
              const name = await fetchData("rooms", part);
              if (name && breadcrumbs[breadcrumbs.length - 1]?.text !== name) {
                breadcrumbs.push({
                  text: name,
                  href: `/overview/rooms/${part}`,
                });
              }
            } catch (error) {
              console.error("Error fetching room details:", error);
            }
          }
          continue;
        }

        if (part === "rooms") {
          breadcrumbs.push({
            text: "Rooms",
            href: "/overview/rooms",
          });
          continue;
        }

        if (part !== "") {
          breadcrumbs.push({
            text: capitalize(part),
            href: `/${part}`,
          });
        } else {
          breadcrumbs.push({ text: "Home", href: "/" });
        }
      }

      resolve(breadcrumbs);
    });
  }

  function capitalize(str: string) {
    return str.charAt(0).toUpperCase() + str.slice(1);
  }

  function isUUID(str: string) {
    const uuidRegex =
      /^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$/;
    return uuidRegex.test(str);
  }
</script>

<div
  class="relative flex flex-col justify-between min-w-screen min-h-screen bg-gradient-to-br from-gray-100 to-gray-200"
>
  <div class="z-50">
    <Navbar />
  </div>
  <div
    class=" mt-3 bg-tertiary px-5 py-1 rounded-md ml-[5.25rem] mr-5"
    id="breadcrumbs"
  >
    {#key breadcrumbs}
      {#each breadcrumbs as breadcrumb, index}
        <a
          href={breadcrumb.href}
          class="text-[#344e41] hover:text-blue-500 duration-300"
          >{breadcrumb.text}</a
        >
        {index < breadcrumbs.length - 1 ? "/ " : " "}
      {/each}
    {/key}
  </div>
  <div class="flex-grow flex-1 overflow-hidden pl-12">
    <slot />
  </div>
  <div class="flex min-w-screen h-full"><Footer /></div>
</div>

<CookieConsent />
