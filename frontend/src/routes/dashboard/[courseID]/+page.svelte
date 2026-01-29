<script lang="ts">
	import { collectionsService, documentsService } from '$lib/services';
	import { onMount } from 'svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import Modal from '$lib/components/ui/Modal.svelte';
	import Skeleton from '$lib/components/ui/Skeleton.svelte';
	import Badge from '$lib/components/ui/Badge.svelte';
	import {
		Plus,
		Upload,
		FileText,
		Image,
		File as FileIcon,
		X,
		BookOpen,
		GraduationCap,
		Folder,
		Camera
	} from 'lucide-svelte';
	import { fly } from 'svelte/transition';

	const { data } = $props<{
		data: {
			courseId: string;
		};
	}>();

	const { courseId } = data;

	let showUploadModal = $state(false);
	let files: File[] = $state([]);
	let selectedType: 'lecture' | 'exam' = $state('lecture');
	let collectionTitle = $state('');
	let dragOver = $state(false);
	let uploading = $state(false);

	onMount(async () => {
		await collectionsService.getByCourse(courseId);
	});

	function handleDrop(event: DragEvent) {
		event.preventDefault();
		dragOver = false;

		if (!event.dataTransfer) return;
		const droppedFiles = Array.from(event.dataTransfer.files);
		files = [...files, ...droppedFiles];
	}

	function handleDragOver(event: DragEvent) {
		event.preventDefault();
		dragOver = true;
	}

	function handleDragLeave() {
		dragOver = false;
	}

	function handleFileSelect(event: Event) {
		const target = event.target as HTMLInputElement;
		if (!target.files) return;

		files = [...files, ...Array.from(target.files)];
	}

	function handleCameraCapture(event: Event) {
		const target = event.target as HTMLInputElement;
		if (!target.files) return;

		files = [...files, ...Array.from(target.files)];
	}

	function openCamera() {
		document.getElementById('camera-input')?.click();
	}

	function removeFile(index: number) {
		files = files.filter((_, i) => i !== index);
	}

	function getFileIcon(file: File) {
		if (file.type.startsWith('image/')) return Image;
		if (file.type.includes('pdf')) return FileText;
		return FileIcon;
	}

	async function handleUpload() {
		if (!collectionTitle.trim() || files.length === 0) return;

		uploading = true;

		// Create collection
		const result = await collectionsService.create({
			title: collectionTitle,
			course: courseId,
			type: selectedType
		});

		if (result.data) {
			// Upload files
			await documentsService.uploadFiles(result.data.collectionID, files);

			// Reset form
			showUploadModal = false;
			collectionTitle = '';
			files = [];
			selectedType = 'lecture';

			// Refresh collections
			await collectionsService.getByCourse(courseId);
		}

		uploading = false;
	}

	function openModal() {
		showUploadModal = true;
	}

	function closeModal() {
		if (!uploading) {
			showUploadModal = false;
			collectionTitle = '';
			files = [];
			selectedType = 'lecture';
		}
	}
</script>

<div class="min-h-[calc(100vh-200px)]">
	<!-- Page header -->
	<div class="mb-8 flex items-center justify-between">
		<div>
			<h1 class="text-3xl font-bold mb-2">{courseId}</h1>
			<p class="text-base-content/60">Manage your study collections</p>
		</div>

		<Button variant="gradient" size="lg" onclick={openModal}>
			{#snippet icon()}
				<Plus class="w-5 h-5" />
			{/snippet}
			New Collection
		</Button>
	</div>

	<!-- Collections grid -->
	{#if collectionsService.loading}
		<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
			{#each Array(6) as _}
				<Skeleton height="200px" class="w-full" />
			{/each}
		</div>
	{:else if collectionsService.collections.length === 0}
		<!-- Empty state -->
		<div class="flex flex-col items-center justify-center py-20" transition:fly={{ y: 20, duration: 400 }}>
			<div class="w-32 h-32 gradient-primary rounded-3xl flex items-center justify-center mb-6">
				<Folder class="w-16 h-16 text-white" />
			</div>
			<h3 class="text-2xl font-bold mb-2">No Collections Yet</h3>
			<p class="text-base-content/60 mb-8 text-center max-w-md">
				Start by creating your first collection. Upload lecture notes or exam materials to get AI-powered study aids.
			</p>
			<Button variant="gradient" size="lg" onclick={openModal}>
				{#snippet icon()}
					<Plus class="w-5 h-5" />
				{/snippet}
				Create First Collection
			</Button>
		</div>
	{:else}
		<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
			{#each collectionsService.collections as collection (collection.ID)}
				<a href={`/dashboard/${courseId}/${collection.ID}`} transition:fly={{ y: 20, duration: 300 }}>
					<Card hover clickable gradient class="h-full">
						{#snippet header()}
							<div class="flex items-start justify-between">
								<div class="flex items-center gap-3">
									<div class="w-12 h-12 bg-gradient-to-br from-brand-500 to-accent-pink rounded-xl flex items-center justify-center flex-shrink-0">
										<BookOpen class="w-6 h-6 text-white" />
									</div>
									<div class="flex-1 min-w-0">
										<h3 class="font-bold text-lg truncate">{collection.title}</h3>
										<Badge variant={collection.type === 'lecture' ? 'primary' : 'secondary'} size="sm">
											{collection.type}
										</Badge>
									</div>
								</div>
							</div>
						{/snippet}

						<div class="mt-4 text-base-content/60 text-sm">
							<p>Click to view files and create analyses</p>
						</div>
					</Card>
				</a>
			{/each}
		</div>
	{/if}
</div>

<!-- Upload Modal -->
<Modal bind:open={showUploadModal} title="Create New Collection" size="lg" onclose={closeModal}>
	<div class="space-y-6">
		<!-- Collection details -->
		<Input
			type="text"
			label="Collection Title"
			bind:value={collectionTitle}
			placeholder="e.g., Week 3 Lecture Notes"
			required
		>
			{#snippet icon()}
				<FileText class="w-5 h-5" />
			{/snippet}
		</Input>

		<!-- Type selector -->
		<div class="form-control">
			<label class="label">
				<span class="label-text font-medium">Collection Type</span>
			</label>
			<div class="flex gap-4">
				<label class="flex-1">
					<input
						type="radio"
						name="type"
						value="lecture"
						bind:group={selectedType}
						class="radio radio-primary"
					/>
					<span class="ml-2">Lecture Notes</span>
				</label>
				<label class="flex-1">
					<input
						type="radio"
						name="type"
						value="exam"
						bind:group={selectedType}
						class="radio radio-primary"
					/>
					<span class="ml-2">Exam Materials</span>
				</label>
			</div>
		</div>

		<!-- Drag-drop area -->
		<div>
			<label class="label">
				<span class="label-text font-medium">Upload Files</span>
			</label>
			<div
				class={`border-2 border-dashed rounded-xl p-8 text-center transition-all cursor-pointer ${
					dragOver
						? 'border-primary bg-primary/10'
						: 'border-base-300 hover:border-primary/50 hover:bg-base-200/50'
				}`}
				ondrop={handleDrop}
				ondragover={handleDragOver}
				ondragleave={handleDragLeave}
				onclick={() => document.getElementById('file-input')?.click()}
			>
				<Upload class="w-12 h-12 mx-auto mb-4 text-base-content/40" />
				<p class="font-semibold mb-1">Drop files here or click to browse</p>
				<p class="text-sm text-base-content/60 mb-4">
					Supports images (PNG, JPG) and PDFs
				</p>
				<input
					id="file-input"
					type="file"
					multiple
					accept="image/*,application/pdf"
					class="hidden"
					onchange={handleFileSelect}
				/>
				<input
					id="camera-input"
					type="file"
					accept="image/*"
					capture="environment"
					class="hidden"
					onchange={handleCameraCapture}
				/>
			</div>

			<!-- Camera capture button (visible on mobile) -->
			<div class="mt-3">
				<Button
					variant="outline"
					size="md"
					class="w-full"
					onclick={(e) => { e.stopPropagation(); openCamera(); }}
					type="button"
				>
					{#snippet icon()}
						<Camera class="w-5 h-5" />
					{/snippet}
					Take Photo
				</Button>
			</div>
		</div>

		<!-- File list -->
		{#if files.length > 0}
			<div class="space-y-2">
				<label class="label">
					<span class="label-text font-medium">{files.length} file(s) selected</span>
				</label>
				<div class="space-y-2 max-h-60 overflow-y-auto">
					{#each files as file, index (file.name + index)}
						<div
							class="flex items-center gap-3 p-3 bg-base-200 rounded-lg"
							transition:fly={{ x: -20, duration: 200 }}
						>
							<svelte:component this={getFileIcon(file)} class="w-5 h-5 text-primary flex-shrink-0" />
							<div class="flex-1 min-w-0">
								<p class="font-medium text-sm truncate">{file.name}</p>
								<p class="text-xs text-base-content/60">
									{(file.size / 1024).toFixed(1)} KB
								</p>
							</div>
							<button
								class="btn btn-ghost btn-circle btn-sm"
								onclick={() => removeFile(index)}
								disabled={uploading}
							>
								<X class="w-4 h-4" />
							</button>
						</div>
					{/each}
				</div>
			</div>
		{/if}

		<!-- Upload progress -->
		{#if documentsService.uploadQueue.length > 0}
			<div class="space-y-2">
				<label class="label">
					<span class="label-text font-medium">Uploading...</span>
				</label>
				{#each documentsService.uploadQueue as item (item.file.name)}
					<div class="space-y-1">
						<div class="flex items-center justify-between text-sm">
							<span class="truncate">{item.file.name}</span>
							<span class="text-base-content/60">
								{item.status === 'completed' ? '100%' : `${item.progress}%`}
							</span>
						</div>
						<progress
							class="progress progress-primary w-full"
							value={item.progress}
							max="100"
						></progress>
					</div>
				{/each}
			</div>
		{/if}
	</div>

	{#snippet actions()}
		<Button variant="ghost" onclick={closeModal} disabled={uploading}>
			Cancel
		</Button>
		<Button
			variant="gradient"
			onclick={handleUpload}
			loading={uploading || collectionsService.loading}
			disabled={!collectionTitle.trim() || files.length === 0}
		>
			{uploading ? 'Uploading...' : 'Create Collection'}
		</Button>
	{/snippet}
</Modal>
