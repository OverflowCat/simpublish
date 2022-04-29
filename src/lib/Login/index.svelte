<script>
    let password = "";
    let message = "";
    export let redirectURL = "";
    export let isAuthenticated = false;
    import { auth } from "./auth";
    import Cookies from "js-cookie";
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

<div>
    <label>
        密码
        <input type="password" bind:value={password} />
    </label>
    <button on:click={login}> 验证 </button>
</div>
