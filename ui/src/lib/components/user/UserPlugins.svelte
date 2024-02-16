<script lang="ts">
    import {User} from "@encedeus/registry-js-api";
    import PluginList from "$lib/components/plugin/PluginList.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import Card from "$lib/components/generic/Card.svelte";

    export let user: User;
    export let isOwn: Boolean;

</script>


<div>

    <Card headerTextSize="lg" >
        <h1 slot="title" class="">Plugins</h1>

        <div slot="content" class="m-2.5">

            <div class="flex justify-center">
                <Button className="" size="sm" redirect="/plugin/">Create New</Button>
            </div>
            <PluginList className="mt-2.5" plugins={user.plugins}/>
        </div>
    </Card>


</div>

<style>
</style>



