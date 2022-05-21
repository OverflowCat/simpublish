<script>
    import Login from "$lib/Login/index.svelte";
    import { page } from "$app/stores";
    let isAuthenticated = false;
    // let url = decodeURIComponent($page.params["redirect"]) || "";
</script>

<div>
    <!--
        Debug: raw is {$page.params["redirect"]}, decoded is {url}.
    -->
    <Login bind:isAuthenticated redirectURL={$page.params["redirect"]} />
</div>

<style>
</style>
