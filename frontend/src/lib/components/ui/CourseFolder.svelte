<script lang="ts">
    let { courseName, onBack } = $props();
  </script>
  
  <div class="max-w-4xl mx-auto text-center animate-in fade-in duration-300">
    <button onclick={onBack} class="text-[10px] uppercase underline mb-8 block hover:font-bold">
      ← Back to Dashboard
    </button>
    
    <h1 class="text-6xl font-black uppercase mb-16 tracking-tighter">{courseName}</h1>
    
    <div class="grid grid-cols-2 gap-6">
      {#each ['Chapter Lectures', 'Summaries', 'Saved Questions', 'Flashcards'] as label}
        <button class="btn btn-lg h-24 rounded-none border-2 border-black bg-neutral-100 uppercase font-black hover:bg-white">
          {label}
        </button>
      {/each}
      <button class="btn btn-lg h-24 rounded-none border-2 border-black bg-black text-white col-span-2 uppercase font-black hover:bg-neutral-800">
        Generate Quiz
      </button>
    </div>
  </div>