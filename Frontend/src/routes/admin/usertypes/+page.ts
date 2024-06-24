import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getUserTypes(): Promise<any> {
        const response = await fetch(API_URL + 'user-types', {
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
        usertypes: await getUserTypes()
    };
}