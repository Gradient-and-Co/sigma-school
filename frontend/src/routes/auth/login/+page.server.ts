import {fail, redirect} from "@sveltejs/kit";
import {setAuthToken} from "$lib/utils/helpers.ts";
import {apiRequest} from "$lib/api/api.ts";
import {AUTH_FINGERPRINT} from "$env/static/private"

export const actions = {
    login: async ({cookies, request}) => {
        const formData = Object.fromEntries(await request.formData());
        const {email, password} = formData;

        const {data: token, error} = await apiRequest('/auth/sign-in', 'post', {
            email: email,
            password: password,
            fingerprint: AUTH_FINGERPRINT
        });
        if (error) {
            return fail(error.status, {
                error: error.error,
                email: email,
            });
        }

        console.log("logged in: " + token)

        setAuthToken({cookies, token});
        throw redirect(302, "/me")
    }
}