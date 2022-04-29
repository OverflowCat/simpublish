<script>
    export let count = 0;
    import VirtualList from "@sveltejs/svelte-virtual-list";
    import { onMount } from "svelte";
    import Item from "./Item.svelte";
    import Login from "./Login/index.svelte";
    let items = [];
    let isAuthenticated = false;
    function refreshItems() {
        if (isAuthenticated) {
            onMount(() => {
                fetch("/api/list")
                    .then((res) => res.json())
                    .then((res) => {
                        items = res;
                    });
            });
        }
    }
    $: count = items.length;
    $: isAuthenticated, refreshItems();
</script>

<div>
    {#if isAuthenticated}
        <VirtualList {items} let:item height="500px">
            <Item id={item.id} title={item.title} size={item.size} />
        </VirtualList>
    {:else}
        <Login bind:isAuthenticated />
    {/if}
</div>
