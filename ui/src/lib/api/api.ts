import {get} from 'svelte/store';
import {EncedeusRegistryApi, type User} from "@encedeus/registry-js-api";
import {accessTokenStore} from "$lib/stores/accessTokenStore";
import {userDataStore} from "$lib/stores/userDataStore";

function onAuth(accessToken: string) {
    accessTokenStore.set(accessToken);
}

function onUser(user: User) {
    userDataStore.set(user);
}

export function getApi(): EncedeusRegistryApi {
    return new EncedeusRegistryApi("http://localhost:3001", get(accessTokenStore), {
        axiosConfig: {},
        callbacks: {
            onAuth,
            onUser
        }
    });
}

export function isAuthenticated(): boolean {
    console.log(get(accessTokenStore).length)
    console.log(get(accessTokenStore).length > 0)
    return get(accessTokenStore).length > 0;
}