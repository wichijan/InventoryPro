import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {

    async function getRoomWithShelves(): Promise<any> {
        const response = await fetch(API_URL + 'rooms-with-shelves/' + params.roomID, {
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

    async function getShelvesWithItems(): Promise<any> {
        const roomData = await getRoomWithShelves();

        if (!roomData.Shelves) {
            return [];
        }

        const shelvesPromises = roomData.Shelves.map(async (shelf: any) => {
            const response = await fetch(API_URL + 'shelves-with-items/' + shelf.ID, {
                method: 'GET',
                credentials: 'include',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            const data = await response.json();
            return data;
        });

        const shelvesWithItems = await Promise.all(shelvesPromises);
        return shelvesWithItems;
    }
    async function getRoom(): Promise<any> {
        const response = await fetch(API_URL + 'rooms/' + params.roomID, {
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
        shelves: await getShelvesWithItems(),
        warehouseID: params.id,
        roomID: params.roomID,
        room: await getRoom()
    };
}
