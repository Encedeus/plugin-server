import {getApi} from "$lib/api/api";
import type {LayoutAuthData} from "$lib/interfaces/intefaces";
import type {User} from "@encedeus/registry-js-api";

export async function load({cookies}: any) {
    const token = cookies.get("encedeus_plugins_refreshToken")

    if (!token) {
        return
    }

    const api = getApi(true, token)

    const refreshResp = await api.AuthService.RefreshAccessToken()

    if (refreshResp.error) {
        return {isAuth: false} as LayoutAuthData
    }

    const getSelfResp = await api.UsersService.GetSelf()

    if (getSelfResp.error) {
        return {isAuth: false} as LayoutAuthData
    }

    const user: User = getSelfResp.response as User
    const accessToken: string = refreshResp.response?.accessToken!

    return {
        isAuth: true,
        user: user,
        accessToken: accessToken
    } as LayoutAuthData
}