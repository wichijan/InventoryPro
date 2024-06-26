import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getQuickShelf(): Promise<any> {
        const response = await fetch(API_URL + 'quick-shelves/' + params.id, {
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
        quickShelf: await getQuickShelf(),
        warehouses: await getAllWarehousesWithRooms()
    };
}