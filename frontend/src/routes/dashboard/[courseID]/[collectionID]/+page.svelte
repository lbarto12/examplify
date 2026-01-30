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
		ArrowLeft,
		ChevronRight,
		Plus
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

	// Group analyses by type
	let analysesByType = $derived(
		analysisTypes.map((typeConfig) => ({
			...typeConfig,
			analyses: collectionsService.analyses.filter((a) => a.type === typeConfig.type)
		}))
	);

	let hasAnyAnalyses = $derived(collectionsService.analyses.length > 0);

	let expandedType = $state<string | null>(null);
	let dropdownDirection = $state<'down' | 'up'>('down');

	function toggleExpand(type: string, event: MouseEvent) {
		if (expandedType === type) {
			expandedType = null;
			return;
		}

		// Calculate available space
		const button = event.currentTarget as HTMLElement;
		const rect = button.getBoundingClientRect();
		const spaceBelow = window.innerHeight - rect.bottom;
		const spaceAbove = rect.top;

		// If less than 200px below and more space above, open upward
		dropdownDirection = spaceBelow < 200 && spaceAbove > spaceBelow ? 'up' : 'down';
		expandedType = type;
	}

	function formatDate(dateString: string | undefined): string {
		if (!dateString) return '';
		const date = new Date(dateString);
		const now = new Date();
		const diffMs = now.getTime() - date.getTime();
		const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));

		if (diffDays === 0) {
			const diffHours = Math.floor(diffMs / (1000 * 60 * 60));
			if (diffHours === 0) {
				const diffMins = Math.floor(diffMs / (1000 * 60));
				return diffMins <= 1 ? 'Just now' : `${diffMins}m ago`;
			}
			return `${diffHours}h ago`;
		}
		if (diffDays === 1) return 'Yesterday';
		if (diffDays < 7) return `${diffDays}d ago`;

		return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
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

	<!-- Study Tools Section -->
	<div>
		<h2 class="text-xl font-bold mb-4 flex items-center gap-2">
			<Sparkles class="w-5 h-5 text-primary" />
			Study Tools
		</h2>

		<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
			{#each analysesByType as typeGroup}
				{@const hasAnalyses = typeGroup.analyses.length > 0}
				{@const isExpanded = expandedType === typeGroup.type}

				<div class="relative group">
					{#if hasAnalyses}
						<!-- Has existing analyses - show expandable card -->
						<div class="rounded-2xl overflow-hidden border border-base-300 bg-base-100 shadow-sm hover:shadow-md transition-shadow flex flex-col">
							<!-- Count badge - top left -->
							<span class="absolute -top-2 -left-2 w-6 h-6 bg-linear-to-br {typeGroup.gradient} rounded-full flex items-center justify-center text-xs font-bold text-white shadow-lg z-10">
								{typeGroup.analyses.length}
							</span>

							<!-- Card header -->
							<button
								class="w-full text-left p-4 pr-14 flex items-center gap-4 hover:bg-base-200/50 transition-colors"
								onclick={(e) => toggleExpand(typeGroup.type, e)}
							>
								<div class="w-12 h-12 bg-linear-to-br {typeGroup.gradient} rounded-xl flex items-center justify-center shrink-0 shadow-md">
									<svelte:component this={typeGroup.icon} class="w-6 h-6 text-white" />
								</div>
								<div class="flex-1 min-w-0">
									<h3 class="font-bold mb-0.5">{typeGroup.title}</h3>
									<p class="text-sm text-base-content/60">{typeGroup.description}</p>
								</div>
								<ChevronRight class="w-5 h-5 text-base-content/40 transition-transform duration-200 shrink-0 {isExpanded ? 'rotate-90' : ''}" />
							</button>

							<!-- Expanded list of analyses -->
							{#if isExpanded}
								<div
									class="bg-base-200/30 p-2 max-h-48 overflow-y-auto {dropdownDirection === 'up' ? 'order-first border-b' : 'border-t'} border-base-300"
								>
									{#each typeGroup.analyses as analysis, index}
										<a
											href={`/dashboard/${courseId}/${collectionID}/${analysis.id}`}
											class="flex items-center gap-3 px-3 py-2.5 rounded-lg hover:bg-base-300/50 transition-colors"
										>
											<span class="w-7 h-7 rounded-full bg-linear-to-br {typeGroup.gradient} flex items-center justify-center text-xs font-bold text-white shadow-sm shrink-0">
												{index + 1}
											</span>
											<span class="flex-1 font-medium text-sm">{typeGroup.title} #{index + 1}</span>
											<span class="text-xs text-base-content/50 shrink-0">{formatDate(analysis.createdAt)}</span>
											<ChevronRight class="w-4 h-4 text-base-content/40 shrink-0" />
										</a>
									{/each}
								</div>
							{/if}

							<!-- Create new button -->
							<button
								class="absolute top-4 right-4 w-7 h-7 bg-primary/10 text-primary rounded-lg flex items-center justify-center hover:bg-primary hover:text-white transition-all z-10"
								onclick={(e) => { e.stopPropagation(); createAnalysis(typeGroup.type); }}
								title="Create new {typeGroup.title}"
								disabled={creatingAnalysis === typeGroup.type}
							>
								{#if creatingAnalysis === typeGroup.type}
									<Loader class="w-4 h-4 animate-spin" />
								{:else}
									<Plus class="w-4 h-4" />
								{/if}
							</button>
						</div>
					{:else}
						<!-- No analyses yet - show create button -->
						<button
							class="w-full h-full"
							onclick={() => createAnalysis(typeGroup.type)}
							disabled={creatingAnalysis === typeGroup.type}
						>
							<div class="rounded-2xl border-2 border-dashed border-base-300 bg-base-100/50 p-4 flex items-center gap-4 hover:border-primary/50 hover:bg-base-200/30 transition-all">
								<div class="w-12 h-12 bg-linear-to-br {typeGroup.gradient} rounded-xl flex items-center justify-center shrink-0 opacity-50 group-hover:opacity-80 transition-opacity shadow-md">
									{#if creatingAnalysis === typeGroup.type}
										<Loader class="w-6 h-6 text-white animate-spin" />
									{:else}
										<svelte:component this={typeGroup.icon} class="w-6 h-6 text-white" />
									{/if}
								</div>
								<div class="flex-1 min-w-0 text-left">
									<h3 class="font-bold mb-0.5">{typeGroup.title}</h3>
									<p class="text-sm text-base-content/50">
										{#if creatingAnalysis === typeGroup.type}
											Creating your {typeGroup.title.toLowerCase()}...
										{:else}
											{typeGroup.description}
										{/if}
									</p>
								</div>
								<div class="w-10 h-10 rounded-full bg-base-200 flex items-center justify-center group-hover:bg-primary group-hover:text-white transition-colors">
									{#if creatingAnalysis === typeGroup.type}
										<Loader class="w-5 h-5 animate-spin" />
									{:else}
										<Plus class="w-5 h-5" />
									{/if}
								</div>
							</div>
						</button>
					{/if}
				</div>
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
