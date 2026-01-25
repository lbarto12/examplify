<script>
    /** @type {import('./$types').PageProps} */
    let { data } = $props(); //  receives the 'load' data
</script>

<div class="p-8 space-y-8">
    <input type="text" placeholder="Search anything..." class="input input-bordered w-full rounded-full" />

    <div class="grid grid-cols-12 gap-6">
        <aside class="col-span-3 flex flex-col gap-2">
            {#each data.quickAccess as action}
                <button class="btn btn-outline rounded-none justify-start">{action.icon} {action.label}</button>
            {/each}
        </aside>

        <main class="col-span-9 flex flex-wrap gap-4">
            {#each data.savedCourses as course}
                <div class="card w-40 h-40 bg-neutral text-neutral-content rounded-none flex items-center justify-center">
                    <span class="font-bold">{course.name}</span>
                </div>
            {/each}
            <button class="card w-40 h-40 border-2 border-dashed border-base-content/20 rounded-none flex items-center justify-center">+</button>
        </main>
    </div>
</div>