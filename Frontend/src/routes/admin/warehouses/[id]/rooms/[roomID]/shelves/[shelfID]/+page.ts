import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {

    async function getShelvesWithItems(): Promise<any> {
        const response = await fetch(API_URL + 'shelves-with-items/' + params.shelfID, {
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
        shelf: await getShelvesWithItems(),
        shelfID: params.shelfID,
        warehouseID: params.id,
        roomID: params.roomID
    };
}
