<script lang="ts">
    import {User} from "@encedeus/registry-js-api";
    import PluginList from "$lib/components/plugin/PluginList.svelte";

    export let user: User;
    export let isOwn: Boolean;
</script>

<div>
    <h1>Plugins</h1>
    <a href="/plugin/">Create new</a>
    <PluginList plugins={user.plugins}/>
</div>



