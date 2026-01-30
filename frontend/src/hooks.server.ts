import { redirect, type Handle } from '@sveltejs/kit';

// Routes that require authentication
const protectedRoutes = ['/dashboard'];

export const handle: Handle = async ({ event, resolve }) => {
	const authCookie = event.cookies.get('auth');
	const path = event.url.pathname;

	// Check if accessing a protected route without auth
	const isProtectedRoute = protectedRoutes.some((route) => path.startsWith(route));
	if (isProtectedRoute && !authCookie) {
		throw redirect(303, '/login');
	}

	return resolve(event);
};
