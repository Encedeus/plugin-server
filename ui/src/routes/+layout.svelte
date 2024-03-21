<script lang="ts">
    import {userDataStore} from "$lib/stores/userDataStore";
    import type {LayoutAuthData} from "$lib/interfaces/intefaces";
    import {accessTokenStore} from "$lib/stores/accessTokenStore";
    import "../app.css";

    export let data: LayoutAuthData;

    if (data.isAuth) {
        userDataStore.set(data.user!)
        accessTokenStore.set(data.accessToken!)
    }
</script>

<main class="h-full">
    <slot/>
</main>