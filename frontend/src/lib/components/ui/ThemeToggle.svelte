<script lang="ts">
	import { Sun, Moon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let theme = $state<'examplify' | 'examplify-dark'>('examplify');

	function applyTheme(newTheme: 'examplify' | 'examplify-dark') {
		document.documentElement.setAttribute('data-theme', newTheme);
		localStorage.setItem('theme', newTheme);
	}

	onMount(() => {
		// Load theme from localStorage
		const saved = localStorage.getItem('theme');
		if (saved === 'examplify-dark' || saved === 'examplify') {
			theme = saved;
			applyTheme(saved);
		} else {
			// Check system preference
			const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
			const initialTheme = prefersDark ? 'examplify-dark' : 'examplify';
			theme = initialTheme;
			applyTheme(initialTheme);
		}
	});

	function toggleTheme() {
		const newTheme = theme === 'examplify' ? 'examplify-dark' : 'examplify';
		theme = newTheme;
		applyTheme(newTheme);
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
