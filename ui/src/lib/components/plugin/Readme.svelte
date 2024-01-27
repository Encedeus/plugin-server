<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import {getApi} from "$lib/api/api";
    import {PluginGetReadmeRequest} from "@encedeus/registry-js-api/src/proto/plugin_api";
    import { marked } from 'marked'
    export let plugin: Plugin;
    let readme: string;
    async function loadData() {
        readme = ""

        const response = await getApi().PluginService.GetReadme({pluginId: plugin.id} as PluginGetReadmeRequest);

        if (response.error) {
            readme = response.error.message;
            return
        }

        // translate markdown to html
        readme = <string>marked(response.response?.readme!);
    }

    $: plugin && loadData()
    loadData()

</script>

<!-- todo: skeletal loading -->

<div>
    {@html readme}
</div>