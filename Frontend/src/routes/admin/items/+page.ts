import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getItems(): Promise<any> {
        const response = await fetch(API_URL + 'items', {
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
    async function getQuickShelves(): Promise<any> {
        const response = await fetch(API_URL + 'quick-shelves', {
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
        quickShelves: await getQuickShelves(),
        items: await getItems(),
        warehouses: await getAllWarehousesWithRooms()
    };
}