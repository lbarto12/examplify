
/** @type {import('./$types').PageLoad} */
export function load({ params }) {
    return {
        // This takes whatever is in the URL and passes it to the Svelte file
        courseId: params.courseID 
    };
}