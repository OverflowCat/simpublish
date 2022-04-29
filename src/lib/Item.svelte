<script>
    export let title = "",
        id = 0,
        size = 0;
    function formatSizeUnits(bytes) {
        if ((bytes >> 30) & 0x3ff)
            bytes = (bytes >>> 30) + "." + (bytes & (3 * 0x3ff)) + " GB";
        else if ((bytes >> 20) & 0x3ff)
            bytes = (bytes >>> 20) + "." + (bytes & (2 * 0x3ff)) + " MB";
        else if ((bytes >> 10) & 0x3ff)
            bytes = (bytes >>> 10) + "." + (bytes & 0x3ff) + " KB";
        else if ((bytes >> 1) & 0x3ff) bytes = (bytes >>> 1) + " Bytes";
        else bytes = bytes + "Byte";
        return bytes;
    }
</script>

<div class="container">
    <div id="id">{id}</div>
    <span id="title"><a href="/articles/{id}/">{title}</a></span>
    <span id="size">{@html formatSizeUnits(size)}</span>
</div>

<style>
    .container {
        padding: 4px;
    }
    #id {
        font-size: smaller;
        font-family: "JetBrains Mono", "JetBrainsMono NF", "Fira Code",
            "FiraCode NF", "Noto Sans Mono CJK SC", "Courier New", Courier,
            Consolas, monospace;
        text-align: center;
        min-width: 40px;
        float: left;
        background-color: rgb(200, 200, 200);
        margin-top: 1px;
        padding: 2px;
        vertical-align: middle;
    }
    #title {
        margin-left: 8px;
        margin-right: 8px;
    }
    #size {
        color: rgb(62, 62, 62);
        font-size: smaller;
        font-family: "JetBrains Mono", "JetBrainsMono NF", "Fira Code",
            "FiraCode NF", "Noto Sans Mono CJK SC", "Courier New", Courier,
            Consolas, monospace;
        float: right;
        margin-top: 1px;
        padding: 2px;
    }
    a {
        color: #2f56b0;
        position: relative;
        transition: color 0.4s ease-out;
        text-decoration-color: rgb(222, 222, 222);
        text-decoration-skip: ink;
    }
    
    a:hover {
        color: #457dfb;
        right: 0;
        text-decoration: rgb(244, 244, 244);
    }

    a:hover:after {
        border-color: #457dfb;
        right: 0;
    }

    a:after {
        border-radius: 1em;
        border-top: 0.1em solid #2f56b0;
        content: "";
        position: absolute;
        right: 100%;
        bottom: 0.14em;
        left: 0;
        transition: right 0.4s cubic-bezier(0, 0.5, 0, 1),
            border-color 0.4s ease-out;
    }

    a:hover:after {
        animation: anchor-underline 2s cubic-bezier(0, 0.5, 0, 1) infinite;
        border-color: #457dfb;
    }

    @keyframes anchor-underline {
        0%,
        10% {
            left: 0;
            right: 100%;
        }
        40%,
        60% {
            left: 0;
            right: 0;
        }
        90%,
        100% {
            left: 100%;
            right: 0;
        }
    }

</style>
