import {User} from "@encedeus/registry-js-api";
import {error} from "@sveltejs/kit";
import {getApi} from "$lib/api/api";

/** @type {import('./$types').PageServerLoad} */
export async function load({params}: any) {
    const api = getApi();

    const apiResp = await api.UsersService.GetUser(params.uid);
    if (apiResp.error) {
        if (apiResp.error.statusCode == 404) {
            error(404, "Not found");
            return;
        }
        error(500, "Server fucked up");
        return;
    }

    const user: User = apiResp.response!;
    return {user: user as User};
}