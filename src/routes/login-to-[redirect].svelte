<script>
    import Login from "$lib/Login/index.svelte";
    import { page } from "$app/stores";
    let isAuthenticated = false;
    let url = decodeURIComponent($page.params["redirect"]) || "";
</script>

<div>
    <div>Debug: raw is {$page.params["redirect"]}, decoded is {url}.</div>
    <Login bind:isAuthenticated redirectURL={url} />
</div>
