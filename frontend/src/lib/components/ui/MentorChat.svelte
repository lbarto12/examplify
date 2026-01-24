<script lang="ts">
    let isExpanded = $state(false);
  
    // Added this to satisfy accessibility rules
    function handleKeyDown(e: KeyboardEvent) {
      if (e.key === 'Enter' || e.key === ' ') {
        isExpanded = !isExpanded;
      }
    }
  </script>
  
  <div 
    role="button"
    tabindex="0"
    aria-expanded={isExpanded}
    class="fixed bottom-20 left-6 transition-all duration-300 border-2 border-black bg-white z-50 shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]
    {isExpanded ? 'w-80 h-96' : 'w-48 h-12'}"
    onmouseenter={() => isExpanded = true}
    onmouseleave={() => isExpanded = false}
    onkeydown={handleKeyDown}
  >
    <div class="p-3 flex justify-between items-center bg-black text-white">
      <span class="text-[10px] font-black uppercase tracking-widest">Ask Mentor</span>
      {#if isExpanded}
         <span class="text-[10px] opacity-50">v1.0</span>
      {/if}
    </div>
    
    {#if isExpanded}
      <div class="p-4 h-64 overflow-y-auto text-xs font-mono space-y-2 animate-in fade-in">
        <p class="bg-neutral-100 p-2 border border-black/10">Hello! I'm your AI assistant. How is your studying going?</p>
      </div>
    {/if}
  
    <div class="absolute bottom-0 w-full p-2 border-t-2 border-black bg-white">
      <input 
        type="text" 
        placeholder="Type a question..." 
        class="w-full text-xs focus:outline-none bg-transparent uppercase font-bold"
      />
    </div>
  </div>