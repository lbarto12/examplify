/**
 * Shared types for service layer
 */

/**
 * Error types that can occur during service operations
 */
export type ServiceErrorCode =
	| 'NETWORK_ERROR'
	| 'AUTH_ERROR'
	| 'VALIDATION_ERROR'
	| 'SERVER_ERROR'
	| 'NOT_FOUND'
	| 'UNKNOWN_ERROR';

/**
 * Structured error returned by services
 */
export interface ServiceError {
	code: ServiceErrorCode;
	message: string;
	details?: unknown;
	statusCode?: number;
}

/**
 * Result wrapper for service operations
 */
export interface ServiceResult<T> {
	data?: T;
	error?: ServiceError;
	loading: boolean;
}

/**
 * Toast notification type
 */
export type ToastType = 'success' | 'error' | 'info' | 'warning';

/**
 * Toast notification message
 */
export interface ToastMessage {
	id: string;
	type: ToastType;
	message: string;
	duration?: number;
}
