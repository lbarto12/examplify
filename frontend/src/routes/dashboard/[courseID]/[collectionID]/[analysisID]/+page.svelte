<script lang="ts">
	import { onMount } from 'svelte';
	import { getClient } from '$lib/apis/core.svelte';
    import SummaryView from '$lib/components/Summary.svelte';
	import FlashcardsView from '$lib/components/FlashCard.svelte';
	import QuizView from '$lib/components/Quiz.svelte';
	import DeepSummaryView from '$lib/components/DeepSummary.svelte';

	type SummaryResult = { summary: string };
	type Flashcard = { question: string; answer: string };
	type FlashcardsResult = Flashcard[];
	type QuizQuestion = { question: string; options: string[]; correct_index: number };
	type QuizResult = QuizQuestion[];
	type DeepSummaryItem = { concept: string; details: string; definition: string };
	type DeepSummaryResult = DeepSummaryItem[];
	type AIResult =
		| { type: "summary"; result: SummaryResult; id: string; createdAt: string }
		| { type: "flashcards"; result: FlashcardsResult; id: string; createdAt: string }
		| { type: "quiz"; result: QuizResult; id: string; createdAt: string }
		| { type: "deep_summary"; result: DeepSummaryResult; id: string; createdAt: string };

    const { data } = $props<{
		data: {
			courseId: string;
            collectionID: string;
            analysisID: string;
		};
	}>();

	const { courseId, collectionID, analysisID } = data;

	// initialize to null
	let result: AIResult | null = $state(null);

    onMount(async () => {
        const api = getClient();

        try {
            const d = await api.getAnalysis({
                params: {
                    id: analysisID,
                    analysisID: analysisID,
                }
            });

            // Optionally validate or cast
            result = {
                ...d,
                result: typeof d.result === 'string' ? JSON.parse(d.result) : d.result
            } as AIResult;
            console.log("AI Result:", result);
        } catch (err) {
            console.error("Failed to fetch AI result", err);
        }
    });
</script>

{#if result}
    {#if result.type === "summary"}
	    <SummaryView data={result.result} />
    {:else if result.type === "flashcards"}
	    <FlashcardsView data={result.result} />
    {:else if result.type === "quiz"}
	    <QuizView data={result.result} />
    {:else if result.type === "deep_summary"}
	    <DeepSummaryView data={result.result} />
    {:else}
	    <p>Unknown AI result type</p>
    {/if}
{:else}
    <p>Loading AI analysis...</p>
{/if}
