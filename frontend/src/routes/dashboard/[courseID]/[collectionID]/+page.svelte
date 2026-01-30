<script lang="ts">
	import { goto } from '$app/navigation';
	import { collectionsService, documentsService } from '$lib/services';
	import { onMount } from 'svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import Modal from '$lib/components/ui/Modal.svelte';
	import Badge from '$lib/components/ui/Badge.svelte';
	import {
		Image,
		FileText,
		File as FileIcon,
		Sparkles,
		Brain,
		ClipboardList,
		BookMarked,
		Loader,
		ZoomIn,
		Download,
		ArrowLeft
	} from 'lucide-svelte';
	import { fly } from 'svelte/transition';

	const { data } = $props<{
		data: {
			courseId: string;
			collectionID: string;
		};
	}>();

	const { courseId, collectionID } = data;

	let files: { id: string; name: string; mimeType: string; url: string; thumbnailUrl?: string }[] = $state([]);
	let selectedFile: { id: string; name: string; mimeType: string; url: string; thumbnailUrl?: string } | null = $state(null);
	let showFileModal = $state(false);
	let creatingAnalysis = $state<string | null>(null);
	let loaded = $state(false);

	const analysisTypes = [
		{
			type: 'summary' as const,
			title: 'Summary',
			description: 'Get a concise overview of all your materials',
			icon: FileText,
			gradient: 'from-purple-500 to-pink-500',
			color: 'primary'
		},
		{
			type: 'flashcards' as const,
			title: 'Flashcards',
			description: 'Interactive Q&A cards for studying',
			icon: Brain,
			gradient: 'from-teal-500 to-blue-500',
			color: 'secondary'
		},
		{
			type: 'quiz' as const,
			title: 'Quiz',
			description: 'Test your knowledge with multiple choice',
			icon: ClipboardList,
			gradient: 'from-orange-500 to-red-500',
			color: 'accent'
		},
		{
			type: 'deep_summary' as const,
			title: 'Deep Summary',
			description: 'Detailed breakdown with concepts and definitions',
			icon: BookMarked,
			gradient: 'from-indigo-500 to-purple-500',
			color: 'info'
		}
	];

	onMount(async () => {
		// Clear stale data before loading
		collectionsService.currentCollection = null;
		collectionsService.analyses = [];
		files = [];

		await collectionsService.getById(collectionID);
		await fetchFiles();
		await collectionsService.getAnalyses(collectionID);
		loaded = true;
	});

	async function fetchFiles() {
		const result = await documentsService.getByCollection(collectionID);
		if (result.data) {
			files = result.data.map((f: any) => ({
				id: f.ID,
				name: f.name ?? f.ID,
				mimeType: f.mimeType ?? 'application/octet-stream',
				url: f.downloadURL,
				thumbnailUrl: f.thumbnailURL
			}));
		}
	}

	function openFile(file: typeof files[0]) {
		selectedFile = file;
		showFileModal = true;
	}

	function closeModal() {
		showFileModal = false;
		selectedFile = null;
	}

	function getFileIcon(mimeType: string) {
		if (mimeType.startsWith('image/')) return Image;
		if (mimeType.includes('pdf')) return FileText;
		return FileIcon;
	}

	async function createAnalysis(type: 'summary' | 'flashcards' | 'quiz' | 'deep_summary') {
		creatingAnalysis = type;

		const result = await collectionsService.analyze(collectionID, type);

		if (result.data) {
			await goto(`/dashboard/${courseId}/${collectionID}/${result.data.id}`);
		}

		creatingAnalysis = null;
	}
</script>

{#if !loaded}
	<!-- Empty placeholder while loading -->
	<div class="min-h-64"></div>
{:else}
	<div class="space-y-8" in:fly={{ y: 15, duration: 250 }}>
		<!-- Page header -->
		<div class="flex items-center gap-4">
			<a href={`/dashboard/${courseId}`} class="btn btn-ghost btn-circle">
				<ArrowLeft class="w-5 h-5" />
			</a>
			<div>
				<h1 class="text-3xl font-bold mb-1">
					{collectionsService.currentCollection?.title || 'Collection'}
				</h1>
				<div class="flex items-center gap-2">
					<Badge
						variant={collectionsService.currentCollection?.type === 'lecture' ? 'primary' : 'secondary'}
					>
						{collectionsService.currentCollection?.type || ''}
					</Badge>
					<span class="text-base-content/60">•</span>
					<span class="text-base-content/60">{files.length} file(s)</span>
				</div>
			</div>
		</div>

		<!-- File Gallery -->
		<div>
			<h2 class="text-xl font-bold mb-4 flex items-center gap-2">
				<Image class="w-5 h-5 text-primary" />
				Uploaded Files
			</h2>

			{#if files.length === 0}
			<Card class="text-center py-12">
				<div class="w-20 h-20 mx-auto mb-4 bg-base-200 rounded-full flex items-center justify-center">
					<FileIcon class="w-10 h-10 text-base-content/40" />
				</div>
				<p class="text-base-content/60">No files uploaded yet</p>
			</Card>
		{:else}
			<div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-4">
				{#each files as file (file.id)}
					<button
						class="group relative aspect-square overflow-hidden rounded-xl bg-base-200 transition-all hover:scale-105 hover:shadow-xl"
						onclick={() => openFile(file)}
					>
						{#if file.mimeType.startsWith('image/')}
							<img
								src={file.thumbnailUrl || file.url}
								alt={file.name}
								class="w-full h-full object-cover"
								loading="lazy"
								decoding="async"
							/>
						{:else}
							<div class="w-full h-full flex flex-col items-center justify-center bg-gradient-to-br from-brand-500 to-accent-pink p-4">
								<svelte:component this={getFileIcon(file.mimeType)} class="w-12 h-12 text-white mb-2" />
								<span class="text-white text-xs font-semibold text-center line-clamp-2">
									{file.name}
								</span>
							</div>
						{/if}

						<!-- Hover overlay -->
						<div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
							<ZoomIn class="w-8 h-8 text-white" />
						</div>
					</button>
				{/each}
			</div>
		{/if}
	</div>

	<!-- Analyses Section -->
	<div>
		<h2 class="text-xl font-bold mb-4 flex items-center gap-2">
			<Sparkles class="w-5 h-5 text-primary" />
			Analyses
		</h2>

		{#if collectionsService.analyses.length === 0}
			<Card class="text-center py-8">
				<Brain class="w-16 h-16 mx-auto mb-3 text-base-content/40" />
				<p class="text-base-content/60 mb-1">No analyses created yet</p>
				<p class="text-sm text-base-content/40">Create your first analysis below</p>
			</Card>
		{:else}
			<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
				{#each collectionsService.analyses as analysis (analysis.id)}
					{@const typeConfig = analysisTypes.find((t) => t.type === analysis.type)}
					<a href={`/dashboard/${courseId}/${collectionID}/${analysis.id}`}>
						<Card hover clickable class="h-full">
							<div class="flex items-center gap-4">
								{#if typeConfig}
									<div class="w-12 h-12 bg-gradient-to-br {typeConfig.gradient} rounded-xl flex items-center justify-center flex-shrink-0">
										<svelte:component this={typeConfig.icon} class="w-6 h-6 text-white" />
									</div>
									<div class="flex-1 min-w-0">
										<h3 class="font-semibold">{typeConfig.title}</h3>
										<p class="text-sm text-base-content/60">Click to view</p>
									</div>
								{/if}
							</div>
						</Card>
					</a>
				{/each}
			</div>
		{/if}
	</div>

	<!-- Create Analysis -->
	<div>
		<h2 class="text-xl font-bold mb-4">Create New Analysis</h2>

		<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
			{#each analysisTypes as analysisType}
				<Card
					hover
					clickable
					gradient
					class="h-full"
					onclick={() => createAnalysis(analysisType.type)}
				>
					<div class="text-center space-y-4">
						<div class="w-16 h-16 mx-auto bg-gradient-to-br {analysisType.gradient} rounded-2xl flex items-center justify-center">
							{#if creatingAnalysis === analysisType.type}
								<Loader class="w-8 h-8 text-white animate-spin" />
							{:else}
								<svelte:component this={analysisType.icon} class="w-8 h-8 text-white" />
							{/if}
						</div>

						<div>
							<h3 class="font-bold text-lg mb-1">{analysisType.title}</h3>
							<p class="text-sm text-base-content/60">{analysisType.description}</p>
						</div>

						{#if creatingAnalysis === analysisType.type}
							<Badge variant="primary">
								<Loader class="w-3 h-3 mr-1 animate-spin" />
								Creating...
							</Badge>
						{:else}
							<div class="text-xs text-primary font-semibold">Click to create</div>
						{/if}
					</div>
				</Card>
			{/each}
		</div>
	</div>
	</div>
{/if}

<!-- File Preview Modal -->
<Modal bind:open={showFileModal} title={selectedFile?.name || 'File Preview'} size="xl" onclose={closeModal}>
	{#if selectedFile}
		<div class="max-h-[70vh] overflow-auto">
			{#if selectedFile.mimeType.startsWith('image/')}
				<img src={selectedFile.url} alt={selectedFile.name} class="max-w-full mx-auto rounded-lg" />
			{:else if selectedFile.mimeType.includes('pdf')}
				<iframe src={selectedFile.url} class="w-full h-[70vh] rounded-lg" title="PDF Preview"></iframe>
			{:else}
				<div class="text-center py-12">
					<FileIcon class="w-16 h-16 mx-auto mb-4 text-base-content/40" />
					<p class="text-base-content/60 mb-4">Preview not available for this file type</p>
					<a href={selectedFile.url} download class="btn btn-primary btn-sm">
						<Download class="w-4 h-4" />
						Download File
					</a>
				</div>
			{/if}
		</div>
	{/if}

	{#snippet actions()}
		{#if selectedFile}
			<a href={selectedFile.url} download class="btn btn-primary">
				<Download class="w-4 h-4" />
				Download
			</a>
		{/if}
		<Button variant="ghost" onclick={closeModal}>
			Close
		</Button>
	{/snippet}
</Modal>
