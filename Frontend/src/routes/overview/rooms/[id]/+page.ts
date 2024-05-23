import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {

    async function getRoom(): Promise<any> {
        return new Promise((resolve, reject) => {
            fetch(API_URL + 'rooms/'+params.id, {
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


    return {room: await getRoom()};
}