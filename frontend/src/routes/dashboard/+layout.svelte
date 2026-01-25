<script lang="ts">
	import { goto } from "$app/navigation";
	import { getClient } from "$lib/apis/core.svelte";
	import { onMount } from "svelte";

    let { children } = $props();

	let courses: string[] = $state([]);
	let loading = $state(true);
	let loggedIn = $state(false);
	let error: string | null = null;

	onMount(async () => {
		loggedIn = document.cookie.split(";").some(cookie => cookie.trim().startsWith("auth="));

        const api = getClient();

        try {
            let d = await api.getCourses();
            console.log(d);
            courses.push(...d);
        }
        catch (e) {
            console.log(e);
        }
        loading = false;
	});

	function handleLogout() {
		// Remove the auth cookie by setting it expired
		document.cookie = "auth=; path=/; max-age=0; SameSite=Lax;";

		// Redirect to login page
		goto("/login");
	}
</script>

<div class="flex flex-col min-h-screen bg-base-200">
	<!-- Titlebar -->
	<header class="navbar bg-base-100 shadow-md px-6">
		<div class="flex-1 justify-center">
			<span class="text-xl font-bold">Examplify</span>
		</div>
		<div class="flex-none">
			{#if loggedIn}
				<span class="badge badge-primary">Logged in</span>
                <button class="btn btn-error" on:click={handleLogout}>
                    Log Out
                </button>
			{/if}
		</div>
	</header>

	<div class="flex flex-1">
		<!-- Sidebar / Navbar -->
		<aside class="w-60 bg-base-100 p-4 border-r border-base-300">
			<h3 class="font-semibold mb-2">Courses</h3>
			{#if loading}
				<p>Loading courses...</p>
			{:else if error}
				<p class="text-error">{error}</p>
			{:else if courses.length === 0}
				<p>No courses available.</p>
			{:else}
				<ul class="menu bg-base-100 w-full">
					{#each courses as course}
						<li>
							<a href={`/dashboard/${course}`} class="hover:bg-primary hover:text-primary-content">
								{course}
							</a>
						</li>
					{/each}
				</ul>
			{/if}
		</aside>

		<!-- Main content area -->
		<main class="flex-1 p-6">
			{@render children()}
		</main>
	</div>
</div>
