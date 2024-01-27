<script lang="ts">
    import {getApi, isAuthenticated} from "$lib/api/api";
    import {userDataStore} from "$lib/stores/userDataStore";
    import {onMount} from 'svelte';
    import {GetUserPageURL} from "$lib/service/userService";
    import {PluginSearchByNameRequest} from "@encedeus/registry-js-api/src/proto/plugin_api";
    import {Plugin} from "@encedeus/registry-js-api";
    import PluginInfo from "$lib/components/plugin/BasicPluginInfo.svelte";
    import {navigating, page} from "$app/stores";
    import {goto} from "$app/navigation";

    const api = getApi();

    let isAuth: boolean,
        pfpURL: string,
        userPageURL: string,
        inputInFocus: boolean = false,
        searchValue: string = "",
        plugins: Plugin[] = [],
        mouseOverSuggestions: boolean = false;

    async function loadData() {


        const isAuthEstablished = isAuthenticated();

        if (isAuthEstablished == isAuth) {
            return
        }

        isAuth = isAuthEstablished

        if (isAuth) {
            if (!$userDataStore) {
                await api.UsersService.GetSelf();
            }
            pfpURL = api.UsersService.GetUserPfpURL($userDataStore);
            userPageURL = GetUserPageURL($userDataStore);
        }
    }

    function onInputFocus() {
        inputInFocus = true;
    }

    function onInputFocusOut() {
        inputInFocus = false;
    }

    function onMouseEnterSuggestionBox() {
        mouseOverSuggestions = true;
    }

    function OnMouseLeaveSuggestionBox() {
        mouseOverSuggestions = false;
    }

    async function getSearchSuggestions() {
        const req = {
            name: searchValue,
            limit: 5
        } as PluginSearchByNameRequest;

        // reset suggestions in 20 ms
        const timeout = setTimeout(() => {
            plugins = [];
        }, 20);

        const res = await api.PluginService.SearchPlugins(req);

        // cancel suggestion reset if request resolved in less the 20 ms
        clearTimeout(timeout);

        if (res.error) {
            plugins = [];
            return;
        }

        plugins = res.response?.plugins!;
    }

    function handleInputSubmit(event: KeyboardEvent) {
        if (event.key != "Enter") {
            return;
        }

        goto(`/search/?q=${searchValue}`, {invalidateAll: true});
    }

    onMount(loadData);
    $: $userDataStore, $page.url.pathname && loadData();
</script>

<header>
    <nav>
        <a href="/">
            <img id="logo" src="/logo.png" alt="Encedeus logo">
        </a>

        <div class="navbarComponent" id="searchBarContainer">
            <input on:focusin={onInputFocus}
                   on:focusout={onInputFocusOut}
                   on:keyup={() =>{onInputFocus(); getSearchSuggestions()}}
                   on:keyup={handleInputSubmit}
                   bind:value={searchValue}
                   type="search" id="searchBar"
                   class="navbarComponent"
                   placeholder="Search plugins"
            />

            <div on:mouseenter={onMouseEnterSuggestionBox}
                 on:mouseleave={OnMouseLeaveSuggestionBox}
                 class:invisible={!(inputInFocus || mouseOverSuggestions)}
                 class="suggestions"
                 id="searchSuggestions"
            >

                {#if plugins.length > 0}
                    {#each plugins as plugin}
                        <PluginInfo plugin={plugin} label="name"/>
                    {/each}
                {:else}
                    <p>...</p>
                {/if}
            </div>
        </div>

        <div>
            <a href="/documentation">
                <button id="docs" class="navbarComponent">Documentation</button>
            </a>
        </div>

        {#if !isAuth}
            <div id="authAnchors">
                <a href="/auth/signup" id="signUp" class="navbarComponent">Sign Up</a>

                <a href="/auth/login" id="signIn" class="navbarComponent">Sign In</a>
            </div>
        {:else}
            <!-- todo: implement lazy loading -->
            <a href={userPageURL}> <img id="pfp" class="navbarComponent" src={pfpURL}> </a>
        {/if}
    </nav>


</header>

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

    #searchBarContainer {
        min-width: 20%;
        width: 50%;
        border: 0;
        padding: 0;
    }

    #searchBar {
        width: 100%;
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

    #pfp {
        border-radius: 5px;
    }

    .invisible {
        display: none;
    }

    .suggestions {
        position: relative;
        background-color: darkgray;
        height: max-content;
    }
</style>