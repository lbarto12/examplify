<script lang="ts">
	import { page } from '$app/stores';
	import { ChevronRight, Home } from 'lucide-svelte';

	interface Breadcrumb {
		label: string;
		href: string;
	}

	let breadcrumbs = $derived.by(() => {
		const path = $page.url.pathname;
		const segments = path.split('/').filter(Boolean);

		const crumbs: Breadcrumb[] = [
			{ label: 'Dashboard', href: '/dashboard' }
		];

		// Build breadcrumbs from URL segments
		let currentPath = '/dashboard';
		for (let i = 1; i < segments.length; i++) {
			currentPath += `/${segments[i]}`;

			// Try to make labels more readable
			let label = segments[i];

			// If it looks like a UUID, try to get a better name from page data
			if (label.length > 20 && label.includes('-')) {
				// For now, use generic labels
				if (i === 1) label = 'Course';
				else if (i === 2) label = 'Collection';
				else if (i === 3) label = 'Analysis';
			} else {
				// Capitalize and replace hyphens
				label = label.replace(/-/g, ' ')
					.split(' ')
					.map(w => w.charAt(0).toUpperCase() + w.slice(1))
					.join(' ');
			}

			crumbs.push({ label, href: currentPath });
		}

		return crumbs;
	});
</script>

<div class="breadcrumbs text-sm">
	<ul>
		{#each breadcrumbs as crumb, i}
			<li>
				{#if i === breadcrumbs.length - 1}
					<span class="text-primary font-semibold">{crumb.label}</span>
				{:else}
					<a href={crumb.href} class="hover:text-primary transition-colors">
						{#if i === 0}
							<Home class="w-4 h-4" />
						{/if}
						{crumb.label}
					</a>
				{/if}
			</li>
		{/each}
	</ul>
</div>
