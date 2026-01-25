<script lang="ts">
	import { getClient } from '$lib/apis/core.svelte';
	import { onMount } from 'svelte';

	const { data } = $props<{
		data: {
			courseId: string;
		};
	}>();

	const { courseId } = data;

    console.log(courseId);


	// Collections (left column)
	let collections: {
        ID: string;
        title: string;
        course: string;
        type: string;
    }[] = $state([]);
	// Modal state
	let showModal = $state(false);

	// Upload state
	let files: File[] = $state([]);
	let selectedType: 'lecture' | 'exam' = $state('lecture');
    let selectedCourse: string = $state('');
    let selectedTitle: string = $state('');

	// Placeholder function for fetching collections
	async function fetchCollections() {
        const api = getClient();

        try {
            console.log(courseId);
            const d = await api.getCourseCollections({
                params: {
                    courseID: courseId
                }
            });
            console.log(d);
            collections = d;
        }
        catch (e) {
            console.log(e);
        }
	}

	onMount(() => {
		fetchCollections();
	});

	// Drag and drop handlers
	function handleDrop(event: DragEvent) {
		event.preventDefault();
		if (!event.dataTransfer) return;
		const droppedFiles = Array.from(event.dataTransfer.files);
		files = [...files, ...droppedFiles];
	}

	function handleDragOver(event: DragEvent) {
		event.preventDefault();
	}

	// Upload function (empty for now)
	async function uploadFiles() {
        const api = getClient();
		console.log('Upload files:', files, 'Type:', selectedType);

        const d = await api.newCollection({
            course: selectedCourse,
            title: selectedTitle,
            type: selectedType
        });

        await Promise.all(files.map(async(f) => {
            const { uploadURL } = await api.uploadFile({
                collectionID: d.collectionID,
                mimeType: f.type,
            });

			// PUT file to MinIO
			const res = await fetch(uploadURL, {
				method: "PUT",
				headers: {
					"Content-Type": f.type
				},
				body: f
			});


			if (!res.ok) {
				throw new Error(`Failed to upload ${f.name}`);
			}

        }));

		// Your upload logic here
	}
</script>

<div class="flex min-h-screen bg-base-200 p-6 space-x-6">
	<!-- Left column: collections -->
	<div class="w-1/3 bg-base-100 p-4 rounded-lg shadow-md">
		<h2 class="text-lg font-bold mb-4">Collections</h2>
		{#if collections.length === 0}
			<p>No collections yet.</p>
		{:else}
			<ul class="menu bg-base-100 w-full">
				{#each collections as collection}
					<li>
						<a href="/dashboard/{courseId}/{collection.ID}" class="hover:bg-primary hover:text-primary-content">{collection.title}</a>
					</li>
				{/each}
			</ul>
		{/if}
	</div>

	<!-- Right column: add collection -->
	<div class="flex-1 bg-base-100 p-4 rounded-lg shadow-md flex flex-col items-center justify-center">
		<button class="btn btn-primary" on:click={() => (showModal = true)}>
			Add Collection
		</button>
	</div>
</div>

<!-- Modal -->
{#if showModal}
	<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
		<div class="bg-base-100 rounded-lg p-6 w-2/3 max-w-2xl shadow-lg">
			<h2 class="text-xl font-bold mb-4">Upload Files</h2>

			<!-- Dropdown -->
			<div class="mb-4">
				<label class="label">
					<span class="label-text">Collection Type</span>
				</label>
				<select bind:value={selectedType} class="select select-bordered w-full">
					<option value="lecture">lecture notes</option>
					<option value="exam">exam</option>
				</select>
			</div>

            			<!-- Dropdown -->
            <div class="mb-4">
                <label class="label">
                    <span class="label-text">Course Code</span>
                </label> 
                <input bind:value={selectedCourse} class="input" type="text"/>
            </div>

            <div class="mb-4">
                <label class="label">
                    <span class="label-text">Collection Title</span>
                </label> 
                <input bind:value={selectedTitle} class="input" type="text"/>
            </div>

			<!-- Drag & drop area -->
			<div
				class="border-2 border-dashed border-gray-400 rounded-lg p-6 text-center mb-4 cursor-pointer"
				on:drop={handleDrop}
				on:dragover={handleDragOver}
			>
				<p class="text-gray-500 mb-2">Drag & drop images or PDFs here</p>
				<p class="text-sm text-gray-400">or click to select files</p>
				<input
					type="file"
					multiple
					class="hidden"
					on:change={(e) => {
						const target = e.target as HTMLInputElement;
						if (!target.files) return;
						files = [...files, ...Array.from(target.files)];
					}}
				/>
			</div>

			<!-- File list -->
			{#if files.length > 0}
				<ul class="mb-4">
					{#each files as file, index}
						<li class="flex justify-between items-center py-1 border-b border-base-300">
							<span>{file.name}</span>
							<button
								class="btn btn-sm btn-ghost text-error"
								on:click={() => files.splice(index, 1)}
							>
								Remove
							</button>
						</li>
					{/each}
				</ul>
			{/if}

			<!-- Buttons -->
			<div class="flex justify-end space-x-2">
				<button class="btn btn-ghost" on:click={() => (showModal = false)}>Cancel</button>
				<button class="btn btn-primary" on:click={uploadFiles}>Submit</button>
			</div>
		</div>
	</div>
{/if}
