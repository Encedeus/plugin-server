import type {User} from "@encedeus/registry-js-api";

export interface SearchLoadData {
    request: {
        query: string
        page: number
        perPage: number
        limit: number
    };

    response: {
        plugins: Plugin[],
        page: number,
        pages: number
    };
}

export interface LayoutAuthData {
    isAuth: boolean,
    user: User | undefined,
    accessToken: string | undefined
}