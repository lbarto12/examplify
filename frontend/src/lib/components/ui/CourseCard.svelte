<script lang="ts">
    let { name, onclick } = $props();
  </script>
  
  <button 
    {onclick}
    class="aspect-square bg-neutral-200 border-2 border-black hover:bg-white transition-all 
           flex items-center justify-center font-black text-xl hover:-translate-y-1 hover:shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]"
  >
    {name}
  </button>