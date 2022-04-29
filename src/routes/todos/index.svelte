<script>
	import { enhance } from '$lib/form';
	import { scale } from 'svelte/transition';
	import { flip } from 'svelte/animate';

	export let todos;
</script>

<svelte:head>
	<title>Todos</title>
</svelte:head>

<div class="todos">
	<h1>Todos</h1>

	<form
		class="new"
		action="/todos"
		method="post"
		use:enhance={{
			result: async ({ form }) => {
				form.reset();
			}
		}}
	>
		<input name="text" aria-label="Add todo" placeholder="+ tap to add a todo" />
	</form>

	{#each todos as todo (todo.uid)}
		<div
			class="todo"
			class:done={todo.done}
			transition:scale|local={{ start: 0.7 }}
			animate:flip={{ duration: 200 }}
		>
			<form
				action="/todos?_method=PATCH"
				method="post"
				use:enhance={{
					pending: ({ data }) => {
						todo.done = !!data.get('done');
					}
				}}
			>
				<input type="hidden" name="uid" value={todo.uid} />
				<input type="hidden" name="done" value={todo.done ? '' : 'true'} />
				<button class="toggle" aria-label="Mark todo as {todo.done ? 'not done' : 'done'}" />
			</form>

			<form class="text" action="/todos?_method=PATCH" method="post" use:enhance>
				<input type="hidden" name="uid" value={todo.uid} />
				<input aria-label="Edit todo" type="text" name="text" value={todo.text} />
				<button class="save" aria-label="Save todo" />
			</form>

			<form
				action="/todos?_method=DELETE"
				method="post"
				use:enhance={{
					pending: () => (todo.pending_delete = true)
				}}
			>
				<input type="hidden" name="uid" value={todo.uid} />
				<button class="delete" aria-label="Delete todo" disabled={todo.pending_delete} />
			</form>
		</div>
	{/each}
</div>

<style>
</style>
