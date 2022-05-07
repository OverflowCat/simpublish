<script>
    let password = "";
    let message = "";
    export let redirectURL = "";
    export let isAuthenticated = false;
    import { auth } from "./auth";
    import Cookies from "js-cookie";
    import { Button, ToastNotification } from "carbon-components-svelte";
    import Login from "carbon-icons-svelte/lib/Login.svelte";
    import Clean from "carbon-icons-svelte/lib/Clean.svelte";
    import { PasswordInput } from "carbon-components-svelte";
    isAuthenticated = Cookies.get("password") !== undefined;
    let isCookieCleared = false;
    let isWrongPassword = false;
    async function login() {
        const res = await auth(password);
        const success = res["result"];
        if (success) {
            isAuthenticated = true;
            message = "验证成功！";
            Cookies.set("password", password, { expires: 31 });
            isCookieCleared = false;
            isWrongPassword = false;
            if (redirectURL !== "") {
                window.location.href = redirectURL;
            }
        } else {
            isAuthenticated = false;
            message = "验证失败！";
            isWrongPassword = true;
        }
    }

    function clearCookie() {
        Cookies.remove("password");
        isAuthenticated = false;
        isCookieCleared = true;
    }
</script>

<div class="passwdinput">
    <PasswordInput
        size="xl"
        labelText="密码"
        placeholder="环境变量中设置的密码…"
        bind:value={password}
    />
    <Button icon={Login} on:click={login}>验证密码</Button>
    <Button
        disabled={isCookieCleared}
        kind="danger-tertiary"
        icon={Clean}
        onLclick={clearCookie}
        >清除 Cookie
    </Button>
    <p class="info">如果你修改了密码，请尝试清除 Cookie 后再次登录。</p>
    {#if isWrongPassword}
        <ToastNotification
            title="验证失败"
            subtitle="密码有误"
            caption={new Date().toLocaleString()}
        />
    {/if}
</div>

<style>
    .passwdinput {
        width: 90%;
    }
    .info {
        margin: 2px;
        margin-top: 10px;
        color: gray;
    }
</style>
