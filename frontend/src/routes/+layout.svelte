<script lang="ts">
	import "../app.css";
	import Navbar from "$lib/components/ui/Navbar.svelte";
	let { children } = $props();
  </script>
  
  <div data-theme="lofi" class="min-h-screen bg-base-100 font-sans">
	<Navbar />
	
	<main class="max-w-4xl mx-auto p-8">
	  {@render children()}
	</main>
  </div>