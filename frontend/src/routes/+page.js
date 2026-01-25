// src/routes/+page.js
/** @type {import('./$types').PageLoad} */
export function load() {
    return {
        // FAKE (for now) Data based on sketches
        savedCourses: [
            { name: 'CS221' },
            { name: 'MAT470' },
            { name: 'PHIL101' }
        ],
        quickAccess: [
            { label: 'Lecture', icon: '📚' },
            { label: 'Summary', icon: '📝' },
            { label: 'Flash Cards', icon: '📇' }
        ]
    };
}