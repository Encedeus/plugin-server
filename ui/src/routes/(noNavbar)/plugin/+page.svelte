<script lang="ts">
    import Input from "$lib/components/generic/Input.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import {getApi} from "$lib/api/api";
    import {PluginCreateRequest} from "@encedeus/registry-js-api/src/proto/plugin_api";
    import {goto} from "$app/navigation";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";


    let name: string = "",
        githubRepository: string = "",
        errorMessage: string = "",
        isLoading: boolean = false;

    const api = getApi()
    async function handleCreatePlugin() {
        const req = {
            name: name,
            repoUri: githubRepository
        } as PluginCreateRequest

        isLoading = true

        const resp = await api.PluginService.CreatePlugin(req)

        isLoading = false

        if (resp.error){
            errorMessage = resp.error.message
            return
        }

        goto(`/plugin/${req.name}`)
    }

</script>
<div id="page">
    <Input bind:value={name}
           placeholder="plugin name"
           label="plugin name"
    />
    <Input bind:value={githubRepository}
           placeholder="GitHub repository link"
           label="GitHub repository link"
    />

    <ErrorTextBox className="m-2.5" bind:value={errorMessage}/>

    <Button className="" onClick= {handleCreatePlugin}>Create plugin</Button>
</div>

<style>
    #page {
        text-align: center;
    }
</style>
