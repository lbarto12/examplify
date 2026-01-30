import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	const authCookie = cookies.get('auth');

	if (authCookie) {
		throw redirect(303, '/dashboard');
	}

	return {};
};
