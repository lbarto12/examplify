<script lang="ts">
	import { fly, fade, scale } from 'svelte/transition';
	import { BookOpen, Sparkles, Brain, GraduationCap, FileText, Zap, ArrowRight } from 'lucide-svelte';
	import Button from '$lib/components/ui/Button.svelte';
</script>

<div class="min-h-screen bg-base-100 overflow-hidden">
	<!-- Animated background -->
	<div class="fixed inset-0 -z-10 overflow-hidden">
		<div class="absolute top-[-20%] left-[-10%] w-150 h-150 rounded-full gradient-primary opacity-20 blur-3xl animate-blob"></div>
		<div class="absolute top-[40%] right-[-10%] w-125 h-125 rounded-full gradient-secondary opacity-15 blur-3xl animate-blob animation-delay-2000"></div>
		<div class="absolute bottom-[-10%] left-[30%] w-100 h-100 rounded-full gradient-accent opacity-10 blur-3xl animate-blob animation-delay-4000"></div>
	</div>

	<!-- Hero Section -->
	<section class="relative min-h-screen flex flex-col items-center justify-center px-6 py-20">
		<div class="max-w-4xl mx-auto text-center">
			<!-- Logo -->
			<div
				class="flex justify-center mb-8"
				in:scale={{ duration: 500, delay: 100, start: 0.8 }}
			>
				<div class="w-24 h-24 gradient-primary rounded-3xl flex items-center justify-center shadow-2xl">
					<BookOpen class="w-14 h-14 text-white" />
				</div>
			</div>

			<!-- Title -->
			<h1
				class="text-5xl md:text-7xl font-bold mb-6"
				in:fly={{ y: 30, duration: 600, delay: 200 }}
			>
				<span class="text-gradient-primary">Examplify</span>
			</h1>

			<!-- Tagline -->
			<p
				class="text-xl md:text-2xl text-base-content/70 mb-12 max-w-2xl mx-auto leading-relaxed"
				in:fly={{ y: 30, duration: 600, delay: 300 }}
			>
				Transform your study materials into interactive learning experiences powered by AI
			</p>

			<!-- CTA Buttons -->
			<div
				class="flex flex-col sm:flex-row gap-4 justify-center mb-20"
				in:fly={{ y: 30, duration: 600, delay: 400 }}
			>
				<a href="/signup" class="no-underline">
					<Button variant="gradient" size="lg" class="w-full sm:w-auto px-8">
						Get Started Free
						{#snippet icon()}
							<ArrowRight class="w-5 h-5 ml-2" />
						{/snippet}
					</Button>
				</a>
				<a href="/login" class="no-underline">
					<Button variant="outline" size="lg" class="w-full sm:w-auto px-8">
						Sign In
					</Button>
				</a>
			</div>

			<!-- Feature Pills -->
			<div
				class="flex flex-wrap justify-center gap-3"
				in:fade={{ duration: 600, delay: 500 }}
			>
				<span class="badge badge-lg badge-outline gap-2 px-4 py-3">
					<Sparkles class="w-4 h-4 text-primary" />
					AI-Powered
				</span>
				<span class="badge badge-lg badge-outline gap-2 px-4 py-3">
					<FileText class="w-4 h-4 text-secondary" />
					Any Document
				</span>
				<span class="badge badge-lg badge-outline gap-2 px-4 py-3">
					<Zap class="w-4 h-4 text-accent" />
					Instant Results
				</span>
			</div>
		</div>
	</section>

	<!-- Features Section -->
	<section class="relative py-24 px-6">
		<div class="max-w-6xl mx-auto">
			<div
				class="text-center mb-16"
				in:fly={{ y: 30, duration: 600, delay: 100 }}
			>
				<h2 class="text-3xl md:text-4xl font-bold mb-4">Study Smarter, Not Harder</h2>
				<p class="text-base-content/60 text-lg max-w-2xl mx-auto">
					Upload your study materials and let AI create personalized learning tools for you
				</p>
			</div>

			<div class="grid md:grid-cols-3 gap-8">
				<!-- Feature 1 -->
				<div
					class="card bg-base-200/50 backdrop-blur-sm border border-base-300 hover:border-primary/30 transition-all duration-300 hover:-translate-y-2"
					in:fly={{ y: 40, duration: 500, delay: 200 }}
				>
					<div class="card-body items-center text-center">
						<div class="w-16 h-16 gradient-primary rounded-2xl flex items-center justify-center mb-4 shadow-lg">
							<FileText class="w-8 h-8 text-white" />
						</div>
						<h3 class="card-title text-xl mb-2">Smart Summaries</h3>
						<p class="text-base-content/60">
							Get concise, well-structured summaries of your documents that highlight key concepts and important details
						</p>
					</div>
				</div>

				<!-- Feature 2 -->
				<div
					class="card bg-base-200/50 backdrop-blur-sm border border-base-300 hover:border-secondary/30 transition-all duration-300 hover:-translate-y-2"
					in:fly={{ y: 40, duration: 500, delay: 300 }}
				>
					<div class="card-body items-center text-center">
						<div class="w-16 h-16 gradient-secondary rounded-2xl flex items-center justify-center mb-4 shadow-lg">
							<Brain class="w-8 h-8 text-white" />
						</div>
						<h3 class="card-title text-xl mb-2">Interactive Flashcards</h3>
						<p class="text-base-content/60">
							Automatically generated flashcards help you memorize key terms and concepts through active recall
						</p>
					</div>
				</div>

				<!-- Feature 3 -->
				<div
					class="card bg-base-200/50 backdrop-blur-sm border border-base-300 hover:border-accent/30 transition-all duration-300 hover:-translate-y-2"
					in:fly={{ y: 40, duration: 500, delay: 400 }}
				>
					<div class="card-body items-center text-center">
						<div class="w-16 h-16 gradient-accent rounded-2xl flex items-center justify-center mb-4 shadow-lg">
							<GraduationCap class="w-8 h-8 text-white" />
						</div>
						<h3 class="card-title text-xl mb-2">Practice Quizzes</h3>
						<p class="text-base-content/60">
							Test your knowledge with AI-generated quizzes that adapt to help you focus on areas that need improvement
						</p>
					</div>
				</div>
			</div>
		</div>
	</section>

	<!-- CTA Section -->
	<section class="relative py-24 px-6">
		<div
			class="max-w-4xl mx-auto text-center"
			in:fly={{ y: 30, duration: 600, delay: 100 }}
		>
			<div class="card gradient-primary text-white shadow-2xl">
				<div class="card-body py-16 px-8">
					<h2 class="text-3xl md:text-4xl font-bold mb-4">Ready to Transform Your Study Habits?</h2>
					<p class="text-white/80 text-lg mb-8 max-w-xl mx-auto">
						Join thousands of students who are already studying smarter with Examplify
					</p>
					<div class="flex flex-col sm:flex-row gap-4 justify-center">
						<a href="/signup" class="no-underline">
							<button class="btn btn-lg bg-white text-primary hover:bg-white/90 border-0 px-8">
								Create Free Account
								<ArrowRight class="w-5 h-5 ml-2" />
							</button>
						</a>
					</div>
				</div>
			</div>
		</div>
	</section>

	<!-- Footer -->
	<footer class="py-8 px-6 border-t border-base-300">
		<div class="max-w-6xl mx-auto flex flex-col md:flex-row items-center justify-between gap-4">
			<div class="flex items-center gap-2">
				<BookOpen class="w-5 h-5 text-primary" />
				<span class="font-semibold">Examplify</span>
			</div>
			<p class="text-sm text-base-content/50">
				Study smarter with AI-powered learning tools
			</p>
		</div>
	</footer>
</div>

<style>
	@keyframes blob {
		0%, 100% {
			transform: translate(0, 0) scale(1);
		}
		25% {
			transform: translate(20px, -30px) scale(1.1);
		}
		50% {
			transform: translate(-20px, 20px) scale(0.9);
		}
		75% {
			transform: translate(30px, 30px) scale(1.05);
		}
	}

	.animate-blob {
		animation: blob 15s ease-in-out infinite;
	}

	.animation-delay-2000 {
		animation-delay: 2s;
	}

	.animation-delay-4000 {
		animation-delay: 4s;
	}
</style>
