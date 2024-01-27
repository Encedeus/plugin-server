<script lang="ts">
    import {User} from "@encedeus/registry-js-api";
    import {getApi} from "$lib/api/api";
    import {goto} from "$app/navigation";

    export let user: User;
    export let isOwn: Boolean;

    const api = getApi();
    let userPfpURL = api.UsersService.GetUserPfpURL(user);

    async function handleLogOut() {
        const api = getApi()
        await api.AuthService.SignOut()
        goto("/")
    }

</script>

<div>
    <div>
        <img class="pfp" src={userPfpURL} alt="user profile picture">
        <h1 class="userNameBig">{user.name}</h1>
    </div>

    {#if isOwn}
        <a href="/user/edit">Edit Profile</a>
        <button on:click={handleLogOut}>Log out</button>
    {/if}
</div>

<style>
    .pfp{
        min-height: 200px;
        min-width: 200px;
        width: 200px;
        height: 200px;

        border-radius: 22px;
    }

    .userNameBig {
        text-align: left;
    }
</style>
