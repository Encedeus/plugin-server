<script lang="ts">
    import ReleaseList from "$lib/components/plugin/ReleaseList.svelte";
    import {Plugin, PluginPublishReleaseRequest} from "@encedeus/registry-js-api";
    import {userDataStore} from "$lib/stores/userDataStore";
    import {goto, invalidateAll} from "$app/navigation";
    import Input from "$lib/components/generic/Input.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {getApi} from "$lib/api/api";
    import Card from "$lib/components/generic/Card.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";


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
    <main class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
        <Card>
            <CardHeader slot="title">
                Manage Releases
            </CardHeader>
            <div class="flex flex-col gap-5 m-6 " slot="content">
                <Input bind:value={releaseName}
                       label="release name"
                       placeholder="release name"
                />

                <Input bind:value={githubTag}
                       label="github tag"
                       placeholder="github tag"
                />

                <Button className="m-auto" onClick={handleCreateRelease}>Publish</Button>

                <ErrorTextBox className="m-auto" bind:value={errorMessage}/>

                <ReleaseList className="m-auto" releases={plugin.releases} isOwn={isOwn} plugin={plugin}/>
            </div>

        </Card>
    </main>


</div>

<style lang="postcss">
    #page {
        max-width: 400px;
        text-align: center;
        margin: auto;
    }
</style>