<script lang="ts">
    import "../app.css";
    import Navbar from "$lib/components/ui/Navbar.svelte";

    let { children } = $props(); // Existing render logic
    let isExpanded = $state(false); // Reactive state for the "hover" effect
	let chatInput = $state("");
</script>

<div data-theme="lofi" class="flex h-screen bg-[#A6A6A6] font-sans overflow-hidden">
    <aside class="w-24 bg-[#4D4D4D]/20 flex flex-col items-center py-6 justify-between">
        <div class="space-y-4">
            <div class="w-12 h-12 bg-[#D9D9D9]/50 rounded-lg"></div>
            <div class="w-12 h-12 bg-[#D9D9D9]/50 rounded-lg"></div>
            <div class="w-12 h-12 bg-[#D9D9D9]/50 rounded-lg"></div>
            <div class="w-12 h-12 bg-[#D9D9D9]/50 rounded-lg"></div>
        </div>

        <div class="flex flex-col gap-6 mb-4">
            <button class="text-2xl opacity-70 hover:opacity-100 transition-opacity">📤</button>
            <button class="text-2xl opacity-70 hover:opacity-100 transition-opacity">👤</button>
            <button class="text-2xl opacity-70 hover:opacity-100 transition-opacity">⚙️</button>
        </div>
    </aside>

    <div class="flex flex-col flex-grow relative">
        <Navbar />
        
        <main class="flex-grow overflow-y-auto">
            {@render children()}
        </main>

        <div 
            role="complementary"
            class="absolute bottom-0 left-0 bg-white/95 shadow-2xl rounded-tr-3xl transition-all duration-300 ease-in-out border-t border-r border-black/10 flex flex-col z-50 overflow-hidden"
            style="width: {isExpanded ? '400px' : '320px'}; height: {isExpanded ? '480px' : '160px'};"
            onmouseenter={() => isExpanded = true}
            onmouseleave={() => isExpanded = false}
        >
            <div class="p-4 flex flex-col h-full">
                <div class="flex justify-between items-start mb-2">
                    <p class="text-[10px] font-black opacity-30 uppercase tracking-widest">Ask Mentor</p>
                    {#if isExpanded}
                        <button class="text-xs opacity-40 hover:opacity-100">⛶</button> {/if}
                </div>

                <div class="flex-grow overflow-y-auto">
                    <p class="text-sm italic text-gray-500">
                        {isExpanded ? "I'm here to help with your performance and answer questions about the lecture." : "Ask me anything..."}
                    </p>
                </div>

                <div class="mt-auto pt-4 border-t border-gray-100">
                    <input 
                        bind:value={chatInput}
                        type="text" 
                        placeholder="Type a message..." 
                        class="w-full bg-[#f3f4f6] rounded-lg px-3 py-2 text-sm focus:outline-none font-bold"
                    />
                </div>
            </div>
        </div>

        <footer class="bg-[#333333] text-white p-3 flex items-center justify-end px-10">
            <div class="text-xl font-black italic tracking-tighter uppercase">AppName</div>
        </footer>
    </div>
</div>