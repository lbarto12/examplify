/**
 * Toast notification store for global toast messages
 */

import type { ToastMessage, ToastType } from './types';

class ToastStore {
	private toasts = $state<ToastMessage[]>([]);
	private idCounter = 0;

	/**
	 * Get all active toasts (reactive)
	 */
	get messages(): ToastMessage[] {
		return this.toasts;
	}

	/**
	 * Add a toast notification
	 */
	private addToast(type: ToastType, message: string, duration = 5000) {
		const id = `toast-${++this.idCounter}`;
		const toast: ToastMessage = {
			id,
			type,
			message,
			duration
		};

		this.toasts = [...this.toasts, toast];

		// Auto-dismiss
		if (duration > 0) {
			setTimeout(() => this.remove(id), duration);
		}
	}

	/**
	 * Show success toast
	 */
	success(message: string, duration?: number) {
		this.addToast('success', message, duration);
	}

	/**
	 * Show error toast
	 */
	error(message: string, duration?: number) {
		this.addToast('error', message, duration);
	}

	/**
	 * Show info toast
	 */
	info(message: string, duration?: number) {
		this.addToast('info', message, duration);
	}

	/**
	 * Show warning toast
	 */
	warning(message: string, duration?: number) {
		this.addToast('warning', message, duration);
	}

	/**
	 * Remove a toast by ID
	 */
	remove(id: string) {
		this.toasts = this.toasts.filter((t) => t.id !== id);
	}

	/**
	 * Clear all toasts
	 */
	clear() {
		this.toasts = [];
	}
}

export const toastStore = new ToastStore();
