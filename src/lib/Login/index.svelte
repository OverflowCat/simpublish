<script>
    import "carbon-components-svelte/css/all.css";
    let password = "";
    let message = "";
    export let redirectURL = "";
    export let isAuthenticated = false;
    import { auth } from "./auth";
    import Cookies from "js-cookie";
    import { Button, ToastNotification, Theme } from "carbon-components-svelte";
    import Login from "carbon-icons-svelte/lib/Login.svelte";
    import Clean from "carbon-icons-svelte/lib/Clean.svelte";
    import { PasswordInput } from "carbon-components-svelte";
    import { onMount } from "svelte";
    isAuthenticated = Cookies.get("password") !== undefined;
    let isCookieCleared = false;
    let isWrongPassword = false;
    async function login() {
        isWrongPassword = false;
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
    
    let theme = "white"; // "white" | "g10" | "g80" | "g90" | "g100"
    onMount(() => {
        const mediaQueryListDark = window.matchMedia(
            "(prefers-color-scheme: dark)"
        );
        if (mediaQueryListDark.matches) {
            theme = "g90";
        }

        function handleChange(mediaQueryListEvent) {
            if (mediaQueryListEvent.matches) {
                theme = "g90";
            } else {
                theme = "white";
            }
        }

        // 添加主题变动监控事件
        mediaQueryListDark.addListener(handleChange);
    });
</script>

<div class="passwdinput">
    <Theme bind:theme />
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
