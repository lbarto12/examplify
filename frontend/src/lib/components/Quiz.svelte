<script lang="ts">
	import '../../app.css';

	type QuizQuestion = {
		question: string;
		options: string[];
		correct_index: number;
	};

	type QuizResult = QuizQuestion[];

	let { data }: { data: QuizResult } = $props();

	// Use $state for reactive arrays
	let selected = $state<number[]>([]);
	let locked = $state<boolean[]>([]);

	function selectAnswer(qIndex: number, optionIndex: number) {
		// Only lock/select if not already locked
		if (!locked[qIndex]) {
			// Reassign arrays to trigger reactivity
			selected = [...selected];
			selected[qIndex] = optionIndex;

			locked = [...locked];
			locked[qIndex] = true;
		}
	}

	function getButtonClass(qIndex: number, oIndex: number) {
		if (!locked[qIndex]) return 'btn btn-sm btn-outline text-left';

		// Correct answer
		if (oIndex === data[qIndex].correct_index) return 'btn btn-sm text-left btn-success';

		// Wrong answer that was selected
		if (selected[qIndex] === oIndex && oIndex !== data[qIndex].correct_index)
			return 'btn btn-sm text-left btn-error';

		// Other buttons after lock
		return 'btn btn-sm btn-outline text-left';
	}
</script>

<div class="card bg-base-100 p-4 shadow-md">
	<h2 class="text-xl font-bold mb-2">Quiz</h2>
	<div class="flex flex-col gap-4">
		{#each data as q, qIndex}
			<div class="border p-3 rounded">
				<p class="font-semibold">Q{qIndex + 1}: {q.question}</p>
				<div class="flex flex-col gap-2 mt-2">
					{#each q.options as option, oIndex}
						<button
							type="button"
							class={getButtonClass(qIndex, oIndex)}
							on:click={() => selectAnswer(qIndex, oIndex)}
						>
							{option}
						</button>
					{/each}
				</div>
			</div>
		{/each}
	</div>
</div>
