import { writable } from "svelte/store";
import {User} from "@encedeus/registry-js-api"
export const userDataStore = writable<User>(undefined);