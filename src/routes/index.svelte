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
		height: 100%;
	}

	h1 {
		font-family: Cambria, Cochin, Georgia, Times, "Times New Roman", serif;
		font-size: large;
	}
</style>
