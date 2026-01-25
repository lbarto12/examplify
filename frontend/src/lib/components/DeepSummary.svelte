<script lang="ts">
	import '../../app.css';
    type DeepSummaryItem = {
        concept: string;
        details: string;
        definition: string;
    };

    type DeepSummaryResult = DeepSummaryItem[];

    let { data }: { data: DeepSummaryResult } = $props();

</script>

<div class="card bg-base-100 p-4 shadow-md">
	<h2 class="text-xl font-bold mb-2">Deep Summary</h2>
	<div class="flex flex-col gap-4">
		{#each data as item}
			<div class="border p-3 rounded">
				<p class="font-bold">{item.concept}</p>
				<p class="text-gray-700 mt-1">{item.details}</p>
				<p class="italic text-gray-500 mt-1">{item.definition}</p>
			</div>
		{/each}
	</div>
</div>

<style>
	@import '../../app.css';
</style>