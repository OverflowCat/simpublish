<script>
    import Login from "$lib/Login/index.svelte";
    import { page } from "$app/stores";
    let isAuthenticated = false;
    let url = decodeURIComponent(atob($page.params["redirect"])) || "";
</script>

<div>
    <Login bind:isAuthenticated redirectURL={url} />
</div>
