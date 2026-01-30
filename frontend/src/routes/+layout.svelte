<script lang="ts">
    import "../app.css";
    import Toast from "$lib/components/ui/Toast.svelte";

    let { children } = $props();
</script>

<div class="min-h-screen bg-base-100">
    {@render children()}
</div>

<!-- Global Toast Notifications -->
<Toast />
