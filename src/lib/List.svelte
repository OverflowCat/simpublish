<script>
    export let count = 0;
    let skeleton = true;
    import VirtualList from "@sveltejs/svelte-virtual-list";
    import Item from "./Item.svelte";
    import Login from "./Login/index.svelte";
    import Cookies from "js-cookie";

    const DEFAULT_ITEMS = [
        {
            title: "Example 1",
            id: 104,
            size: 2500,
        },
        {
            title: "Example 2",
            id: 233,
            size: 1246,
        },
        {
            title: "Example 3",
            id: 567,
            size: 1246,
        },
        {
            title: "Example 6",
            id: 666,
            size: 114514,
        },
    ];
    let items = DEFAULT_ITEMS;
    let isAuthenticated = true;
    isAuthenticated = Cookies.get("password") !== undefined;
    function refreshItems() {
        if (isAuthenticated) {
            fetch("/api/list")
                .then((res) => res.json())
                .then((res) => {
                    items = res.sort((a, b) => b.id - a.id);
                    skeleton = false;
                });
        } else {
            items = DEFAULT_ITEMS;
            skeleton = true;
        }
    }
    $: count = items.length;
    $: isAuthenticated, refreshItems();
</script>

<div class="container">
    {#if isAuthenticated}
        <VirtualList {items} let:item height="500px">
            <Item id={item.id} title={item.title} size={item.size} {skeleton} />
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
