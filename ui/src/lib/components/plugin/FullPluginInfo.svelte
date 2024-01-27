<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import PluginName from "$lib/components/plugin/PluginName.svelte";
    import PluginInfo from "$lib/components/plugin/BasicPluginInfo.svelte";

    export let plugin: Plugin;

</script>

<div class="pluginContainer">
    <PluginName pluginName={plugin.name}/>
    <PluginInfo plugin={plugin}/>
</div>

<style>
    .pluginContainer {
        text-align: left;
    }
</style>