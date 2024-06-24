import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {

    async function getItem(): Promise<any> {
        const response = await fetch(API_URL + 'items/' + params.itemID, {
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
        item: await getItem(),
        itemID: params.itemID,
        shelfID: params.shelfID,
        warehouseID: params.id,
        roomID: params.roomID
    };
}
