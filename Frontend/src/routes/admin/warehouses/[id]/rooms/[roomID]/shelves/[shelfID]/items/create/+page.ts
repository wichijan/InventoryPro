export const load = async ({ fetch, params }) => {

    return {
        shelfID: params.shelfID,
        warehouseID: params.id,
        roomID: params.roomID
    };
}
