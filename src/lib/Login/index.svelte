<script>
    let password = "";
    let message = "";
    export let redirectURL = "";
    export let isAuthenticated = false;
    import { auth } from "./auth";
    import Cookies from "js-cookie";
    import { Button } from "carbon-components-svelte";
    import Login from "carbon-icons-svelte/lib/Login.svelte";
    import { PasswordInput } from "carbon-components-svelte";
    isAuthenticated = Cookies.get("password") !== undefined;

    async function login() {
        const res = await auth(password);
        const success = res["result"];
        if (success) {
            isAuthenticated = true;
            message = "验证成功！";
            Cookies.set("password", password, { expires: 31 });
            if (redirectURL !== "") {
                window.location.href = redirectURL;
            }
        } else {
            isAuthenticated = false;
            message = "验证失败！";
        }
    }
</script>

<div class="passwdinput">
    <PasswordInput
        size="xl"
        labelText="密码"
        placeholder="环境变量中设置的密码…"
        bind:value={password}
    />
    <Button icon={Login} on:click={login}>验证</Button>
</div>

<style>
    .passwdinput {
        width: 90%;
    }
</style>