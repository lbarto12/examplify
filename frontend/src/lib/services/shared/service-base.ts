/**
 * Base service class with error handling and loading state management
 */

import type { ServiceError, ServiceErrorCode, ServiceResult } from './types';
import { toastStore } from './toast-store.svelte';

/**
 * Base class for all services providing common functionality
 */
export abstract class BaseService {
	/**
	 * Execute an async operation with automatic error handling
	 * @param operation The async operation to execute
	 * @param showErrorToast Whether to show error as toast (default: true)
	 * @returns ServiceResult with data or error
	 */
	protected async execute<T>(
		operation: () => Promise<T>,
		showErrorToast = true
	): Promise<ServiceResult<T>> {
		try {
			const data = await operation();
			return {
				data,
				loading: false
			};
		} catch (error) {
			const serviceError = this.handleError(error);

			if (showErrorToast) {
				toastStore.error(serviceError.message);
			}

			return {
				error: serviceError,
				loading: false
			};
		}
	}

	/**
	 * Convert unknown error into ServiceError
	 */
	protected handleError(error: unknown): ServiceError {
		// If it's already a ServiceError, return it
		if (this.isServiceError(error)) {
			return error;
		}

		// Handle Zodios/Axios errors
		if (this.isHttpError(error)) {
			const statusCode = error.response?.status;
			const message = error.response?.data?.message || error.message;

			return {
				code: this.getErrorCode(statusCode),
				message: this.getUserFriendlyMessage(message, statusCode),
				details: error.response?.data,
				statusCode
			};
		}

		// Handle generic errors
		if (error instanceof Error) {
			return {
				code: 'UNKNOWN_ERROR',
				message: error.message || 'An unexpected error occurred'
			};
		}

		// Fallback
		return {
			code: 'UNKNOWN_ERROR',
			message: 'An unexpected error occurred'
		};
	}

	/**
	 * Check if error is a ServiceError
	 */
	private isServiceError(error: unknown): error is ServiceError {
		return (
			typeof error === 'object' &&
			error !== null &&
			'code' in error &&
			'message' in error
		);
	}

	/**
	 * Check if error is an HTTP error (from Axios/Zodios)
	 */
	private isHttpError(error: unknown): error is { response?: { status: number; data?: { message?: string } }; message: string } {
		return (
			typeof error === 'object' &&
			error !== null &&
			'response' in error
		);
	}

	/**
	 * Map HTTP status code to ServiceErrorCode
	 */
	private getErrorCode(statusCode?: number): ServiceErrorCode {
		if (!statusCode) return 'NETWORK_ERROR';

		if (statusCode === 401 || statusCode === 403) return 'AUTH_ERROR';
		if (statusCode === 404) return 'NOT_FOUND';
		if (statusCode >= 400 && statusCode < 500) return 'VALIDATION_ERROR';
		if (statusCode >= 500) return 'SERVER_ERROR';

		return 'UNKNOWN_ERROR';
	}

	/**
	 * Get user-friendly error message
	 */
	private getUserFriendlyMessage(message: string, statusCode?: number): string {
		// Use backend message if available
		if (message && message !== 'Request failed') {
			return message;
		}

		// Fallback based on status code
		switch (statusCode) {
			case 401:
				return 'You need to sign in to perform this action';
			case 403:
				return 'You do not have permission to perform this action';
			case 404:
				return 'The requested resource was not found';
			case 500:
				return 'Server error. Please try again later';
			default:
				return 'An error occurred. Please try again';
		}
	}
}
