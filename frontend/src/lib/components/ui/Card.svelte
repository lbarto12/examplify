<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		hover?: boolean;
		clickable?: boolean;
		gradient?: boolean;
		onclick?: (event: MouseEvent) => void;
		class?: string;
		children: Snippet;
		header?: Snippet;
		footer?: Snippet;
	}

	let {
		hover = false,
		clickable = false,
		gradient = false,
		onclick,
		class: className = '',
		children,
		header,
		footer
	}: Props = $props();

	const hoverClass = hover || clickable ? 'card-hover' : '';
	const cursorClass = clickable ? 'cursor-pointer' : '';
	const gradientClass = gradient ? 'border-2 border-transparent bg-gradient-to-br from-brand-500 to-accent-pink p-[2px]' : '';
</script>

<div
	class={`card bg-base-100 shadow-lg ${hoverClass} ${cursorClass} ${gradientClass} ${className}`}
	role={clickable ? 'button' : undefined}
	tabindex={clickable ? 0 : undefined}
	{onclick}
	onkeydown={(e) => {
		if (clickable && (e.key === 'Enter' || e.key === ' ')) {
			e.preventDefault();
			onclick?.(e as unknown as MouseEvent);
		}
	}}
>
	{#if gradient}
		<div class="bg-base-100 rounded-xl h-full">
			{#if header}
				<div class="card-header p-6 pb-0">
					{@render header()}
				</div>
			{/if}

			<div class="card-body p-6">
				{@render children()}
			</div>

			{#if footer}
				<div class="card-footer p-6 pt-0">
					{@render footer()}
				</div>
			{/if}
		</div>
	{:else}
		{#if header}
			<div class="card-header p-6 pb-0">
				{@render header()}
			</div>
		{/if}

		<div class="card-body p-6">
			{@render children()}
		</div>

		{#if footer}
			<div class="card-footer p-6 pt-0">
				{@render footer()}
			</div>
		{/if}
	{/if}
</div>

<style>
	.card {
		border-radius: 1rem;
		transition: all 0.2s ease;
	}
</style>
