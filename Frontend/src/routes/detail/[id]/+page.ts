import { API_URL } from '$lib/_services/ShelfService.js';

export const load = async ({ fetch, params }) => {

    async function getItem(){
        // return {
        //     "ID": "2cbdabf2-f8f4-4aab-8d2e-59ef464abf6c",
        //     "Name": "Pen",
        //     "Description": "A pen is a writing instrument used to apply ink to a surface, usually paper, for writing or drawing.",
        //     "ClassOne": false,
        //     "ClassTwo": true,
        //     "ClassThree": false,
        //     "ClassFour": false,
        //     "Damaged": false,
        //     "DamagedDescription": null,
        //     "Quantity": 20,
        //     "StatusID": null
        // };
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