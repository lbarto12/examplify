<script lang="ts">
    import { userState } from '$lib/state/user.svelte.ts';
    import { goto } from '$app/navigation';
	import { getClient } from '$lib/apis/sessions.svelte';
  
    let email = $state("");
    let password = $state("");

    async function handleLogin() {
      const sessionAPI = getClient();

      const { token } = await sessionAPI.signIn({ 
        email: email,
        password: password,
      });

      await fetch("/actions/set-cookie", {
        method: "POST",
        body: JSON.stringify({
          key: "auth",
          value: token,
        }),
      });

      console.log("Session token: ", token);
    }
  </script>
  
  <div class="flex flex-col items-center justify-center min-h-[70vh] px-4">
    <div class="w-full max-w-sm">
      
      <header class="mb-8 text-center">
        <h1 class="text-3xl font-black uppercase tracking-tighter">Access System</h1>
        <p class="text-xs font-mono opacity-50 uppercase mt-2">Education Interface v1.0</p>
      </header>
  
      <div class="space-y-6">
        <div class="form-control w-full">
          <label class="label py-1">
            <span class="label-text text-[10px] font-bold uppercase tracking-widest">Identification</span>
          </label>
          <input 
            type="text" 
            bind:value={email}
            placeholder="ENTER USERNAME" 
            class="input input-bordered w-full rounded-none border-black border-2 focus:outline-none placeholder:opacity-30 text-sm"
          />
        </div>
  
        <div class="form-control w-full">
          <label class="label py-1">
            <span class="label-text text-[10px] font-bold uppercase tracking-widest">Security Token</span>
          </label>
          <input 
            type="password" 
            bind:value={password}
            placeholder="••••••••" 
            class="input input-bordered w-full rounded-none border-black border-2 focus:outline-none placeholder:opacity-30 text-sm"
          />
        </div>
  
        <button 
          class="btn btn-primary w-full rounded-none border-2 border-black font-black uppercase tracking-widest mt-4 hover:bg-black hover:text-white transition-colors"
          onclick={handleLogin}
        >
          Authenticate
        </button>
      </div>
  
      <footer class="mt-12 border-t border-black/10 pt-4">
        <p class="text-[9px] font-mono text-center opacity-40 uppercase">
          Secure Terminal / Authorized Personnel Only
        </p>
      </footer>
    </div>
  </div>