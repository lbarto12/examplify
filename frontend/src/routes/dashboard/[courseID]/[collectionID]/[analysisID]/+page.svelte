<script lang="ts">
	import { onMount } from 'svelte';
	import { analysesService } from '$lib/services';
	import SummaryView from '$lib/components/Summary.svelte';
	import FlashcardsView from '$lib/components/FlashCard.svelte';
	import QuizView from '$lib/components/Quiz.svelte';
	import DeepSummaryView from '$lib/components/DeepSummary.svelte';
	import Skeleton from '$lib/components/ui/Skeleton.svelte';
	import Badge from '$lib/components/ui/Badge.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import {
		FileText,
		Brain,
		ClipboardList,
		BookMarked,
		Download,
		Share2,
		RotateCw
	} from 'lucide-svelte';

	const { data } = $props<{
		data: {
			courseId: string;
			collectionID: string;
			analysisID: string;
		};
	}>();

	const { courseId, collectionID, analysisID } = data;

	const analysisTypeConfig = {
		summary: {
			title: 'Summary',
			icon: FileText,
			gradient: 'from-purple-500 to-pink-500',
			badge: 'primary'
		},
		flashcards: {
			title: 'Flashcards',
			icon: Brain,
			gradient: 'from-teal-500 to-blue-500',
			badge: 'secondary'
		},
		quiz: {
			title: 'Quiz',
			icon: ClipboardList,
			gradient: 'from-orange-500 to-red-500',
			badge: 'accent'
		},
		deep_summary: {
			title: 'Deep Summary',
			icon: BookMarked,
			gradient: 'from-indigo-500 to-purple-500',
			badge: 'info'
		}
	};

	onMount(async () => {
		await analysesService.getById(collectionID, analysisID);
	});

	async function handleExport() {
		if (!analysesService.currentAnalysis) return;

		// Create a blob with the analysis content
		const content = JSON.stringify(analysesService.currentAnalysis.result, null, 2);
		const blob = new Blob([content], { type: 'application/json' });
		const url = URL.createObjectURL(blob);

		// Create download link
		const a = document.createElement('a');
		a.href = url;
		a.download = `${analysesService.currentAnalysis.type}_${analysisID}.json`;
		a.click();

		URL.revokeObjectURL(url);
	}
</script>

<div class="space-y-6">
	{#if analysesService.loading}
		<!-- Loading state -->
		<div class="space-y-4">
			<Skeleton height="60px" class="w-full" />
			<Skeleton height="400px" class="w-full" />
		</div>
	{:else if analysesService.error}
		<!-- Error state -->
		<div class="text-center py-20">
			<div class="w-20 h-20 mx-auto mb-4 bg-error/10 rounded-full flex items-center justify-center">
				<ClipboardList class="w-10 h-10 text-error" />
			</div>
			<h3 class="text-2xl font-bold mb-2">Failed to Load Analysis</h3>
			<p class="text-base-content/60 mb-6">{analysesService.error}</p>
			<Button variant="primary" onclick={() => window.location.reload()}>
				<RotateCw class="w-4 h-4" />
				Try Again
			</Button>
		</div>
	{:else if analysesService.currentAnalysis}
		{@const config = analysisTypeConfig[analysesService.currentAnalysis.type]}

		<!-- Header -->
		<div class="flex items-start justify-between gap-4 flex-wrap">
			<div class="flex items-center gap-4">
				<div class="w-16 h-16 bg-gradient-to-br {config.gradient} rounded-2xl flex items-center justify-center flex-shrink-0">
					<svelte:component this={config.icon} class="w-8 h-8 text-white" />
				</div>
				<div>
					<h1 class="text-3xl font-bold mb-1">{config.title}</h1>
					<Badge variant={config.badge as any}>
						AI-Generated Analysis
					</Badge>
				</div>
			</div>

			<div class="flex items-center gap-2">
				<Button variant="outline" onclick={handleExport}>
					{#snippet icon()}
						<Download class="w-4 h-4" />
					{/snippet}
					Export
				</Button>
			</div>
		</div>

		<!-- Divider -->
		<div class="divider"></div>

		<!-- Content -->
		<div class="pb-8">
			{#if analysesService.currentAnalysis.type === 'summary'}
				<SummaryView data={analysesService.currentAnalysis.result as any} />
			{:else if analysesService.currentAnalysis.type === 'flashcards'}
				<FlashcardsView data={analysesService.currentAnalysis.result as any} />
			{:else if analysesService.currentAnalysis.type === 'quiz'}
				<QuizView data={analysesService.currentAnalysis.result as any} />
			{:else if analysesService.currentAnalysis.type === 'deep_summary'}
				<DeepSummaryView data={analysesService.currentAnalysis.result as any} />
			{:else}
				<p class="text-center text-base-content/60 py-12">Unknown analysis type</p>
			{/if}
		</div>
	{:else}
		<div class="text-center py-20">
			<p class="text-base-content/60">Analysis not found</p>
		</div>
	{/if}
</div>
