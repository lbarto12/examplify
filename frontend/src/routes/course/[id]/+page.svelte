<script>
    /** @type {import('./$types').PageProps} */
    let { data } = $props(); // This pulls in the courseId from +page.js
</script>

<div class="flex flex-col min-h-screen bg-[#A6A6A6] font-sans">
    <header class="p-8 flex justify-center">
        <div class="relative w-full max-w-2xl">
            <input type="text" placeholder="🔍" 
                   class="input w-full rounded-full border-none bg-[#D9D9D9] px-6 py-2 shadow-inner focus:outline-none" />
        </div>
    </header>

    <main class="flex-grow flex flex-col items-center justify-center -mt-16">
        <h1 class="text-6xl font-black italic text-[#4D4D4D] tracking-widest mb-12 uppercase">
            {data.courseId}
        </h1>

        <div class="grid grid-cols-2 gap-6 w-full max-w-4xl px-10">
            <button class="btn btn-lg h-20 rounded-xl bg-[#D9D9D9] border-none text-[#4D4D4D] font-bold shadow-md hover:bg-black hover:text-white transition-all uppercase">Chapter Lectures</button>
            <button class="btn btn-lg h-20 rounded-xl bg-[#D9D9D9] border-none text-[#4D4D4D] font-bold shadow-md hover:bg-black hover:text-white transition-all uppercase">Summaries</button>
            <button class="btn btn-lg h-20 rounded-xl bg-[#D9D9D9] border-none text-[#4D4D4D] font-bold shadow-md hover:bg-black hover:text-white transition-all uppercase">Saved Questions</button>
            <button class="btn btn-lg h-20 rounded-xl bg-[#D9D9D9] border-none text-[#4D4D4D] font-bold shadow-md hover:bg-black hover:text-white transition-all uppercase">Flashcards</button>
            
            <div class="col-span-2 flex justify-center mt-4">
                <button class="btn btn-lg w-1/2 h-16 rounded-xl bg-[#D9D9D9] border-none text-[#4D4D4D] font-bold shadow-md hover:bg-black hover:text-white transition-all uppercase">Generate Quiz</button>
            </div>
        </div>
    </main>

    <footer class="bg-[#333333] text-white p-4 flex items-center justify-between px-10 mt-auto">
        <div class="text-2xl font-black italic tracking-tighter">AppName</div>
        <div class="flex gap-8 text-2xl items-center">
            <button class="hover:opacity-70">📤</button>
            <button class="hover:opacity-70">👤</button>
            <button class="hover:opacity-70">⚙️</button>
        </div>
    </footer>
</div>