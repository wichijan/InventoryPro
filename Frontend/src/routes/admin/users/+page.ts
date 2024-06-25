import { API_URL } from '$lib/_services/ShelfService';

export const load = async ({ fetch, params }) => {


    async function registrationRequests(): Promise<any> {
        return new Promise((resolve, reject) => {
            fetch(API_URL + 'registration-requests', {
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
    async function getUsers() {
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
    async function getRoles() {
        const response = await fetch(API_URL + 'roles', {
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
        roles: await getRoles(),
        registrationRequests: await registrationRequests(),
        users: await getUsers()
    };
}