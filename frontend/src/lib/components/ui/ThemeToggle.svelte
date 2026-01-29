<script lang="ts">
	import { Sun, Moon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let theme = $state<'examplify' | 'examplify-dark'>('examplify');

	onMount(() => {
		// Load theme from localStorage
		const saved = localStorage.getItem('theme');
		if (saved === 'examplify-dark' || saved === 'examplify') {
			theme = saved;
		} else {
			// Check system preference
			const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
			theme = prefersDark ? 'examplify-dark' : 'examplify';
		}
		applyTheme();
	});

	function applyTheme() {
		if (typeof window !== 'undefined') {
			document.documentElement.setAttribute('data-theme', theme);
			localStorage.setItem('theme', theme);
		}
	}

	function toggleTheme() {
		theme = theme === 'examplify' ? 'examplify-dark' : 'examplify';
		applyTheme();
	}
</script>

<button
	class="btn btn-circle btn-ghost swap swap-rotate"
	onclick={toggleTheme}
	aria-label="Toggle theme"
>
	{#if theme === 'examplify'}
		<Moon class="w-5 h-5" />
	{:else}
		<Sun class="w-5 h-5" />
	{/if}
</button>
