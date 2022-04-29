<script context="module">
	export const prerender = true;
</script>

<script>
	import List from "$lib/List.svelte";
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>
<h1>Simpublish</h1>
<section>
	<List />
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 1;
	}
</style>
