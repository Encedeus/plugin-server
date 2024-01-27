import {get} from "svelte/store";
import {getApi} from "$lib/api/api";
import type {PluginSearchByNameRequest} from "@encedeus/registry-js-api/src/proto/plugin_api";
import {error} from "@sveltejs/kit";
import type {SearchPluginsByNameResponse} from "@encedeus/registry-js-api/src/services/pluginService";
import type {SearchLoadData} from "$lib/interfaces/intefaces";



/** @type {import('./$types').PageServerLoad} */
export async function load({params, url}: any) {
    const api = getApi();

    const name = url.searchParams.get("q"),
        page = url.searchParams.get("page"),
        pluginsPerPage = url.searchParams.get("perpage"),
        limit = url.searchParams.get("limit");

    const req = {name, limit, page, pluginsPerPage} as PluginSearchByNameRequest;

    const resp = await api.PluginService.SearchPlugins(req);

    let plugins = resp.response?.plugins
    let pagesTotal = resp.response?.pages
    let currentPage = resp.response?.page

    if (resp.error) {
        if (resp.error.statusCode == 500) {
            error(500, "Server fucked up");
        }

        plugins = []
        pagesTotal = 0
        currentPage = 0
    }

    return {
        request: {
            query: name,
            limit: limit,
            perPage: pluginsPerPage,
            page: page
        },
        response: {
            plugins: plugins,
            page: currentPage,
            pages: pagesTotal
        }
    } ;
}