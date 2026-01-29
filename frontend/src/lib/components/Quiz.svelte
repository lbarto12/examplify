<script lang="ts">
	import { Check, X } from 'lucide-svelte';
	import Card from './ui/Card.svelte';
	import Badge from './ui/Badge.svelte';

	type QuizQuestion = {
		question: string;
		options: string[];
		correct_index: number;
	};

	type QuizResult = QuizQuestion[];

	let { data }: { data: QuizResult } = $props();

	let selected = $state<number[]>([]);
	let locked = $state<boolean[]>([]);

	let score = $derived(
		locked.reduce((acc, isLocked, qIndex) => {
			if (isLocked && selected[qIndex] === data[qIndex].correct_index) {
				return acc + 1;
			}
			return acc;
		}, 0)
	);

	let totalAnswered = $derived(locked.filter(Boolean).length);
	let percentComplete = $derived(Math.round((totalAnswered / data.length) * 100));
	let allAnswered = $derived(totalAnswered === data.length);
	let scorePercent = $derived(allAnswered ? Math.round((score / data.length) * 100) : 0);

	function selectAnswer(qIndex: number, optionIndex: number) {
		if (!locked[qIndex]) {
			selected = [...selected];
			selected[qIndex] = optionIndex;

			locked = [...locked];
			locked[qIndex] = true;
		}
	}

	function getButtonClass(qIndex: number, oIndex: number) {
		const base = 'btn btn-block text-left justify-start';

		if (!locked[qIndex]) return `${base} btn-outline hover:btn-primary`;

		if (oIndex === data[qIndex].correct_index) return `${base} btn-success`;

		if (selected[qIndex] === oIndex && oIndex !== data[qIndex].correct_index)
			return `${base} btn-error`;

		return `${base} btn-ghost opacity-50`;
	}
</script>

<div class="space-y-6">
	<!-- Progress header -->
	<Card>
		<div class="flex items-center justify-between mb-4">
			<div>
				<h3 class="font-bold text-lg">Quiz Progress</h3>
				<p class="text-sm text-base-content/60">
					{totalAnswered} of {data.length} questions answered
				</p>
			</div>
			{#if allAnswered}
				<Badge variant={scorePercent >= 80 ? 'success' : scorePercent >= 60 ? 'warning' : 'error'} size="lg">
					Score: {score}/{data.length} ({scorePercent}%)
				</Badge>
			{/if}
		</div>

		<!-- Progress bar -->
		<progress class="progress progress-primary w-full" value={percentComplete} max="100"></progress>

		{#if allAnswered}
			<div class="mt-4 text-center">
				{#if scorePercent === 100}
					<p class="text-success font-bold text-xl">Perfect Score! 🎉</p>
				{:else if scorePercent >= 80}
					<p class="text-success font-semibold">Great job! Keep it up!</p>
				{:else if scorePercent >= 60}
					<p class="text-warning font-semibold">Good effort! Review the material and try again.</p>
				{:else}
					<p class="text-error font-semibold">Keep studying! You'll get there.</p>
				{/if}
			</div>
		{/if}
	</Card>

	<!-- Questions -->
	<div class="space-y-4">
		{#each data as q, qIndex (qIndex)}
			<Card class={locked[qIndex] ? 'opacity-90' : ''}>
				<div class="flex items-start gap-4">
					<div class="flex-shrink-0">
						<div
							class={`w-10 h-10 rounded-full flex items-center justify-center font-bold ${
								!locked[qIndex]
									? 'bg-primary/20 text-primary'
									: selected[qIndex] === q.correct_index
										? 'bg-success text-success-content'
										: 'bg-error text-error-content'
							}`}
						>
							{#if locked[qIndex]}
								{#if selected[qIndex] === q.correct_index}
									<Check class="w-5 h-5" />
								{:else}
									<X class="w-5 h-5" />
								{/if}
							{:else}
								{qIndex + 1}
							{/if}
						</div>
					</div>

					<div class="flex-1">
						<p class="font-semibold text-lg mb-4">{q.question}</p>

						<div class="space-y-2">
							{#each q.options as option, oIndex}
								<button
									type="button"
									class={getButtonClass(qIndex, oIndex)}
									onclick={() => selectAnswer(qIndex, oIndex)}
									disabled={locked[qIndex]}
								>
									<span class="flex items-center gap-3 w-full">
										<span
											class={`flex-shrink-0 w-6 h-6 rounded-full border-2 flex items-center justify-center ${
												locked[qIndex] && oIndex === q.correct_index
													? 'bg-success border-success'
													: locked[qIndex] && selected[qIndex] === oIndex && oIndex !== q.correct_index
														? 'bg-error border-error'
														: !locked[qIndex]
															? 'border-base-content/30'
															: 'border-base-content/10'
											}`}
										>
											{#if locked[qIndex]}
												{#if oIndex === q.correct_index}
													<Check class="w-4 h-4 text-white" />
												{:else if selected[qIndex] === oIndex}
													<X class="w-4 h-4 text-white" />
												{/if}
											{/if}
										</span>
										<span class="flex-1 text-left">{option}</span>
									</span>
								</button>
							{/each}
						</div>
					</div>
				</div>
			</Card>
		{/each}
	</div>
</div>
