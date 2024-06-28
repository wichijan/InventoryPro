import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getUserItems(): Promise<any> {
        const response = await fetch(API_URL + 'users/items', {
            method: 'GET',
            credentials: 'include',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        const data = await response.json();

        return data ?? [];
    }

    return {
        items: await getUserItems()
    };
}