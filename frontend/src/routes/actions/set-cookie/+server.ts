import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request, cookies }) => {
	try {
		const body = await request.json();
		const { auth } = body;

		// Cookie options - secure: false for localhost development
		const cookieOptions = {
			path: '/',
			httpOnly: false, // Allow JavaScript access for the auth plugin
			secure: false, // Set to true in production with HTTPS
			sameSite: 'lax' as const
		};

		if (auth) {
			// Set the auth cookie
			cookies.set('auth', auth, {
				...cookieOptions,
				maxAge: 60 * 60 * 24 * 7 // 7 days
			});
		} else {
			// Clear the auth cookie if auth is empty
			cookies.delete('auth', cookieOptions);
		}

		return json({ success: true });
	} catch (error) {
		console.error('Error setting cookie:', error);
		return json({ success: false, error: 'Failed to set cookie' }, { status: 500 });
	}
};
