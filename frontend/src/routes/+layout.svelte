<script lang="ts">
	import "../app.css";
	let { children } = $props();
  </script>
  
  <div data-theme="lofi" class="min-h-screen bg-base-100 font-sans">
	<main class="max-w-4xl mx-auto p-8">
	  {@render children()}
	</main>
  </div>