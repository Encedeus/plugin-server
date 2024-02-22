<script lang="ts">
    import Plugin from "$lib/components/plugin/FullPluginInfo.svelte";
    import TextArea from "$lib/components/generic/TextArea.svelte";
    import FilledContainer from "$lib/components/generic/FilledContainer.svelte";
    import PluginRow from "$lib/components/plugin/PluginRow.svelte";

    export let plugins: Plugin[],
        className: string;
</script>

<div class="pb-4 text-white w-full {className}">
    <table class="border-collapse">
        <thead>
        <tr class="tableRowNames text-xs">
            <th scope="col">Name</th>
            <th scope="col">Author</th>
            <th scope="col">Latest release</th>
            <th scope="col">Latest Release Time</th>
        </tr>
        </thead>
        {#if plugins}
            {#each plugins as plugin}
                <PluginRow plugin={plugin}/>
            {/each}
        {/if}

    </table>
</div>

<style>
    table {
        text-align: center;
        table-layout: fixed;
        width: 100%;
        border-collapse: collapse;
    }

    thead th {
        width: 7.5%;
    }
</style>