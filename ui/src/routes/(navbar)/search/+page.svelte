<script lang="ts">
    import PluginList from "$lib/components/plugin/PluginList.svelte";
    import PageSelector from "$lib/components/search/PageSelector.svelte";
    import type {SearchLoadData} from "$lib/interfaces/intefaces";

    export let data: SearchLoadData;

    let plugins, currentPage: number, pagesTotal: number;

    function getAnchor(page: number): string {
        let href: URL = new URL("/search", "http://localhost:5173");

        if (data.request.query) href.searchParams.set("q", data.request.query.toString());
        if (data.request.page) href.searchParams.set("page", page.toString());
        if (data.request.perPage) href.searchParams.set("perpage", data.request.perPage.toString());
        if (data.request.limit) href.searchParams.set("limit", data.request.limit.toString());

        return href.toString();
    }

    function loadData() {
        if (data.response.plugins.length == 0) {
            return;
        }

        plugins = data.response.plugins!;
        currentPage = data.response.page!;
        pagesTotal = data.response.pages!;
    }

    $: data && loadData();
    loadData();
</script>


<div id="page">
    {#if data.response.plugins.length !== 0}
        <PluginList plugins={plugins}/>
        <PageSelector className="m-auto w-fit" currentPage={currentPage}, pagesTotal={pagesTotal} getAnchor={getAnchor}/>
    {:else }
        <p>no matches found</p>
    {/if}
</div>

<style>
    #page {
        margin: auto;
        margin-top: 10px;
        max-width: 85rem;
    }
</style>