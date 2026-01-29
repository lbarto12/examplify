<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		type?: 'text' | 'email' | 'password' | 'number' | 'tel' | 'url';
		label?: string;
		placeholder?: string;
		value?: string | number;
		error?: string;
		disabled?: boolean;
		required?: boolean;
		autocomplete?: HTMLInputElement['autocomplete'];
		class?: string;
		oninput?: (event: Event) => void;
		icon?: Snippet;
	}

	let {
		type = 'text',
		label,
		placeholder = '',
		value = $bindable(''),
		error,
		disabled = false,
		required = false,
		autocomplete,
		class: className = '',
		oninput,
		icon
	}: Props = $props();

	let focused = $state(false);
	let hasValue = $derived(!!value);
	let shouldFloat = $derived(focused || hasValue);
</script>

<div class={`form-control w-full ${className}`}>
	<div class="relative">
		{#if icon}
			<div class="absolute left-4 top-1/2 -translate-y-1/2 text-base-content/50 z-10">
				{@render icon()}
			</div>
		{/if}

		<input
			{type}
			{placeholder}
			bind:value
			{disabled}
			{required}
			{autocomplete}
			class={`input input-bordered w-full ${icon ? 'pl-12' : ''} ${error ? 'input-error' : ''} ${label ? 'pt-6' : ''}`}
			{oninput}
			onfocus={() => (focused = true)}
			onblur={() => (focused = false)}
		/>

		{#if label}
			<label
				class={`absolute left-4 transition-all duration-200 pointer-events-none ${icon ? 'left-12' : 'left-4'} ${
					shouldFloat
						? 'top-2 text-xs text-primary font-medium'
						: 'top-1/2 -translate-y-1/2 text-base text-base-content/50'
				}`}
			>
				{label}
				{#if required}
					<span class="text-error">*</span>
				{/if}
			</label>
		{/if}
	</div>

	{#if error}
		<label class="label">
			<span class="label-text-alt text-error">{error}</span>
		</label>
	{/if}
</div>

<style>
	input {
		transition: all 0.2s ease;
	}
</style>
