import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {

    type Warehouse = {
        Description: string;
        Id: string;
        Name: string;
      };

    async  function getAllWarehouses(): Promise<Warehouse[]> {
        return new Promise((resolve, reject) => {
            fetch(API_URL + 'warehouses', {
                method: 'GET',
                credentials: 'include',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(async (response) => {
                if (response.ok) {
                    let aPromise = [];
                    await response.json().then((data) => {
                        aPromise = data.map(async (warehouse) => {
                            let warehouseData = await fetch(API_URL + 'warehouses/' + warehouse.ID, {
                                method: 'GET',
                                credentials: 'include',
                                mode: 'cors',
                                headers: {
                                    'Content-Type': 'application/json'
                                }
                            });
                            return warehouseData.json();
                        });
                        Promise.all(aPromise).then((data) => {
                            resolve(data);
                            (data);
                        });
                    });
                } else {
                    reject(response.statusText);
                }
            });
        });
    }

    async function getRooms(): Promise<any> {
        return new Promise((resolve, reject) => {
            fetch(API_URL + 'rooms', {
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


    return {warehouses: await getAllWarehouses(), rooms: await getRooms()};
}