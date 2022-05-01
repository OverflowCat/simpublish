<script>
    export let count = 0;
    import VirtualList from "@sveltejs/svelte-virtual-list";
    import Item from "./Item.svelte";
    import Login from "./Login/index.svelte";
    let items = [];
    import Cookies from "js-cookie";
    let isAuthenticated = true;
    isAuthenticated = Cookies.get("password") !== undefined;
    function refreshItems() {
        if (isAuthenticated) {
            fetch("/api/list")
                .then((res) => res.json())
                .then((res) => {
                    items = res.sort((a, b) => b.id - a.id);
                });
        }
    }
    $: count = items.length;
    $: isAuthenticated, refreshItems();
</script>

<div class="container">
    {#if isAuthenticated}
        <VirtualList {items} let:item height="500px">
            <Item id={item.id} title={item.title} size={item.size} />
        </VirtualList>
    {:else}
        <Login bind:isAuthenticated />
    {/if}
</div>

<style>
    .container {
        width: 90%;
        margin-top: 8px;
    }
</style>