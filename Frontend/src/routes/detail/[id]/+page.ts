import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {

    type Item = {
        ID: string;
        Name: string;
        Description: string;
        ClassOne: boolean;
        ClassTwo: boolean;
        ClassThree: boolean;
        ClassFour: boolean;
        Damaged: boolean;
        DamagedDesc: string;
        Quantity: number;
        Status: string;
        Keywords: string[];
        Subject: string;
        Pictures: string[];
      };

    async function getItem(): Promise<Item> {
        return new Promise((resolve, reject) => {
            const id = params.id;
            fetch(API_URL + 'items/' + id, {
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


    return { item: await getItem()};
}