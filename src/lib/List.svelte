<script>
    export let count = 0;
    import VirtualList from "@sveltejs/svelte-virtual-list";
    const things = [
        // these can be any values you like
        { name: "one", number: 1 },
        { name: "two", number: 2 },
        { name: "three", number: 3 },
        // ...
        { name: "six thousand and ninety-two", number: 6092 },
    ];
    $: count = things.length;
</script>

<div>
    From List component
    <VirtualList items={things} let:item height="500px">
        <!-- this will be rendered for each currently visible item -->
        <p>{item.number}: {item.name}</p>
    </VirtualList>
</div>
