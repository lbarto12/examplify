/**
 * Authentication service
 */

import { BaseService } from '../shared/service-base';
import { sessionsApiClient } from '../shared/api-client';
import { toastStore } from '../shared/toast-store.svelte';
import type { ServiceResult } from '../shared/types';
import { goto } from '$app/navigation';

interface AuthSession {
	token: string;
}

class AuthService extends BaseService {
	isAuthenticated = $state(false);
	loading = $state(false);

	constructor() {
		super();
		// Check if user is already authenticated on init (only in browser)
		if (typeof window !== 'undefined') {
			this.checkAuth();
		}
	}

	/**
	 * Sign in with email and password
	 */
	async signIn(email: string, password: string): Promise<ServiceResult<AuthSession>> {
		this.loading = true;

		const result = await this.execute(async () => {
			const response = await sessionsApiClient.signIn({ email, password });

			// Store token in cookie (call the set-cookie server action)
			await fetch('/actions/set-cookie', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ auth: response.token })
			});

			this.isAuthenticated = true;
			toastStore.success('Welcome back!');

			return response;
		});

		this.loading = false;
		return result;
	}

	/**
	 * Sign up with email and password
	 */
	async signUp(email: string, password: string): Promise<ServiceResult<AuthSession>> {
		this.loading = true;

		const result = await this.execute(async () => {
			const response = await sessionsApiClient.signUp({ email, password });

			// Store token in cookie
			await fetch('/actions/set-cookie', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ auth: response.token })
			});

			this.isAuthenticated = true;
			toastStore.success('Account created successfully!');

			return response;
		});

		this.loading = false;
		return result;
	}

	/**
	 * Sign out current user
	 */
	async signOut() {
		// Clear auth cookie
		await fetch('/actions/set-cookie', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ auth: '' })
		});

		this.isAuthenticated = false;
		toastStore.info('Signed out successfully');

		// Redirect to login
		await goto('/login');
	}

	/**
	 * Check if user is authenticated (by checking cookie)
	 */
	checkAuth(): boolean {
		// Only check auth in browser environment
		if (typeof window === 'undefined') {
			return false;
		}

		const cookie = document.cookie
			.split('; ')
			.find((row) => row.startsWith('auth='))
			?.split('=')[1];

		this.isAuthenticated = !!cookie;
		return this.isAuthenticated;
	}
}

export const authService = new AuthService();
