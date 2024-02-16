<script lang="ts">
    import {Release, PluginDeprecateReleaseRequest, Plugin} from "@encedeus/registry-js-api";
    import {getRelativeTimeString} from "$lib/service/relativeTimeService";
    import {getApi} from "$lib/api/api";
    import {invalidateAll} from "$app/navigation";
    import Button from "$lib/components/generic/Button.svelte";
    import ElipsisVertical from "$lib/components/heroicons/ElipsisVertical.svelte";

    export let release: Release;
    export let isOwn: boolean = false;
    export let plugin: Plugin;
    export let className: string;

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
<div class="releaseInfoContainer w-fit {className}">
    <p class="releaseInfoComponent m-auto" class:deprecated={release.isDeprecated}>{release.name} ----- {relativeTime}</p>
    {#if isOwn}
        <button class="releaseInfoComponent" on:click={toggleDropdown}><ElipsisVertical/></button>
        <div class="releaseInfoComponent releaseDropdown" class:invisible={!isShowingDropdown}>
            <Button isDisabled={release.isDeprecated} size="xs" onClick={handeDeprecate}>Deprecate</Button>
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
        position: absolute;
    }
</style>