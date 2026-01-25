<script lang="ts">
	import '../../app.css';

    type Flashcard = {
        question: string;
        answer: string;
    };

    type FlashcardsResult = Flashcard[];

	let { data }: {data: FlashcardsResult} = $props();
</script>

<div class="card bg-base-100 p-4 shadow-md">
	<h2 class="text-xl font-bold mb-2">Flashcards</h2>
	<div class="flex flex-col gap-3">
		{#each data as fc, index}
			<div class="border p-3 rounded">
				<p class="font-semibold">Q{index + 1}: {fc.question}</p>
				<p class="text-gray-700 mt-1">A: {fc.answer}</p>
			</div>
		{/each}
	</div>
</div>
<style>
	@import '../../app.css';
</style>