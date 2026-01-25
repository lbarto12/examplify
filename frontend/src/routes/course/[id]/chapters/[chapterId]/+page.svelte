<script>
    /** @type {import('./$types').PageProps} */
    let { data } = $props(); // Strict Svelte 5 props
    
    let chapters = [
        { id: 'ch1', name: 'CH1 - NAME OF CHAPTER' },
        { id: 'ch2', name: 'CH2 - NAME OF CHAPTER' },
        { id: 'ch3', name: 'CH3 - NAME OF CHAPTER' }
    ];
</script>

<div class="flex flex-col min-h-screen bg-[#A6A6A6] font-sans">
    <header class="p-8 flex items-center justify-between max-w-4xl mx-auto w-full">
        <a href="/course/{data.id.toLowerCase()}" class="btn btn-circle bg-[#D9D9D9] border-none text-2xl no-underline flex items-center justify-center">←</a>
        <div class="relative flex-grow mx-10">
            <input type="text" placeholder="🔍" class="input w-full rounded-full border-none bg-[#D9D9D9] shadow-inner" />
        </div>
    </header>

    <main class="flex-grow flex flex-col items-center px-10">
        <h1 class="text-5xl font-black italic text-[#4D4D4D] tracking-widest mb-10 uppercase">
            {data.id} - CHAPTERS
        </h1>

        <div class="w-full max-w-2xl space-y-4">
            {#each chapters as ch}
                <a href="/course/{data.id.toLowerCase()}/chapters/{ch.id}" 
                   class="flex items-center bg-[#D9D9D9] p-5 rounded-xl border-none shadow-md font-bold text-[#4D4D4D] hover:bg-black hover:text-white transition-all no-underline">
                    {ch.name}
                </a>
            {/each}
        </div>

        <button class="btn btn-circle btn-lg w-20 h-20 bg-[#D9D9D9] border-none mt-12 shadow-2xl text-5xl text-[#4D4D4D]">
            +
        </button>
    </main>

    <footer class="bg-[#333333] text-white p-4 flex items-center justify-between px-10 mt-10">
        <div class="text-2xl font-black italic">AppName</div>
        <div class="flex gap-8 text-2xl">
            <span>📤</span> <span>👤</span> <span>⚙️</span>
        </div>
    </footer>
</div>