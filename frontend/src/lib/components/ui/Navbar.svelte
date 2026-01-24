<script lang="ts">
    import { userState } from '$lib/state/user.svelte.ts';
</script>

<div class="navbar bg-base-100 border-b-2 border-black">
    <div class="flex-1">
        <a href="/" class="btn btn-ghost text-xl rounded-none font-black uppercase tracking-widest">
            Project
        </a>
    </div>

    <div class="flex-none px-2">
        {#if userState.isLoggedIn}
            <div class="flex items-center gap-4">
                <span class="text-xs font-mono uppercase">{userState.username}</span>
                <button 
                    class="btn btn-sm btn-outline rounded-none" 
                    onclick={() => userState.clearSession()}
                >
                    LOGOUT
                </button>
            </div>
        {:else}
            <a 
                href="/login"
                class="btn btn-sm btn-primary rounded-none"
            >
                LOGIN
            </a>
        {/if}
    </div>
</div>