<script lang="ts">
    import ReleaseList from "$lib/components/plugin/ReleaseList.svelte";
    import {Plugin, PluginPublishReleaseRequest} from "@encedeus/registry-js-api";
    import {userDataStore} from "$lib/stores/userDataStore";
    import {goto, invalidateAll} from "$app/navigation";
    import Input from "$lib/components/generic/Input.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {getApi} from "$lib/api/api";


    export let data: Plugin;
    let plugin: Plugin,
        isOwn: boolean = false,
        releaseName: string = "",
        githubTag: string = "",
        errorMessage: string = "",
        isLoading: boolean = false;

    async function handleCreateRelease() {
        const req = {
            pluginId: plugin.id,
            name: releaseName,
            githubReleaseTag: githubTag,
        } as PluginPublishReleaseRequest;

        const api = getApi();


        const resp = await api.PluginService.CreatePluginRelease(req);

        if (resp.error) {
            errorMessage = resp.error.message;
            return;
        }

        isLoading = true;
        await invalidateAll();
        isLoading = false;
    }

    function loadData() {
        plugin = data;
        if (!$userDataStore || plugin.ownerName != $userDataStore.name) {
            isOwn = false;
            goto("/");
            return;
        }

        isOwn = true;
    }

    $: data && loadData();
</script>

<div id="page">

    <Input bind:value={releaseName}
           label="release name"
           placeholder="release name"
    />

    <Input bind:value={githubTag}
           label="github tag"
           placeholder="github tag"
    />

    <ErrorTextBox className="m-2.5" bind:value={errorMessage}/>

    <Button className="mb-2.5" onClick={handleCreateRelease}>Publish</Button>

    <ReleaseList className="m-auto" releases={plugin.releases} isOwn={isOwn} plugin={plugin}/>
</div>

<style lang="postcss">
    #page {
        max-width: 400px;
        text-align: center;
        margin: auto;
    }
</style>