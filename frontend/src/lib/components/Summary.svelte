<script lang="ts">
    import '../../app.css';

    type SummaryResult = {
        summary: string;
    };
    
    let { data }: {data: SummaryResult}  = $props();
</script>

<div class="card bg-base-100 p-4 shadow-md">
	<h2 class="text-xl font-bold mb-2">Summary</h2>
	<p>{data.summary}</p>
</div>

<style>
	@import '../../app.css';
</style>