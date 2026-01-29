<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		variant?: 'primary' | 'secondary' | 'ghost' | 'outline' | 'gradient';
		size?: 'sm' | 'md' | 'lg';
		loading?: boolean;
		disabled?: boolean;
		type?: 'button' | 'submit' | 'reset';
		onclick?: (event: MouseEvent) => void;
		class?: string;
		children: Snippet;
		icon?: Snippet;
	}

	let {
		variant = 'primary',
		size = 'md',
		loading = false,
		disabled = false,
		type = 'button',
		onclick,
		class: className = '',
		children,
		icon
	}: Props = $props();

	const variantClasses = {
		primary: 'btn btn-primary gradient-primary text-white btn-hover',
		secondary: 'btn btn-secondary btn-hover',
		ghost: 'btn btn-ghost btn-hover',
		outline: 'btn btn-outline btn-hover',
		gradient: 'gradient-primary text-white btn btn-hover'
	};

	const sizeClasses = {
		sm: 'btn-sm',
		md: 'btn-md',
		lg: 'btn-lg'
	};

	const isDisabled = $derived(disabled || loading);
</script>

<button
	{type}
	class={`${variantClasses[variant]} ${sizeClasses[size]} ${className}`}
	disabled={isDisabled}
	{onclick}
>
	{#if loading}
		<span class="loading loading-spinner loading-sm"></span>
	{:else if icon}
		{@render icon()}
	{/if}
	{@render children()}
</button>

<style>
	button {
		transition: all 0.2s ease;
	}
</style>
