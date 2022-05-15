<script context="module">
	export const prerender = true;
	import config from "../custom.js";
</script>

<script>
	import List from "$lib/List.svelte";
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>
<h1>{config["site-name"]}<span id="slogan">{config["site-slogan"]}</span></h1>
<section>
	<List />
</section>

<style>
	section {
		height: 100%;
	}

	#slogan {
		margin-left: 20px;
		font-size: 0.8em;
		font-family: "Ysabeau", Tahoma, Helvetica, "PingFang SC", "Hiragino Sans GB",
			"Franklin Gothic Medium", "Arial Narrow", "Noto Sans CJK SC",
			sans-serif;
		font-weight: 600;
		color: #aaa;
	}

	h1 {
		font-family: Cambria, Cochin, Georgia, "Noto Sans CJK SC", Times, "Times New Roman", serif;
		font-size: large;
		font-weight: 800;
	}
</style>
