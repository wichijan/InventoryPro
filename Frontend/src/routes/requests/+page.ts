import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getAllTransferRequests() {
        const response = await fetch(API_URL + 'items/transfer-requests', {
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

    async function getAllUsers() {
        const response = await fetch(API_URL + 'users', {
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


    return { users: await getAllUsers(), transferRequests: await getAllTransferRequests() };
}