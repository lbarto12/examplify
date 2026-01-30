<script lang="ts">
	import { ChevronLeft, ChevronRight, RotateCcw } from 'lucide-svelte';
	import Button from './ui/Button.svelte';

	type Flashcard = {
		question: string;
		answer: string;
	};

	type FlashcardsResult = Flashcard[];

	let { data }: { data: FlashcardsResult } = $props();

	let currentIndex = $state(0);
	let isFlipped = $state(false);
	let isAnimating = $state(false);
	let slideDirection = $state<'left' | 'right'>('right');

	function nextCard() {
		if (isAnimating || currentIndex >= data.length - 1) return;
		isAnimating = true;
		slideDirection = 'right';
		isFlipped = false;

		setTimeout(() => {
			currentIndex++;
			setTimeout(() => {
				isAnimating = false;
			}, 300);
		}, 150);
	}

	function prevCard() {
		if (isAnimating || currentIndex <= 0) return;
		isAnimating = true;
		slideDirection = 'left';
		isFlipped = false;

		setTimeout(() => {
			currentIndex--;
			setTimeout(() => {
				isAnimating = false;
			}, 300);
		}, 150);
	}

	function flipCard() {
		if (isAnimating) return;
		isFlipped = !isFlipped;
	}

	function resetCards() {
		currentIndex = 0;
		isFlipped = false;
	}

	function handleKeydown(event: KeyboardEvent) {
		if (event.key === 'ArrowLeft') {
			prevCard();
		} else if (event.key === 'ArrowRight') {
			nextCard();
		} else if (event.key === ' ' || event.key === 'Enter') {
			event.preventDefault();
			flipCard();
		}
	}
</script>

<svelte:window on:keydown={handleKeydown} />

<div class="w-full max-w-2xl mx-auto">
	<!-- Progress indicator -->
	<div class="flex items-center justify-between mb-6">
		<span class="text-sm text-base-content/60">
			Card {currentIndex + 1} of {data.length}
		</span>
		<div class="flex items-center gap-2">
			<div class="flex gap-1">
				{#each data as _, i}
					<button
						class="w-2 h-2 rounded-full transition-all duration-300 {i === currentIndex
							? 'bg-primary w-6'
							: i < currentIndex
								? 'bg-primary/50'
								: 'bg-base-300'}"
						onclick={() => {
							if (!isAnimating) {
								slideDirection = i > currentIndex ? 'right' : 'left';
								isFlipped = false;
								currentIndex = i;
							}
						}}
					></button>
				{/each}
			</div>
			<button class="btn btn-ghost btn-xs" onclick={resetCards} title="Reset to first card">
				<RotateCcw class="w-4 h-4" />
			</button>
		</div>
	</div>

	<!-- Flashcard carousel -->
	<div class="relative">
		<!-- Previous button -->
		<button
			class="absolute left-0 top-1/2 -translate-y-1/2 -translate-x-4 z-10 btn btn-circle btn-ghost bg-base-100 shadow-lg border border-base-300 disabled:opacity-30"
			onclick={prevCard}
			disabled={currentIndex === 0 || isAnimating}
		>
			<ChevronLeft class="w-6 h-6" />
		</button>

		<!-- Card container -->
		<div
			class="perspective-1000 cursor-pointer mx-8"
			onclick={flipCard}
			role="button"
			tabindex="0"
			onkeydown={(e) => e.key === 'Enter' && flipCard()}
		>
			<div
				class="relative w-full aspect-[3/2] transition-all duration-500 transform-style-3d {isFlipped
					? 'rotate-y-180'
					: ''} {isAnimating ? (slideDirection === 'right' ? 'translate-x-4 opacity-0' : '-translate-x-4 opacity-0') : 'translate-x-0 opacity-100'}"
			>
				<!-- Front of card (Question) -->
				<div
					class="absolute inset-0 backface-hidden rounded-2xl bg-gradient-to-br from-primary/10 via-base-100 to-accent/10 border-2 border-primary/20 shadow-xl p-8 flex flex-col"
				>
					<div class="flex items-center gap-2 mb-4">
						<span class="badge badge-primary">Question</span>
						<span class="text-xs text-base-content/40">Click to reveal answer</span>
					</div>
					<div class="flex-1 flex items-center justify-center">
						<p class="text-xl md:text-2xl font-medium text-center leading-relaxed">
							{data[currentIndex]?.question}
						</p>
					</div>
					<div class="text-center text-xs text-base-content/40 mt-4">
						Press Space or click to flip
					</div>
				</div>

				<!-- Back of card (Answer) -->
				<div
					class="absolute inset-0 backface-hidden rotate-y-180 rounded-2xl bg-gradient-to-br from-secondary/10 via-base-100 to-primary/10 border-2 border-secondary/20 shadow-xl p-8 flex flex-col"
				>
					<div class="flex items-center gap-2 mb-4">
						<span class="badge badge-secondary">Answer</span>
						<span class="text-xs text-base-content/40">Click to see question</span>
					</div>
					<div class="flex-1 flex items-center justify-center">
						<p class="text-lg md:text-xl text-center leading-relaxed">
							{data[currentIndex]?.answer}
						</p>
					</div>
					<div class="text-center text-xs text-base-content/40 mt-4">
						Use arrow keys to navigate
					</div>
				</div>
			</div>
		</div>

		<!-- Next button -->
		<button
			class="absolute right-0 top-1/2 -translate-y-1/2 translate-x-4 z-10 btn btn-circle btn-ghost bg-base-100 shadow-lg border border-base-300 disabled:opacity-30"
			onclick={nextCard}
			disabled={currentIndex === data.length - 1 || isAnimating}
		>
			<ChevronRight class="w-6 h-6" />
		</button>
	</div>

	<!-- Navigation hints -->
	<div class="flex justify-center gap-8 mt-8 text-sm text-base-content/50">
		<div class="flex items-center gap-2">
			<kbd class="kbd kbd-sm">←</kbd>
			<span>Previous</span>
		</div>
		<div class="flex items-center gap-2">
			<kbd class="kbd kbd-sm">Space</kbd>
			<span>Flip</span>
		</div>
		<div class="flex items-center gap-2">
			<kbd class="kbd kbd-sm">→</kbd>
			<span>Next</span>
		</div>
	</div>
</div>

<style>
	.perspective-1000 {
		perspective: 1000px;
	}

	.transform-style-3d {
		transform-style: preserve-3d;
	}

	.backface-hidden {
		backface-visibility: hidden;
	}

	.rotate-y-180 {
		transform: rotateY(180deg);
	}
</style>
