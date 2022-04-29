<script>
	import { page } from '$app/stores';
	import logo from './logo.svg';
</script>

<header>
	<div class="corner">
		<a href="/">
			<img src={logo} alt="Logo" />
		</a>
	</div>

	<nav>

	</nav>

	<div class="corner">
		<!-- TODO put something else here? github link? -->
	</div>
</header>

<style>
	header {
		display: flex;
		justify-content: space-between;
	}

	.corner {
		width: 3em;
		height: 3em;
	}

	.corner a {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100%;
	}

	.corner img {
		width: 2em;
		height: 2em;
		object-fit: contain;
	}

	nav {
		display: flex;
		justify-content: center;
		--background: rgba(255, 255, 255, 0.7);
	}

	a:hover {
		color: var(--accent-color);
	}
</style>
