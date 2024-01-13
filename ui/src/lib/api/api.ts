import {get} from 'svelte/store';
import {EncedeusRegistryApi} from "@encedeus/registry-js-api";
import {accessTokenStore} from "$lib/stores/accessTokenStore";

function onAuth(accessToken: string) {
    accessTokenStore.set(accessToken);
}

export function getApi(): EncedeusRegistryApi {
    return new EncedeusRegistryApi("http://localhost:3001", get(accessTokenStore), {}, onAuth);
}

export function isAuthenticated(): boolean {
    console.log(get(accessTokenStore).length)
    console.log(get(accessTokenStore).length > 0)
    return get(accessTokenStore).length > 0;
}