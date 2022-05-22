<script>
    import {
        Tile,
        Button,
        SkeletonPlaceholder,
        truncate,
    } from "carbon-components-svelte";
    import CopyLink from "carbon-icons-svelte/lib/CopyLink.svelte";

    export let title = "",
        id = 0,
        size = 0,
        skeleton = true;
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
    function copyLink(text) {
        const url = window.location.origin + text;
        navigator.clipboard.writeText(url);
    }
</script>

<div class="item">
    <Tile>
        <div class="container">
            {#if skeleton}
                <span id="skeleton">
                    <SkeletonPlaceholder style="height: 2rem; width: 100%;" />
                </span>
            {:else}
                <div id="id">
                    {#if id === 0}
                        -
                    {:else}
                        {id}
                    {/if}
                </div>
                <span id="title" use:truncate>
                    <a href="/articles/{id}/">{title}</a>
                </span>
                <span id="size" class="hidden-mobile">
                    {@html formatSizeUnits(size)}
                </span>
            {/if}
            <span id="copy">
                <Button
                    tooltipPosition="left"
                    tooltipAlignment="end"
                    iconDescription="复制永久链接"
                    icon={CopyLink}
                    kind="ghost"
                    size="small"
                    on:click={() => copyLink(`/articles/${id}/`)}
                    disabled={skeleton}
                />
            </span>
        </div>
    </Tile>
</div>

<style>
    .container {
        display: flex;
        column-gap: 4px;
    }
    #skeleton {
        width: 100%;
    }
    .item {
        padding: 2px;
    }
    #copy {
        width: 32px;
    }
    #id {
        font-size: smaller;
        font-family: "JetBrains Mono", "JetBrainsMono NF", "Fira Code",
            "FiraCode NF", "Noto Sans Mono CJK SC", "Courier New", Courier,
            Consolas, monospace;
        text-align: center;
        min-width: 40px;
        float: left;
        background-color: --tertiary-color-light;
        margin-top: 1px;
        padding: 2px;
        padding-top: 7px;
    }
    #title {
        margin-left: 8px;
        margin-right: 8px;
        padding-top: 6px;
        vertical-align: text-bottom;
        flex-grow: 1;
        font-weight: 600;
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
        padding-top: 7.5px;
        min-width: 80px;
        text-align: right;
    }
    a {
        color: var(--accent-color-light);
        position: relative;
        transition: color 0.4s ease-out;
        text-decoration-color: var(--tertiary-color-light);
        /* mix-blend-mode: difference; */
        text-decoration-skip: ink;
    }

    a:hover {
        color: var(--accent-color-light-hover);
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
    @media (max-width: 767px) {
        .hidden-mobile {
            display: none;
        }
    }

    @media screen and (prefers-color-scheme: dark) {
        a {
            text-decoration-color: var(--tertiary-color-dark);
            color: var(--accent-color-dark);
        }

        #id {
            background-color: var(--tertiary-color-dark);
        }

        a:hover {
            color: var(--accent-color-dark-hover);
        }

        #size {
            color: #cccccc;
        }
    }
</style>
