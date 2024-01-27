<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import PluginInfo from "$lib/components/plugin/BasicPluginInfo.svelte";
    import Readme from "$lib/components/plugin/Readme.svelte";
    import ReleaseList from "$lib/components/plugin/ReleaseList.svelte";
    import {onMount} from "svelte";
    import {isOwnPlugin} from "$lib/service/userService";
    import {page} from "$app/stores";

    /** @type {import('./$types').PageData} */
    export let data: { plugin: Plugin };

    // local
    let selected: "readme" | "releases" = "readme",
        isOwn: boolean = false;

    function loadData() {
        isOwn = isOwnPlugin(data.plugin);
    }

    function handleReadmeClick() {
        selected = "readme";
    }

    function handleReleasesClick() {
        selected = "releases";
    }

    onMount(loadData);
    $: data , $page.url.pathname && loadData();
</script>

<div id="page">
    <h1>{data.plugin.name}</h1>
    <PluginInfo plugin={data.plugin} label="owner"/>
    <a href={`/plugin/${data.plugin.name}/releases`}>Manage Releases</a>

    <br>
    <br>

    <div id="options">
        <button on:click={handleReadmeClick}>Readme</button>
        <button on:click={handleReleasesClick}>Releases</button>
    </div>

    <div class:invisible={selected !== "readme"}>
        <Readme plugin={data.plugin}/>
    </div>
    <div class:invisible={selected !== "releases"}>
        <ReleaseList releases={data.plugin.releases} isOwn={isOwn} plugin={data.plugin}/>
    </div>

</div>

<style>
    h1 {
        margin-bottom: 5px;
    }

    .invisible {
        display: none;
    }
</style>

