<script lang="ts">
    import Plugin from "$lib/components/plugin/FullPluginInfo.svelte";
    import TextArea from "$lib/components/generic/TextArea.svelte";
    import FilledContainer from "$lib/components/generic/FilledContainer.svelte";

    export let plugins: Plugin[],
        className: string;
</script>


<div class="pluginList {className}">
    {#if plugins}
        {#each plugins as plugin}
            <FilledContainer>
                <Plugin className="ml-2 mb-1" plugin={plugin}/>
            </FilledContainer>
        {/each}
    {/if}
</div>

<style>
    .pluginList {
        display: flex;
        flex-direction: column;
        gap: 10px;
    }
</style>