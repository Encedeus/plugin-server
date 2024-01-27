<script lang="ts">
    import {Release} from "@encedeus/registry-js-api";
    import ReleaseInfo from "$lib/components/plugin/ReleaseInfo.svelte";

    export let releases: Release[];
    export let isOwn: boolean = false;
    export let plugin: Plugin;
</script>

{#if releases !== undefined}
    {#each releases as release}
        <ReleaseInfo release={release} isOwn={isOwn} plugin={plugin}/>
    {/each}
{:else}
    <p>no published releases</p>
{/if}