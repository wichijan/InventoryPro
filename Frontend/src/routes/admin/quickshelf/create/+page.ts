import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getAllWarehousesWithRooms(): Promise<any> {
        const response = await fetch(API_URL + 'warehouses-with-rooms', {
            method: 'GET',
            credentials: 'include',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        const data = await response.json();
        return data;
    }

    return {
        warehouses: await getAllWarehousesWithRooms()
    };
}