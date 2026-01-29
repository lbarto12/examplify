<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		variant?: 'primary' | 'secondary' | 'accent' | 'success' | 'warning' | 'error' | 'info' | 'neutral';
		size?: 'sm' | 'md' | 'lg';
		class?: string;
		children: Snippet;
	}

	let {
		variant = 'primary',
		size = 'md',
		class: className = '',
		children
	}: Props = $props();

	const variantClasses = {
		primary: 'badge-primary',
		secondary: 'badge-secondary',
		accent: 'badge-accent',
		success: 'badge-success',
		warning: 'badge-warning',
		error: 'badge-error',
		info: 'badge-info',
		neutral: 'badge-neutral'
	};

	const sizeClasses = {
		sm: 'badge-sm',
		md: 'badge-md',
		lg: 'badge-lg'
	};
</script>

<span class={`badge ${variantClasses[variant]} ${sizeClasses[size]} ${className}`}>
	{@render children()}
</span>
