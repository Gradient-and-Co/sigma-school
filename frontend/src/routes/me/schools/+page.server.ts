import {error, redirect} from "@sveltejs/kit";
import {apiRequest} from "$lib/api/api.ts";

export const load = async ({cookies, locals}) => {
    const session = locals.session;
    if (!session) {
        throw redirect(302, '/auth/login')
    }

    const token = cookies.get('AuthorizationToken');
    const {data: schools, error: err} = await apiRequest('/users/me/schools', 'get', undefined, token);
    if (err) {
        throw error(500, {message: err.error})
    }

    return {schools, token};
};