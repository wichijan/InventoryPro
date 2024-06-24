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


    async function getRealItem() {
        const item = await getItem();
        const itemType = item.ItemTypes;
        let response;
        if (itemType === "book") {
            response = await fetch(API_URL + 'items/book/' + params.itemID, {
                method: 'GET',
                credentials: 'include',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
        } else {
            response = await fetch(API_URL + 'items/set-of-objects/' + params.itemID, {
                method: 'GET',
                credentials: 'include',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
        }

        let data = await response.json();



        data = { ...item, ...data };


        return data;
    }

    return {
        item: await getRealItem(),
        itemID: params.itemID,
        shelfID: params.shelfID,
        warehouseID: params.id,
        roomID: params.roomID
    };
}
