<script lang="ts">
    import {Release, PluginDeprecateReleaseRequest, Plugin} from "@encedeus/registry-js-api";
    import {getRelativeTimeString} from "$lib/service/relativeTimeService";
    import {getApi} from "$lib/api/api";
    import {invalidateAll} from "$app/navigation";

    export let release: Release;
    export let isOwn: boolean = false;
    export let plugin: Plugin;

    let relativeTime: string = getRelativeTimeString(Date.parse(<string>release.publishedAt));
    let isShowingDropdown: boolean = false;


    function toggleDropdown() {
        isShowingDropdown = !isShowingDropdown;
    }

    async function handeDeprecate() {
        const api = getApi();

        const req = {
            releaseName: release.name,
            pluginId: plugin.id
        } as PluginDeprecateReleaseRequest;

        const resp = await api.PluginService.DeprecateRelease(req);

        if (resp.error) {
            console.log("deprecation failed");
            return
        }

        await invalidateAll()
    }
</script>
<div class="releaseInfoContainer">
    <p class="releaseInfoComponent" class:deprecated={release.isDeprecated}>{release.name} ----- {relativeTime}</p>
    {#if isOwn}
        <button class="releaseInfoComponent" on:click={toggleDropdown}>...</button>
        <div class="releaseInfoComponent releaseDropdown" class:invisible={!isShowingDropdown}>
            <button on:click={handeDeprecate}>Deprecate</button>
        </div>
    {/if}
</div>


<style>
    .releaseInfoComponent {
        display: inline;
    }

    .deprecated {
        color: red;
    }

    .invisible {
        display: none;
    }

    .releaseDropdown {
        background-color: darkgray;
        position: relative;
    }
</style>