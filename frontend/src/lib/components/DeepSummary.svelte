<script lang="ts">
	import Latex from './ui/Latex.svelte';
	import { BookOpen } from 'lucide-svelte';

	type DeepSummaryItem = {
		concept: string;
		details: string;
		definition: string;
	};

	type DeepSummaryResult = DeepSummaryItem[];

	let { data }: { data: DeepSummaryResult } = $props();
</script>

<div class="space-y-6">
	<h2 class="text-xl font-bold flex items-center gap-2">
		<BookOpen class="w-5 h-5 text-primary" />
		Deep Summary
	</h2>

	<div class="grid gap-4">
		{#each data as item, index}
			<div class="card bg-base-100 shadow-md border border-base-300 overflow-hidden">
				<div class="bg-gradient-to-r from-primary/10 to-secondary/10 px-6 py-4 border-b border-base-300">
					<h3 class="font-bold text-lg">
						<Latex content={item.concept} />
					</h3>
				</div>
				<div class="p-6 space-y-4">
					<div>
						<h4 class="text-sm font-semibold text-base-content/60 uppercase tracking-wide mb-2">Details</h4>
						<div class="text-base-content">
							<Latex content={item.details} />
						</div>
					</div>
					{#if item.definition}
						<div class="bg-base-200/50 rounded-lg p-4 border-l-4 border-primary">
							<h4 class="text-sm font-semibold text-base-content/60 uppercase tracking-wide mb-2">Definition</h4>
							<div class="italic text-base-content/80">
								<Latex content={item.definition} />
							</div>
						</div>
					{/if}
				</div>
			</div>
		{/each}
	</div>
</div>
