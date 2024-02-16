<script lang="ts">
    import {User} from "@encedeus/registry-js-api";
    import {getApi} from "$lib/api/api";
    import {goto} from "$app/navigation";
    import Button from "$lib/components/generic/Button.svelte";

    export let user: User;
    export let isOwn: Boolean;
    export let className: string;

    const api = getApi();
    let userPfpURL = api.UsersService.GetUserPfpURL(user);

    async function handleLogOut() {
        const api = getApi();
        await api.AuthService.SignOut();
        goto("/");
    }

</script>

<div class={className}>
    <div>
        <img class="pfp" src={userPfpURL} alt="user profile picture">

        <h1 class="userNameBig font-bold text-3xl">{user.name}</h1>
    </div>

    {#if isOwn}
        <div class="profileButtons">
            <Button redirect="/user/edit" size="sm">Edit Profile</Button>
            <Button onClick={handleLogOut} size="sm">Log Out</Button>
        </div>
    {/if}
</div>

<style lang="css">
    .pfp {
        min-height: 200px;
        min-width: 200px;
        width: 200px;
        height: 200px;

        border-radius: 22px;
    }

    .userNameBig {
        text-align: left;
        container-type: inline-size;
        max-width: 200px;
    }

    .profileButtons {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }
</style>
