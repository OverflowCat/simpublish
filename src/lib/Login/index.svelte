<script>
    let password = "";
    export let isAuthenticated = false;
    import { onMount } from "svelte";
    import { auth } from "./auth";
    async function login() {
        const res = await auth(password);
        console.log("Res is", res);
        const success = res["result"];
        if (success) {
            isAuthenticated = true;
            onMount(() => {
                alert("验证成功！");
            });
        } else {
            isAuthenticated = false;
            onMount(() => {
                alert("密码错误或网络连接失败");
            });
        }
    }
</script>

<div>
    <label>
        密码
        <input type="password" bind:value={password} />
    </label>
    <button on:click={login}> 验证 </button>
</div>
