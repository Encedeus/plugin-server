<script lang="ts">
    import {getApi, isAuthenticated} from "$lib/api/api";
    import {userDataStore} from "$lib/stores/userDataStore";

    const api = getApi()

    let isAuth: boolean = isAuthenticated();

    if (isAuth) await api.UsersService.GetSelf()

    let pfpURL: string = api.UsersService.GetUserPfpURL($userDataStore.id)

</script>

<nav>
    <a href="/">
        <img id="logo" src="logo.png" alt="Encedeus logo">
    </a>


    <input id="searchBar" class="navbarComponent" type="search" placeholder="Search plugins">

    <div>
        <a href="./documentation">
            <button id="docs" class="navbarComponent">Documentation</button>
        </a>
    </div>

    {#if !isAuth}
        <div id="authAnchors">
            <a href="./auth/signup" id="signUp" class="navbarComponent">Sign Up</a>

            <a href="./auth/login" id="signIn" class="navbarComponent">Sign In</a>
        </div>
    {:else}
        <img src={pfpURL} alt="user profile picture">
    {/if}
</nav>

<style>
    nav {
        top: 0;
        left: 0;
        margin-top: 0;
        padding-bottom: 5px;

        position: relative;

        background-color: gray;

        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;

        gap: 5px 20px;
    }

    #logo,
    #searchBar,
    #authAnchors {
        max-height: 100px;
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        gap: inherit;
    }

    #searchBar {
        min-width: 20%;
        width: 50%;
        border: 0;
        padding: 0;
    }

    button,
    #signIn,
    #signUp {
        max-height: 100px;
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;

        border: 0;
        border-bottom: 50px;
    }

    #signIn,
    #signUp {
        background-color: aliceblue;
        padding-inline: 4px 4px;
    }

    .navbarComponent {
        height: 45px;
    }

</style>