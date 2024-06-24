import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function getRoles(): Promise<any> {
        const response = await fetch(API_URL + 'roles', {
            method: 'GET',
            credentials: 'include',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        const data = await response.json();
        console.log(data);
        return data;
    }

    return {
        roles: await getRoles()
    };
}