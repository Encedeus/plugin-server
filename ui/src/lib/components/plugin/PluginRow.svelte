<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import XMark from "$lib/components/heroicons/XMark.svelte";
    import {getRelativeTimeString} from "$lib/service/relativeTimeService";
    import PluginName from "$lib/components/plugin/PluginName.svelte";
    import UserName from "$lib/components/user/UserName.svelte";

    export let plugin: Plugin;
    let latestRelativeTime: string = "";
    if (plugin.releases) {
        latestRelativeTime = getRelativeTimeString(Date.parse(<string>plugin.releases[0].publishedAt));
    }
</script>

<tr class="border-y-2 border-indigo-900">
    <td class="name min-w-full"><PluginName pluginName={plugin.name}/></td>
    <td class="author min-w-full "><UserName username={plugin.ownerName}/></td>

    {#if plugin.releases}
        <td class="latestRelease min-w-full">{plugin.releases[0].name}</td>
        <td class="latestReleaseTime min-w-full">{latestRelativeTime}</td>
    {:else }
        <td class="latestRelease min-w-full">no releases</td>
        <td class="latestReleaseTime min-w-full flex flex-row justify-center">
            <XMark/>
        </td>
    {/if}
</tr>

<style>
    td {
        /*padding-left: 10px;
        padding-right: 10px;*/
    }

    td:not(:first-child) {
        padding-top: 1rem;
        padding-bottom: 1rem;
    }
</style>