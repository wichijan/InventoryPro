import { API_URL } from "./ShelfService";

export function isUserAdmin(): Promise<boolean> {
    return new Promise((resolve, reject) => {
        resolve(true);
    });
}
export function isUserLoggedIn(): Promise<boolean> {
    return new Promise((resolve, reject) => {
        fetch(API_URL + "auth/logged-in", {
            method: "GET",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
          }).then((response) => {
            if (response.ok) {
              response.json().then((data) => {
                resolve(data.loggedIn);
              });
            } else {
                resolve(false);
            }
          });
        
    });
}