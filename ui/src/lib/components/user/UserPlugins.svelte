<script lang="ts">
    import {User} from "@encedeus/registry-js-api";
    import PluginList from "$lib/components/plugin/PluginList.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import Card from "$lib/components/generic/Card.svelte";

    export let user: User;
    export let isOwn: Boolean;
    export let className: string;

</script>


<div class={className}>

    <Card className="" >
        <div slot="title">
            Plugins
            {#if isOwn}
                <Button className="rounded-xl ml-8" size="md" redirect="/plugin/">Create New</Button>
            {/if}
        </div>
        <div slot="content" class="">
            <PluginList className="mt-2.5" plugins={user.plugins}/>
        </div>
    </Card>


</div>

<style>
</style>



