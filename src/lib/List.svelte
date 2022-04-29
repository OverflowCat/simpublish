<script>
    export let count = 0;
    import VirtualList from "@sveltejs/svelte-virtual-list";
    const things = [
        // these can be any values you like
        { title: "one", id: 1 },
        { title: "two", id: 2 },
        { title: "three", id: 3 },
        // ...
        { title: "six thousand and ninety-two", id: 6092 },
    ];
    $: count = things.length;
</script>

<div>
    From List component
    <VirtualList items={things} let:item height="500px">
        <!-- this will be rendered for each currently visible item -->
        <p>{item.id}: {item.name}</p>
    </VirtualList>
</div>
