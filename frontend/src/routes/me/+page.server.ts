import {redirect} from "@sveltejs/kit";
import {apiRequest} from "$lib/api/api.ts";
import { error } from '@sveltejs/kit';

export const load = async ({locals}) => {
    const session = locals.session;
    if (!session) {
        throw redirect(302, '/auth/login')
    }

    const {data: user, error: err} = await apiRequest('/users/' + session.userID);
    if (err) {
        throw error(500, {message: err.error})
    }
    console.log(user, err)

    return {user};
};

export const actions = {
    logout: async ({cookies, locals}) => {
        cookies.delete('AuthorizationToken', { path: '/' });
        locals.session = undefined;
        throw redirect(302, '/auth/login');
    }
}
