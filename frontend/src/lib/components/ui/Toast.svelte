<script lang="ts">
	import { toastStore } from '$lib/services';
	import { CheckCircle, XCircle, Info, AlertTriangle, X } from 'lucide-svelte';
	import { fly } from 'svelte/transition';

	const icons = {
		success: CheckCircle,
		error: XCircle,
		info: Info,
		warning: AlertTriangle
	};

	const alertClasses = {
		success: 'alert-success',
		error: 'alert-error',
		info: 'alert-info',
		warning: 'alert-warning'
	};
</script>

<div class="toast toast-top toast-end z-50">
	{#each toastStore.messages as toast (toast.id)}
		<div
			class={`alert ${alertClasses[toast.type]} shadow-lg min-w-[300px]`}
			transition:fly={{ x: 300, duration: 300 }}
		>
			<svelte:component this={icons[toast.type]} class="w-6 h-6" />
			<span>{toast.message}</span>
			<button
				class="btn btn-sm btn-ghost btn-circle"
				onclick={() => toastStore.remove(toast.id)}
			>
				<X class="w-4 h-4" />
			</button>
		</div>
	{/each}
</div>
