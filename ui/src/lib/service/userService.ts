import type {User} from "@encedeus/registry-js-api";
import {userDataStore} from "$lib/stores/userDataStore";
import {get} from "svelte/store";
import {Plugin} from "@encedeus/registry-js-api";

export function GetUserPageURL(user: User): string {
    return `/user/${user.id}`;
}

export function isOwnPlugin(plugin: Plugin): boolean {
    if (!plugin) {
        return false;
    }

    const user = get(userDataStore);

    if (!user) {
        return false;
    }

    const plugins = user.plugins;

    for (let i = 0; i < plugins.length; i++) {
        if (plugins[i].id === plugin.id) {
            return true;
        }
    }
    return false;
}