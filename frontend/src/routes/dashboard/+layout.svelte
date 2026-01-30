<script lang="ts">
	import { authService, coursesService, toastStore } from '$lib/services';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import Breadcrumbs from '$lib/components/Breadcrumbs.svelte';
	import ThemeToggle from '$lib/components/ui/ThemeToggle.svelte';
	import Skeleton from '$lib/components/ui/Skeleton.svelte';
	import Modal from '$lib/components/ui/Modal.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import {
		BookOpen,
		Menu,
		X,
		LogOut,
		User,
		Search,
		Plus,
		GraduationCap
	} from 'lucide-svelte';
	import { fly, slide, fade } from 'svelte/transition';
	import { cubicOut } from 'svelte/easing';

	let { children } = $props();

	let sidebarOpen = $state(false);
	let searchQuery = $state('');
	let createCourseModalOpen = $state(false);
	let newCourseName = $state('');
	let courseNameError = $state('');

	// Course colors for visual distinction
	const courseColors = [
		'from-purple-500 to-pink-500',
		'from-teal-500 to-blue-500',
		'from-orange-500 to-red-500',
		'from-green-500 to-emerald-500',
		'from-indigo-500 to-purple-500',
		'from-yellow-500 to-orange-500',
		'from-pink-500 to-rose-500',
		'from-cyan-500 to-teal-500'
	];

	function getCourseColor(index: number): string {
		return courseColors[index % courseColors.length];
	}

	let filteredCourses = $derived(
		searchQuery
			? coursesService.courses.filter((course) =>
					course.toLowerCase().includes(searchQuery.toLowerCase())
				)
			: coursesService.courses
	);

	onMount(async () => {
		// Check auth
		if (!authService.checkAuth()) {
			authService.signOut();
			return;
		}

		// Load courses
		await coursesService.getCourses();
	});

	function toggleSidebar() {
		sidebarOpen = !sidebarOpen;
	}

	function closeSidebar() {
		sidebarOpen = false;
	}

	function openCreateCourseModal() {
		createCourseModalOpen = true;
		newCourseName = '';
		courseNameError = '';
	}

	function closeCreateCourseModal() {
		createCourseModalOpen = false;
		newCourseName = '';
		courseNameError = '';
	}

	async function handleCreateCourse(e: Event) {
		e.preventDefault();
		courseNameError = '';

		if (!newCourseName.trim()) {
			courseNameError = 'Course name is required';
			return;
		}

		const result = await coursesService.createCourse(newCourseName.trim());

		if (result.error) {
			courseNameError = result.error.message;
			toastStore.error(result.error.message);
		} else {
			toastStore.success(`Course "${newCourseName}" created successfully!`);
			closeCreateCourseModal();
		}
	}
</script>

<div class="min-h-screen bg-base-200/50">
	<!-- Header -->
	<header class="sticky top-0 z-30 bg-base-100 border-b border-base-300 shadow-sm">
		<div class="navbar px-4 lg:px-6">
			<!-- Mobile menu button -->
			<div class="lg:hidden">
				<button class="btn btn-ghost btn-circle" onclick={toggleSidebar}>
					<Menu class="w-6 h-6" />
				</button>
			</div>

			<!-- Logo -->
			<div class="flex-1 lg:flex-none">
				<a href="/dashboard" class="flex items-center gap-2 group">
					<div
						class="w-10 h-10 gradient-primary rounded-xl flex items-center justify-center group-hover:scale-105 transition-transform"
					>
						<BookOpen class="w-6 h-6 text-white" />
					</div>
					<span class="text-xl font-bold hidden lg:inline">Examplify</span>
				</a>
			</div>

			<!-- Breadcrumbs (Desktop) -->
			<div class="hidden lg:flex flex-1 ml-8">
				<Breadcrumbs />
			</div>

			<!-- Right side actions -->
			<div class="flex-none flex items-center gap-2">
				<ThemeToggle />

				<!-- User menu -->
				<div class="dropdown dropdown-end">
					<button class="btn btn-ghost btn-circle">
						<User class="w-5 h-5" />
					</button>
					<ul
						class="dropdown-content z-[1] menu p-2 shadow-lg bg-base-100 rounded-box w-52 mt-3 border border-base-300"
					>
						<li class="menu-title">Account</li>
						<li>
							<button onclick={() => authService.signOut()}>
								<LogOut class="w-4 h-4" />
								Sign Out
							</button>
						</li>
					</ul>
				</div>
			</div>
		</div>

		<!-- Mobile breadcrumbs -->
		<div class="lg:hidden px-4 pb-3">
			<Breadcrumbs />
		</div>
	</header>

	<div class="flex">
		<!-- Sidebar (Desktop) -->
		<aside class="hidden lg:block w-72 bg-base-100 border-r border-base-300 min-h-[calc(100vh-65px)] p-6">
			<div class="mb-6">
				<h2 class="text-lg font-bold mb-2 flex items-center gap-2">
					<GraduationCap class="w-5 h-5 text-primary" />
					My Courses
				</h2>

				<!-- Search -->
				<div class="relative mb-3">
					<Search class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-base-content/50" />
					<input
						type="text"
						placeholder="Search courses..."
						bind:value={searchQuery}
						class="input input-bordered input-sm w-full pl-10"
					/>
				</div>

				<!-- Create Course Button -->
				<button
					onclick={openCreateCourseModal}
					class="btn btn-sm gradient-primary text-white w-full"
				>
					<Plus class="w-4 h-4" />
					Create Course
				</button>
			</div>

			<!-- Course list -->
			<div class="space-y-3">
				{#if coursesService.loading}
					{#each Array(5) as _}
						<Skeleton height="80px" class="w-full" />
					{/each}
				{:else if filteredCourses.length === 0}
					<div class="text-center py-8 text-base-content/60">
						<p class="text-sm">
							{searchQuery ? 'No courses found' : 'No courses yet'}
						</p>
					</div>
				{:else}
					{#each filteredCourses as course, i}
						<a
							href={`/dashboard/${course}`}
							class="block group"
							transition:slide={{ duration: 200 }}
						>
							<div
								class="relative overflow-hidden rounded-xl bg-gradient-to-br {getCourseColor(
									i
								)} p-[2px] card-hover"
							>
								<div class="bg-base-100 rounded-xl p-4 h-full">
									<div class="flex items-center gap-3">
										<div class="w-10 h-10 bg-white/10 rounded-lg flex items-center justify-center flex-shrink-0">
											<BookOpen class="w-5 h-5 text-{getCourseColor(i).split('-')[1]}-500" />
										</div>
										<div class="flex-1 min-w-0">
											<h3 class="font-semibold text-sm truncate">{course}</h3>
											<p class="text-xs text-base-content/60">View collections</p>
										</div>
									</div>
								</div>
							</div>
						</a>
					{/each}
				{/if}
			</div>
		</aside>

		<!-- Mobile Sidebar Drawer -->
		{#if sidebarOpen}
			<div
				class="fixed inset-0 z-40 lg:hidden"
				transition:fly={{ x: -300, duration: 300 }}
			>
				<!-- Backdrop -->
				<div
					class="absolute inset-0 bg-black/50 backdrop-blur-sm"
					onclick={closeSidebar}
				></div>

				<!-- Drawer -->
				<div class="relative w-80 max-w-[80vw] h-full bg-base-100 p-6 overflow-y-auto">
					<div class="flex items-center justify-between mb-6">
						<h2 class="text-lg font-bold flex items-center gap-2">
							<GraduationCap class="w-5 h-5 text-primary" />
							My Courses
						</h2>
						<button class="btn btn-ghost btn-circle btn-sm" onclick={closeSidebar}>
							<X class="w-5 h-5" />
						</button>
					</div>

					<!-- Search -->
					<div class="relative mb-3">
						<Search
							class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-base-content/50"
						/>
						<input
							type="text"
							placeholder="Search courses..."
							bind:value={searchQuery}
							class="input input-bordered input-sm w-full pl-10"
						/>
					</div>

					<!-- Create Course Button -->
					<button
						onclick={openCreateCourseModal}
						class="btn btn-sm gradient-primary text-white w-full mb-6"
					>
						<Plus class="w-4 h-4" />
						Create Course
					</button>

					<!-- Course list -->
					<div class="space-y-3">
						{#if coursesService.loading}
							{#each Array(5) as _}
								<Skeleton height="80px" class="w-full" />
							{/each}
						{:else if filteredCourses.length === 0}
							<div class="text-center py-8 text-base-content/60">
								<p class="text-sm">
									{searchQuery ? 'No courses found' : 'No courses yet'}
								</p>
							</div>
						{:else}
							{#each filteredCourses as course, i}
								<a
									href={`/dashboard/${course}`}
									class="block"
									onclick={closeSidebar}
								>
									<div
										class="relative overflow-hidden rounded-xl bg-gradient-to-br {getCourseColor(
											i
										)} p-[2px]"
									>
										<div class="bg-base-100 rounded-xl p-4 h-full">
											<div class="flex items-center gap-3">
												<div
													class="w-10 h-10 bg-white/10 rounded-lg flex items-center justify-center flex-shrink-0"
												>
													<BookOpen class="w-5 h-5" />
												</div>
												<div class="flex-1 min-w-0">
													<h3 class="font-semibold text-sm truncate">{course}</h3>
													<p class="text-xs text-base-content/60">View collections</p>
												</div>
											</div>
										</div>
									</div>
								</a>
							{/each}
						{/if}
					</div>
				</div>
			</div>
		{/if}

		<!-- Main content -->
		<main class="flex-1 p-6 lg:p-8 max-w-7xl mx-auto w-full">
			{@render children()}
		</main>
	</div>

	<!-- Create Course Modal -->
	<Modal bind:open={createCourseModalOpen} title="Create New Course" onclose={closeCreateCourseModal}>
		{#snippet children()}
			<form onsubmit={handleCreateCourse} class="space-y-4">
				<Input
					type="text"
					label="Course Name"
					bind:value={newCourseName}
					error={courseNameError}
					placeholder="e.g., Computer Science 101"
					required
				>
					{#snippet icon()}
						<GraduationCap class="w-5 h-5" />
					{/snippet}
				</Input>

				<p class="text-sm text-base-content/60">
					Create a course to organize your study materials and collections.
				</p>
			</form>
		{/snippet}

		{#snippet actions()}
			<Button variant="ghost" onclick={closeCreateCourseModal}>Cancel</Button>
			<Button
				type="submit"
				variant="gradient"
				loading={coursesService.loading}
				onclick={handleCreateCourse}
			>
				{#snippet icon()}
					{#if !coursesService.loading}
						<Plus class="w-5 h-5" />
					{/if}
				{/snippet}
				Create Course
			</Button>
		{/snippet}
	</Modal>
</div>

<style>
	.navbar {
		height: 65px;
	}
</style>
