<script context="module">
	export const prerender = true;
</script>

<script>
	import List from '$lib/List.svelte';
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<section>
	<h1>
		<div class="welcome">

		</div>
		Simpublish
	</h1>
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

	h1 {
		width: 100%;
	}

</style>
