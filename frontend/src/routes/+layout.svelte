<script lang="ts">
    import "../app.css";

    let { children } = $props(); // Existing render logic
</script>

<main class="flex-grow overflow-y-auto">
    {@render children()}
</main>
