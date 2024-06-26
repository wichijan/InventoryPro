import { writable } from "svelte/store";
type State = {
    requests: Array<Request>;
};
export const state = writable<State>({
    requests: [],
});
export const connect = () => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    ws.addEventListener("message", (message: any) => {
        const data: Request = JSON.parse(message.data);
        state.update((state) => ({
            ...state,
            requests: [data].concat(state.requests),
        }));
    });
};