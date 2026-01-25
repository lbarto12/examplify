/** @type {import('./$types').PageLoad} */
export function load() {
    return {
        // Mock data to be replaced by backend later
        savedCourses: [
            { id: 'cs221', name: 'CS221' },
            { id: 'mat470', name: 'MAT470' },
            { id: 'phil101', name: 'PHIL101' }
        ],
        quickAccess: [
            { label: 'Lecture', icon: '📚' },
            { label: 'Summary', icon: '📝' },
            { label: 'Flash Cards', icon: '📇' },
            { label: 'Quizzes', icon: '❓' }
        ]
    };
}