<script lang="ts">
	import type { Snippet } from 'svelte';
	import { X } from 'lucide-svelte';
	import { fly, fade } from 'svelte/transition';

	interface Props {
		open?: boolean;
		title?: string;
		size?: 'sm' | 'md' | 'lg' | 'xl';
		onclose?: () => void;
		closeOnBackdrop?: boolean;
		children: Snippet;
		actions?: Snippet;
	}

	let {
		open = $bindable(false),
		title,
		size = 'md',
		onclose,
		closeOnBackdrop = true,
		children,
		actions
	}: Props = $props();

	const sizeClasses = {
		sm: 'modal-box max-w-sm',
		md: 'modal-box max-w-2xl',
		lg: 'modal-box max-w-4xl',
		xl: 'modal-box max-w-6xl'
	};

	function handleClose() {
		open = false;
		onclose?.();
	}

	function handleBackdropClick() {
		if (closeOnBackdrop) {
			handleClose();
		}
	}

	function handleKeyDown(e: KeyboardEvent) {
		if (e.key === 'Escape') {
			handleClose();
		}
	}
</script>

{#if open}
	<div
		class="modal modal-open"
		role="dialog"
		aria-modal="true"
		tabindex="-1"
		onkeydown={handleKeyDown}
	>
		<div
			class="modal-backdrop"
			onclick={handleBackdropClick}
			transition:fade={{ duration: 200 }}
		></div>

		<div
			class={sizeClasses[size]}
			transition:fly={{ y: 20, duration: 300 }}
		>
			{#if title}
				<div class="flex items-center justify-between mb-4">
					<h3 class="font-bold text-2xl">{title}</h3>
					<button
						class="btn btn-sm btn-circle btn-ghost"
						onclick={handleClose}
						aria-label="Close modal"
					>
						<X class="w-5 h-5" />
					</button>
				</div>
			{:else}
				<button
					class="btn btn-sm btn-circle btn-ghost absolute right-4 top-4"
					onclick={handleClose}
					aria-label="Close modal"
				>
					<X class="w-5 h-5" />
				</button>
			{/if}

			<div class="modal-content">
				{@render children()}
			</div>

			{#if actions}
				<div class="modal-action">
					{@render actions()}
				</div>
			{/if}
		</div>
	</div>
{/if}

<style>
	.modal-backdrop {
		position: fixed;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		background-color: rgba(0, 0, 0, 0.5);
		backdrop-filter: blur(4px);
	}
</style>
