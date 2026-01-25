/** @type {import('./$types').PageLoad} */
export function load({ params }) {
    // This feeds 'id' into the props so {data.id} works
    return {
        id: params.id.toUpperCase()
    };
}