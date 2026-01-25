<script>
    /** @type {import('./$types').PageProps} */
    let { data } = $props(); 
</script>

<div class="grid grid-cols-12 gap-10 p-10 min-h-[90vh] bg-base-100 font-sans">
    <aside class="col-span-3 space-y-4">
        <p class="text-[10px] font-black opacity-30 tracking-[0.3em] uppercase">Temporary Action</p>
        <div class="flex flex-col gap-2">
            {#each data.quickAccess as action}
                <button class="btn btn-outline btn-md rounded-none justify-start border-2 border-black font-bold uppercase italic hover:bg-black hover:text-white transition-all">
                    <span>{action.icon}</span> {action.label}
                </button>
            {/each}
        </div>
    </aside>

    <main class="col-span-9 space-y-12">
        <div class="relative">
            <input type="text" placeholder="SEARCH ANYTHING..." 
                   class="input input-bordered input-lg w-full rounded-full border-2 border-black font-black italic shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] focus:outline-none focus:translate-x-1 focus:translate-y-1 focus:shadow-none transition-all" />
        </div>

        <div class="grid grid-cols-3 gap-8">
            {#each data.savedCourses as course}
                <div class="aspect-square bg-black text-white flex flex-col items-center justify-center p-6 hover:bg-neutral cursor-pointer shadow-[12px_12px_0px_0px_rgba(0,0,0,0.1)] transition-all group border-2 border-black">
                    <span class="text-2xl font-black italic tracking-tighter uppercase text-center leading-none">
                        {course.name}
                    </span>
                </div>
            {/each}
            <button class="aspect-square border-4 border-dashed border-black/10 flex items-center justify-center hover:bg-base-200 transition-all group">
                <span class="text-6xl font-thin opacity-20 group-hover:opacity-100">+</span>
            </button>
        </div>
    </main>
</div>