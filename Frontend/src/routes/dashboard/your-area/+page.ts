import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getUser(): Promise<any> {
        const response = await fetch(API_URL + 'users/get-me', {
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
        userData: await getUser()
    };
}