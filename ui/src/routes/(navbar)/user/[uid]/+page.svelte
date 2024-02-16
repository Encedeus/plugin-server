<script lang="ts">
    import UserInfo from "$lib/components/user/UserInfo.svelte";
    import UserPlugins from "$lib/components/user/UserPlugins.svelte";
    import {userDataStore} from "$lib/stores/userDataStore";

    /** @type {import('./$types').PageData} */
    export let data: any;

    // data from server
    let user = data.user;
    const isOwn: Boolean = $userDataStore ? $userDataStore.id === user.id : false;
</script>


<div id="page" class="m-auto">
    <div class="info">
        <UserInfo className=""  user={user} isOwn={isOwn}/>
        <UserPlugins user={user} isOwn={isOwn}/>
    </div>
</div>

<style>
    #page {
        margin: auto;
        margin-top: 20px;
        max-width: 800px;
    }

    .info {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-between;

        gap: 5px 50px;
    }


</style>


