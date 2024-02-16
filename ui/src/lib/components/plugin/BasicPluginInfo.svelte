<script lang="ts">
    import UserName from "$lib/components/user/UserName.svelte";
    import {getRelativeTimeString} from "$lib/service/relativeTimeService.js";
    import PluginName from "$lib/components/plugin/PluginName.svelte";
    import {goto} from "$app/navigation";

    export let plugin;
    export let className: string;
    export let label: "owner" | "name";

    const pluginPage = `/plugin/${plugin.name}/`;

    let hasReleases = plugin.releases !== undefined;
    let latestReleaseName: string = "", latestRelativeTime: string = "";

    function loadData() {
        hasReleases = plugin.releases !== undefined;

        if (hasReleases) {
            latestReleaseName = plugin.releases[0].name;
            latestRelativeTime = getRelativeTimeString(Date.parse(<string>plugin.releases[0].publishedAt));
        }
    }

    loadData();
    $: plugin && loadData();
</script>

<div class="pluginInfo {className}" >
    {#if label === "owner"}
        <UserName style="margin-top: 5px; margin-bottom: 5px" username={plugin.ownerName}/>
    {:else if label === "name"}
        <PluginName style="margin-top: 5px; margin-bottom: 5px" pluginName={plugin.name}/>
    {/if}


    {#if hasReleases}
        <p class="pluginInfoComponent">published {latestReleaseName}</p>
        <p class="pluginInfoComponent">{latestRelativeTime}</p>
    {:else}
        <p class="pluginInfoComponent">no published releases</p>
    {/if}

</div>

<style>
    .pluginInfo {
        display: flex;
        flex-direction: row;
        justify-content: left;
        /*width: 100%;*/
        width: fit-content;
        gap: 5px 20px;
    }

    p {
        margin-top: 5px;
    }
</style>