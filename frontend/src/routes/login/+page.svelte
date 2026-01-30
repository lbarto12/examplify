<script lang="ts">
	import { goto } from '$app/navigation';
	import { authService } from '$lib/services';
	import Input from '$lib/components/ui/Input.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import { Mail, Lock, BookOpen, Sparkles, Brain, GraduationCap } from 'lucide-svelte';
	import { fly, fade } from 'svelte/transition';

	let email = $state('');
	let password = $state('');
	let emailError = $state('');
	let passwordError = $state('');

	async function handleLogin(e: Event) {
		e.preventDefault();

		// Reset errors
		emailError = '';
		passwordError = '';

		// Basic validation
		if (!email) {
			emailError = 'Email is required';
			return;
		}
		if (!password) {
			passwordError = 'Password is required';
			return;
		}

		const result = await authService.signIn(email, password);

		if (!result.error) {
			// Navigate to dashboard on success
			await goto('/dashboard');
		}
	}
</script>

<div class="min-h-screen grid lg:grid-cols-2">
	<!-- Left Side - Hero / Branding -->
	<div class="hidden lg:flex flex-col justify-center items-center p-12 gradient-primary relative overflow-hidden">
		<!-- Animated background pattern -->
		<div class="absolute inset-0 opacity-10">
			<div class="absolute top-0 left-0 w-96 h-96 bg-white rounded-full blur-3xl animate-pulse"></div>
			<div class="absolute bottom-0 right-0 w-96 h-96 bg-white rounded-full blur-3xl animate-pulse" style="animation-delay: 1s;"></div>
		</div>

		<div class="relative z-10 text-white text-center max-w-md" in:fade={{ duration: 400 }}>
			<div class="flex justify-center mb-8">
				<div class="w-20 h-20 bg-white/20 rounded-3xl flex items-center justify-center backdrop-blur-sm">
					<BookOpen class="w-12 h-12" />
				</div>
			</div>

			<h1 class="text-5xl font-bold mb-4">Examplify</h1>
			<p class="text-xl mb-12 text-white/90">Transform your study materials into interactive learning experiences</p>

			<div class="space-y-6 text-left">
				<div class="flex items-start gap-4" in:fly={{ y: 20, delay: 100, duration: 400 }}>
					<div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center flex-shrink-0 backdrop-blur-sm">
						<Sparkles class="w-6 h-6" />
					</div>
					<div>
						<h3 class="font-semibold text-lg mb-1">AI-Powered Analysis</h3>
						<p class="text-white/80 text-sm">Generate summaries, flashcards, and quizzes from your documents</p>
					</div>
				</div>

				<div class="flex items-start gap-4" in:fly={{ y: 20, delay: 200, duration: 400 }}>
					<div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center flex-shrink-0 backdrop-blur-sm">
						<Brain class="w-6 h-6" />
					</div>
					<div>
						<h3 class="font-semibold text-lg mb-1">Smart Organization</h3>
						<p class="text-white/80 text-sm">Organize by course and access your materials anywhere</p>
					</div>
				</div>

				<div class="flex items-start gap-4" in:fly={{ y: 20, delay: 300, duration: 400 }}>
					<div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center flex-shrink-0 backdrop-blur-sm">
						<GraduationCap class="w-6 h-6" />
					</div>
					<div>
						<h3 class="font-semibold text-lg mb-1">Study Smarter</h3>
						<p class="text-white/80 text-sm">Interactive quizzes and flashcards for effective learning</p>
					</div>
				</div>
			</div>
		</div>
	</div>

	<!-- Right Side - Login Form -->
	<div class="flex flex-col justify-center items-center p-8 lg:p-12 bg-base-100">
		<div class="w-full max-w-md" in:fly={{ x: 20, duration: 400 }}>
			<!-- Mobile logo -->
			<div class="lg:hidden flex justify-center mb-8">
				<div class="w-16 h-16 gradient-primary rounded-2xl flex items-center justify-center">
					<BookOpen class="w-10 h-10 text-white" />
				</div>
			</div>

			<div class="mb-8">
				<h2 class="text-4xl font-bold mb-2">Welcome back!</h2>
				<p class="text-base-content/60">Sign in to continue your learning journey</p>
			</div>

			<form class="space-y-6" onsubmit={handleLogin}>
				<Input
					type="email"
					label="Email"
					bind:value={email}
					error={emailError}
					required
					autocomplete="email"
				>
					{#snippet icon()}
						<Mail class="w-5 h-5" />
					{/snippet}
				</Input>

				<Input
					type="password"
					label="Password"
					bind:value={password}
					error={passwordError}
					required
					autocomplete="current-password"
				>
					{#snippet icon()}
						<Lock class="w-5 h-5" />
					{/snippet}
				</Input>

				<div class="flex items-center justify-between">
					<label class="flex items-center gap-2 cursor-pointer">
						<input type="checkbox" class="checkbox checkbox-primary checkbox-sm" />
						<span class="text-sm">Remember me</span>
					</label>
					<a href="/forgot-password" class="text-sm text-primary hover:underline">
						Forgot password?
					</a>
				</div>

				<Button
					type="submit"
					variant="gradient"
					size="lg"
					loading={authService.loading}
					class="w-full"
				>
					{authService.loading ? 'Signing in...' : 'Sign In'}
				</Button>
			</form>

			<div class="divider my-8">OR</div>

			<div class="text-center">
				<p class="text-base-content/60">
					Don't have an account?
					<a href="/signup" class="text-primary font-semibold hover:underline ml-1">
						Sign up for free
					</a>
				</p>
			</div>
		</div>
	</div>
</div>

<style>
	@keyframes pulse {
		0%, 100% {
			opacity: 0.1;
		}
		50% {
			opacity: 0.2;
		}
	}
</style>
