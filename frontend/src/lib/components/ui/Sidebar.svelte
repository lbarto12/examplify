<script lang="ts">
    let actions = ['Lecture', 'Summary', 'Flashcards', 'Quizzes'];
  </script>
  
  <aside class="w-48 border-r-2 border-black bg-white flex flex-col p-4 gap-4 h-full">
    <div class="text-[10px] font-black uppercase tracking-widest opacity-40 mb-2">Actions</div>
    
    {#each actions as action}
      <button class="btn btn-outline btn-sm rounded-none border-2 border-black font-black uppercase text-[10px] tracking-tighter hover:bg-black hover:text-white transition-all">
        {action}
      </button>
    {/each}
  </aside>