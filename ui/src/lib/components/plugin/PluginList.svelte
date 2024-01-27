<script lang="ts">


    import Plugin from "$lib/components/plugin/FullPluginInfo.svelte";

    export let plugins: Plugin[];
</script>

{#each plugins as plugin}
    <Plugin plugin={plugin}/>
{/each}