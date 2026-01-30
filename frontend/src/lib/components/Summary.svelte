<script lang="ts">
	import Latex from './ui/Latex.svelte';

	type SummaryResult = {
		summary: string;
	};

	let { data }: { data: SummaryResult } = $props();
</script>

<div class="card bg-base-100 p-6 shadow-md">
	<h2 class="text-xl font-bold mb-4">Summary</h2>
	<div class="prose prose-lg max-w-none">
		<Latex content={data.summary} />
	</div>
</div>
