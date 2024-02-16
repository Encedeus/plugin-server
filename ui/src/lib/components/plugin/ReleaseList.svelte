<script lang="ts">
    import {Release} from "@encedeus/registry-js-api";
    import ReleaseInfo from "$lib/components/plugin/ReleaseInfo.svelte";
    import FilledContainer from "$lib/components/generic/FilledContainer.svelte";

    export let releases: Release[];
    export let isOwn: boolean = false;
    export let plugin: Plugin;
    export let className: string;
</script>

<div class="releaseListContainer {className}">
    {#if releases !== undefined}
        {#each releases as release}
            <FilledContainer>
                <ReleaseInfo className="m-auto" release={release} isOwn={isOwn} plugin={plugin}/>
            </FilledContainer>
        {/each}
    {:else}
        <p>no published releases</p>
    {/if}
</div>

<style>
    .releaseListContainer {
        display: flex;
        flex-direction: column;
        gap: 10px 10px;
    }

</style>