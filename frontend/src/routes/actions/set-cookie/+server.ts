import type { RequestHandler } from './$types';
import { json } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, cookies }) => {
  const { key, value } = await request.json();

  // set cookie
  cookies.set(key, value, {
    path: '/',
    httpOnly: false,        // must be false if JS needs to read it
    sameSite: 'lax',
    secure: false,          // set true in production (HTTPS)
    maxAge: 60 * 60 * 24,   // 1 day
  });

  return json({ ok: true });
};
