<script lang="ts">
    import Input from "$lib/components/Input.svelte";
    import Button from "$lib/components/Button.svelte";
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
    <Input bind:value={name} helperText="pluign name"/>
    <Input bind:value={githubRepository} helperText="GitHub repository link"/>

    <Button onclick={handleCreatePlugin}>Create plugin</Button>

    <ErrorTextBox bind:value={errorMessage}/>
</div>

<style>
    #page {
        text-align: center;
    }
</style>
