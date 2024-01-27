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


<div id="page">
    <div class="info">
        <UserInfo  user={user} isOwn={isOwn}/>
        <UserPlugins user={user} isOwn={isOwn}/>
    </div>
</div>

<style>
    #page {
        margin-top: 20px;
        margin-left: 10px
    }

    .info {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: left;

        gap: 5px 50px;
    }


</style>


