<script lang="ts">
    import '../../app.css';
    
    import SummaryView from './Summary.svelte';
	import FlashcardsView from './FlashCard.svelte';
	import QuizView from './Quiz.svelte';
	import DeepSummaryView from './DeepSummary.svelte';
    
    type SummaryResult = {
        summary: string;
    };

    type Flashcard = {
        question: string;
        answer: string;
    };

    type FlashcardsResult = Flashcard[];

    type QuizQuestion = {
        question: string;
        options: string[];
        correct_index: number;
    };

    type QuizResult = QuizQuestion[];

    type DeepSummaryItem = {
        concept: string;
        details: string;
        definition: string;
    };

    type DeepSummaryResult = DeepSummaryItem[];

    type AIResult =
	| { type: "summary"; result: SummaryResult; id: string; createdAt: string }
	| { type: "flashcards"; result: FlashcardsResult; id: string; createdAt: string }
	| { type: "quiz"; result: QuizResult; id: string; createdAt: string }
	| { type: "deep_summary"; result: DeepSummaryResult; id: string; createdAt: string };

	export let aiResult: AIResult;
</script>

{#if aiResult.type === "summary"}
	<SummaryView data={aiResult.result} />
{:else if aiResult.type === "flashcards"}
	<FlashcardsView data={aiResult.result} />
{:else if aiResult.type === "quiz"}
	<QuizView data={aiResult.result} />
{:else if aiResult.type === "deep_summary"}
	<DeepSummaryView data={aiResult.result} />
{:else}
	<p>Unknown AI result type</p>
{/if}