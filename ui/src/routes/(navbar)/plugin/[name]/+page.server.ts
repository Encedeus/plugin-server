import {Plugin} from "@encedeus/registry-js-api";
import {error} from "@sveltejs/kit";
import {getApi} from "$lib/api/api";

/** @type {import('./$types').PageServerLoad} */
export async function load({params}: any) {
    const api = getApi();

    const apiResp = await api.PluginService.GetPlugin(params.name);
    if (apiResp.error) {
        if (apiResp.error.statusCode == 404) {
            error(404, "Not found");
            return;
        }
        error(500, "Server fucked up");
        return;
    }

    const plugin: Plugin = apiResp.response!;
    return {plugin: plugin as Plugin};
}