<script lang="ts">
	import { goto } from '$app/navigation';
	import { getClient } from '$lib/apis/sessions.svelte';

	let email = $state("");
	let password = $state("");
	let error: string | null = null;
	let loading = false;

	async function handleLogin() {
		error = null;
		loading = true;
		const sessionAPI = getClient();

		try {
			const { token } = await sessionAPI.signIn({ 
				email: email,
				password: password,
			});

			await fetch("/actions/set-cookie", {
				method: "POST",
				body: JSON.stringify({
					key: "auth",
					value: token,
				}),
			});

			await goto("/dashboard");
		} catch (e: any) {
			error = e?.message || "Failed to sign in";
		} finally {
			loading = false;
		}
	}
</script>

<div class="min-h-screen flex items-center justify-center bg-base-200">
	<div class="card w-full max-w-sm shadow-xl bg-base-100 p-6">
		<h2 class="text-2xl font-bold text-center mb-6">Sign In</h2>

		{#if error}
			<div class="alert alert-error mb-4">
				<span>{error}</span>
			</div>
		{/if}

		<form class="form-control space-y-4" on:submit|preventDefault={handleLogin}>
			<label class="label">
				<span class="label-text">Email</span>
			</label>
			<input
				type="email"
				bind:value={email}
				placeholder="you@example.com"
				class="input input-bordered w-full"
				required
			/>

			<label class="label">
				<span class="label-text">Password</span>
			</label>
			<input
				type="password"
				bind:value={password}
				placeholder="••••••••"
				class="input input-bordered w-full"
				required
			/>

			<label class="label justify-end">
				<a href="/forgot-password" class="link link-hover text-sm">Forgot password?</a>
			</label>

			<button type="submit" class="btn btn-primary w-full" disabled={loading}>
				{#if loading}
					<span class="loading loading-spinner loading-sm"></span> Signing In...
				{:else}
					Sign In
				{/if}
			</button>
		</form>

		<p class="text-center mt-4 text-sm">
			New here? <a href="/signup" class="link link-primary">Create an account</a>
		</p>
	</div>
</div>
