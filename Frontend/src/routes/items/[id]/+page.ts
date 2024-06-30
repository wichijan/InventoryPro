import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {

    type Item = {
        ID: string;
        ItemTypes: string;
        QuantityInShelf: number;
        Name: string;
        Description: string;
        RegularShelfID: string;
        ClassOne: boolean;
        ClassTwo: boolean;
        ClassThree: boolean;
        ClassFour: boolean;
        Damaged: boolean;
        DamagedDesc: string;
        Picture: string;
        HintText: string;
        Quantity: number;
        UsersBorrowed: any[];
        Keywords: any[];
        Subject: any[];
        Reservations: any[];
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


    async function getQuickshelves() {
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


    return { item: await getItem(), users: await getAllUsers(), quickshelves: await getQuickshelves(), warehouses: await getAllWarehousesWithRooms() };
}