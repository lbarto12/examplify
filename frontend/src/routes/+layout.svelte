<script lang="ts">
    import "../app.css";
    import Toast from "$lib/components/ui/Toast.svelte";

    let { children } = $props(); // Existing render logic
</script>

<main class="flex-grow overflow-y-auto">
    {@render children()}
</main>

<!-- Global Toast Notifications -->
<Toast />
