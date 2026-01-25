<script lang="ts">
	import { goto } from '$app/navigation';
	import { getClient } from '$lib/apis/core.svelte';
	import { onMount } from 'svelte';

	const { data } = $props<{
		data: {
			courseId: string;
            collectionID: string;
		};
	}>();

	const { courseId, collectionID } = data;

	// Data
	let files: { id: string; name: string; mimeType: string; url: string }[] = $state([]);
	let analyses: { id: string; type: string; result: string }[] = $state([]);

	// UI state
	let selectedFile:
		| { id: string; name: string; mimeType: string; url: string }
		| null = $state(null);

	let showFileModal = $state(false);

	// --- Fetch stubs (you implement) ---
	async function fetchFiles() {
		const api = getClient();

		const fs = await api.getCollectionDocuments({
			params: {
				id: collectionID
			}
		});

		files = fs.map((f: any) => ({
			id: f.ID,
			name: f.name ?? f.ID,       // Use name if available, fallback to ID
			mimeType: f.mimeType ?? "application/octet-stream",
			url: f.downloadURL          // presigned URL
		}));

		console.log("FF", fs);
	}

	async function fetchAnalyses() {
        const api = getClient();
		analyses = await api.getCollectionAnalyses({
            params: {
                id: collectionID
            }
        });
	}

	onMount(() => {
		fetchFiles();
		fetchAnalyses();
	});

	function openFile(file: { id: string; name: string; mimeType: string; url: string }) {
		selectedFile = file;
		showFileModal = true;
	}

	function closeModal() {
		showFileModal = false;
		selectedFile = null;
	}

	async function createAnalysis(type: 'summary' | 'flashcards' | 'quiz' | 'deep_summary') {
		console.log('create analysis', type);

        const api = getClient();

        const d = await api.analyzeCollection({
            type: type
        }, {
            params: {
                id: collectionID,
            }
        });

        await goto(`/dashboard/${courseId}/${collectionID}/${d.id}`)
	}
</script>

<div class="flex flex-col gap-6">

	<!-- Top section -->
	<div class="grid grid-cols-2 gap-6">

<!-- Uploaded files as grid thumbnails -->
<div class="card bg-base-100 shadow">
    <div class="card-body">
        <h2 class="card-title">Uploaded Files</h2>

        {#if files.length === 0}
            <p class="text-sm opacity-70">No files uploaded yet.</p>
        {:else}
            <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-4">
                {#each files as file}
                    <button
                        class="relative w-full aspect-square overflow-hidden rounded-lg bg-base-200 hover:scale-105 transition-transform"
                        on:click={() => openFile(file)}
                        title="Click to preview"
                    >
                        {#if file.mimeType.startsWith('image/')}
                            <img 
                                src={file.url} 
                                alt="Preview" 
                                class="w-full h-full object-cover"
                            />
                        {:else if file.mimeType === 'application/pdf'}
                            <div class="w-full h-full flex items-center justify-center bg-base-300 text-sm font-semibold">
                                PDF
                            </div>
                        {:else}
                            <div class="w-full h-full flex items-center justify-center bg-base-300 text-sm font-semibold">
                                FILE
                            </div>
                        {/if}
                    </button>
                {/each}
            </div>
        {/if}
    </div>
</div>


		<!-- Analyses recycler -->
		<div class="card bg-base-100 shadow">
			<div class="card-body">
				<h2 class="card-title">Analyses</h2>

				{#if analyses.length === 0}
					<p class="text-sm opacity-70">No analyses yet.</p>
				{:else}
					<div class="flex flex-col gap-2 max-h-80 overflow-y-auto">
						{#each analyses as analysis}
                            <a href="/dashboard/{courseId}/{collectionID}/{analysis.id}">
							    <div class="p-3 rounded bg-base-200 flex justify-between">
                                    <span class="capitalize">{analysis.type.replace('_', ' ')}</span>
                                </div>
                            </a>
						{/each}
					</div>
				{/if}
			</div>
		</div>

	</div>

	<!-- Actions -->
	<div class="card bg-base-100 shadow">
		<div class="card-body">
			<h2 class="card-title">Create New</h2>

			<div class="grid grid-cols-2 md:grid-cols-4 gap-3">
				<button class="btn btn-primary" on:click={() => createAnalysis('summary')}>
					New Summary
				</button>
				<button class="btn btn-secondary" on:click={() => createAnalysis('flashcards')}>
					New Flashcards
				</button>
				<button class="btn btn-accent" on:click={() => createAnalysis('quiz')}>
					New Quiz
				</button>
				<button class="btn btn-outline" on:click={() => createAnalysis('deep_summary')}>
					New Deep Summary
				</button>
			</div>
		</div>
	</div>

</div>

<!-- File Preview Modal -->
{#if showFileModal && selectedFile}
	<div class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
		<div class="bg-base-100 rounded-lg w-11/12 max-w-4xl p-4">
			<div class="flex justify-between items-center mb-2">
				<h3 class="font-bold">{selectedFile.name}</h3>
				<button class="btn btn-sm btn-ghost" on:click={closeModal}>✕</button>
			</div>

			<div class="max-h-[70vh] overflow-auto">
				{#if selectedFile.mimeType.startsWith('image/')}
					<img src={selectedFile.url} class="max-w-full mx-auto" />
				{:else if selectedFile.mimeType === 'application/pdf'}
					<iframe
						src={selectedFile.url}
						class="w-full h-[70vh]"
					/>
				{:else}
					<p>Preview not supported.</p>
				{/if}
			</div>
		</div>
	</div>
{/if}
