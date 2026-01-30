<script lang="ts">
	import { goto } from '$app/navigation';
	import { authService } from '$lib/services';
	import Input from '$lib/components/ui/Input.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import { Mail, Lock, BookOpen, Sparkles, Brain, GraduationCap, UserPlus } from 'lucide-svelte';
	import { fly, fade } from 'svelte/transition';

	let email = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	let emailError = $state('');
	let passwordError = $state('');
	let confirmPasswordError = $state('');

	async function handleSignup(e: Event) {
		e.preventDefault();

		// Reset errors
		emailError = '';
		passwordError = '';
		confirmPasswordError = '';

		// Basic validation
		if (!email) {
			emailError = 'Email is required';
			return;
		}

		// Email format validation
		const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
		if (!emailRegex.test(email)) {
			emailError = 'Please enter a valid email address';
			return;
		}

		if (!password) {
			passwordError = 'Password is required';
			return;
		}

		if (password.length < 8) {
			passwordError = 'Password must be at least 8 characters';
			return;
		}

		if (!confirmPassword) {
			confirmPasswordError = 'Please confirm your password';
			return;
		}

		if (password !== confirmPassword) {
			confirmPasswordError = 'Passwords do not match';
			return;
		}

		const result = await authService.signUp(email, password);

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

			<h1 class="text-5xl font-bold mb-4">Join Examplify</h1>
			<p class="text-xl mb-12 text-white/90">Start transforming your study materials today</p>

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

	<!-- Right Side - Signup Form -->
	<div class="flex flex-col justify-center items-center p-8 lg:p-12 bg-base-100">
		<div class="w-full max-w-md" in:fly={{ x: 20, duration: 400 }}>
			<!-- Mobile logo -->
			<div class="lg:hidden flex justify-center mb-8">
				<div class="w-16 h-16 gradient-primary rounded-2xl flex items-center justify-center">
					<BookOpen class="w-10 h-10 text-white" />
				</div>
			</div>

			<div class="mb-8">
				<h2 class="text-4xl font-bold mb-2">Create Account</h2>
				<p class="text-base-content/60">Sign up to start your learning journey</p>
			</div>

			<form class="space-y-6" onsubmit={handleSignup}>
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
					autocomplete="new-password"
				>
					{#snippet icon()}
						<Lock class="w-5 h-5" />
					{/snippet}
				</Input>

				<Input
					type="password"
					label="Confirm Password"
					bind:value={confirmPassword}
					error={confirmPasswordError}
					required
					autocomplete="new-password"
				>
					{#snippet icon()}
						<Lock class="w-5 h-5" />
					{/snippet}
				</Input>

				<div class="text-sm text-base-content/60">
					<p>Password must be at least 8 characters long</p>
				</div>

				<Button
					type="submit"
					variant="gradient"
					size="lg"
					loading={authService.loading}
					class="w-full"
				>
					{#snippet icon()}
						{#if !authService.loading}
							<UserPlus class="w-5 h-5" />
						{/if}
					{/snippet}
					{authService.loading ? 'Creating account...' : 'Sign Up'}
				</Button>
			</form>

			<div class="divider my-8">OR</div>

			<div class="text-center">
				<p class="text-base-content/60">
					Already have an account?
					<a href="/login" class="text-primary font-semibold hover:underline ml-1">
						Sign in
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
