import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getWarehouse(): Promise<any> {
        return new Promise((resolve, reject) => {
            fetch(API_URL + 'warehouses-with-rooms/' + params.id, {
                method: 'GET',
                credentials: 'include',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                }
            }).then(async (response) => {
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

    return {
        warehouse: await getWarehouse()
    };
}