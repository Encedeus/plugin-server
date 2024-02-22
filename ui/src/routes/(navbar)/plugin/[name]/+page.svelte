<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import PluginInfo from "$lib/components/plugin/BasicPluginInfo.svelte";
    import Readme from "$lib/components/plugin/Readme.svelte";
    import ReleaseList from "$lib/components/plugin/ReleaseList.svelte";
    import {onMount} from "svelte";
    import {isOwnPlugin} from "$lib/service/userService";
    import {page} from "$app/stores";
    import Button from "$lib/components/generic/Button.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import FilledContainer from "$lib/components/generic/FilledContainer.svelte";

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

<div id="page" class="m-auto w-[80rem]">
    <h1 class="text-4xl">{data.plugin.name}</h1>
    <FilledContainer className="w-fit m-2.5 ml-0 pl-2.5 pr-2.5">
        <PluginInfo plugin={data.plugin} label="owner"/>
    </FilledContainer>

    <div id="options">
        <Button onClick={handleReadmeClick} className="w-5/12" size="sm">Readme</Button>
        <Button onClick={handleReleasesClick} className="w-5/12" size="sm">Releases</Button>
    </div>

    <div class:invisible={selected !== "readme"}>
        <Card className="mt-2" headerTextSize="lg">
            <h1 slot="title" class="m-0">README.MD</h1>
            <Readme slot="content" className="mt-2.5 ml-2.5" plugin={data.plugin}/>
        </Card>


    </div>
    <div class:invisible={selected !== "releases"}>
        <Card className="mt-2" headerTextSize="lg">
            <div slot="title" class="flex flex-row justify-between">
                <h1 class="m-0">Published Releases</h1>
                {#if isOwn}
                        <Button className="w-36 rounded-xl mr-8" size="sm"
                                redirect={`/plugin/${data.plugin.name}/releases`}>Manage releases</Button>
                {/if}
            </div>
            <span slot="content">

                <ReleaseList className="mt-2.5 m-auto w-fit min-w-[400px]" releases={data.plugin.releases} isOwn={isOwn}
                             plugin={data.plugin}/>
            </span>
        </Card>


    </div>

</div>

<style>
    #options {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
    }

    .invisible {
        display: none;
    }

    #page {
        padding: 10px 50px;
    }
</style>

