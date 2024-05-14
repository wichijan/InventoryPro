import { API_URL } from '$lib/_services/ShelfService.js';

export const load = async ({ fetch, params }) => {

    async  function getAllShelvesWithItems() {
        return new Promise((resolve, reject) => {
            fetch(API_URL + 'shelveswithitems', {
                method: 'GET',
                credentials: 'include',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(async (response) => {
                if (response.ok) {
                    await response.json().then((data) => {
                        resolve(data);
                    });
                } else {
                    reject(response.statusText);
                }
            });
        });
    }

    return { shelvesItems: await getAllShelvesWithItems()};
}