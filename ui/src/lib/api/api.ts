import {get} from 'svelte/store';
import {EncedeusRegistryApi, type User} from "@encedeus/registry-js-api";
import {accessTokenStore} from "$lib/stores/accessTokenStore";
import {userDataStore} from "$lib/stores/userDataStore";

function onAuthUpdate(accessToken: string) {
    accessTokenStore.set(accessToken);
    getApi().UsersService.GetSelf().catch(err => err.response);
}

function onUserUpdate(user: User) {
    userDataStore.set(user);
}

export function getApi(noCallbacks: boolean = false, refreshCookie: string | undefined = undefined): EncedeusRegistryApi {
    return new EncedeusRegistryApi("http://localhost:3001", get(accessTokenStore), {
        axiosConfig: {
            withCredentials: true,
            headers: refreshCookie ? {Cookie: `encedeus_plugins_refreshToken=${refreshCookie};`} : {}
        },
        callbacks: {
            onAuth: noCallbacks ? () => {} : onAuthUpdate,
            onUser: noCallbacks ? () => {} : onUserUpdate
        }
    });
}

export function isAuthenticated(): boolean {
    return get(accessTokenStore).length > 0;
}