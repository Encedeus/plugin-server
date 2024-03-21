<script lang="ts">
    import Input from "$lib/components/generic/Input.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import {getApi} from "$lib/api/api";
    import {PluginCreateRequest} from "@encedeus/registry-js-api/src/proto/plugin_api";
    import {goto} from "$app/navigation";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import Checkbox from "$lib/components/Checkbox.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";


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


    <main class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
        <Card>
            <CardHeader slot="title">
                Create plugin
            </CardHeader>
            <div class="flex flex-col gap-5 m-6" slot="content">
                <Input bind:value={name}
                       placeholder="plugin name"
                       label="plugin name"
                />
                <Input bind:value={githubRepository}
                       placeholder="GitHub repository link"
                       label="GitHub repository link"
                />

                <Button className="m-auto" onClick= {handleCreatePlugin}>Create plugin</Button>

                <ErrorTextBox className="m-2.5" bind:value={errorMessage}/>
            </div>

        </Card>
    </main>
</div>
