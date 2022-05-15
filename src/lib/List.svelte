<script>
    export let count = 0;
    let skeleton = true;
    import VirtualList from "@sveltejs/svelte-virtual-list";
    import Item from "./Item.svelte";
    import Login from "./Login/index.svelte";
    import Cookies from "js-cookie";
    let isDebug = true;

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
            title: "Example 6 Long long long longLong long long longLong long long long long long long long long title",
            id: 666,
            size: 114514,
        },
        {
            title: "Example 1",
            id: 88,
            size: 2500,
        },
        {
            title: "Example 2jj",
            id: 4894,
            size: 1246,
        },
        {
            title: "Example 6 Long long long longLong long long longLong long long long long long long long long title",
            id: 666,
            size: 114514,
        },
        {
            title: "Example 1",
            id: 88,
            size: 2500,
        },
        {
            title: "Example 2jj",
            id: 4894,
            size: 1246,
        },
    ];
    let items = DEFAULT_ITEMS;
    let isAuthenticated = isDebug || (Cookies.get("password") !== undefined);
    function refreshItems() {
        if (isAuthenticated && !isDebug) {
            fetch("/api/list")
                .then((res) => res.json())
                .then((res) => {
                    items = res.sort((a, b) => b.id - a.id);
                    skeleton = false;
                });
        } else if (isDebug) {
            skeleton = false;
        } else {
            items = DEFAULT_ITEMS;
            skeleton = true;
        }
    }
    $: count = items.length;
    $: isAuthenticated, refreshItems();
</script>

<div class="container scrollview">
    {#if isAuthenticated}
        <VirtualList {items} let:item>
            <Item id={item.id} title={item.title} size={item.size} {skeleton} />
        </VirtualList>
    {:else}
        <Login bind:isAuthenticated />
    {/if}
</div>

<style>
    .container {
        width: 100%;
        margin-top: 8px;
        margin-left: 4px;
        margin-right: 4px;
        height: calc(100vh - 10rem);
    }
</style>
