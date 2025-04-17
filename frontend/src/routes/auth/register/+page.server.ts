import {apiRequest} from "$lib/api/api.ts";
import {redirect, fail} from "@sveltejs/kit";

export const actions = {
    register: async ({request}) => {
        const formData = await request.formData();
        const email = String(formData.get('email'));
        const password = String(formData.get('password'));
        const name = String(formData.get('name'));
        const surname = String(formData.get('surname'));
        const phone = String(formData.get('phone'));
        const avatarUrl = String(formData.get('avatar_url'));

        const {error} = await apiRequest('/auth/sign-up', 'post', {
            name: name,
            surname: surname,
            phone: phone,
            email: email,
            password: password,
            avatar_url: avatarUrl,
        })

        if (error) {
            return fail(error.status, {
                error: error.error,
                name: name,
                surname: surname,
                phone: phone,
                email: email,
                avatar_url: avatarUrl,
            })
        }

        throw redirect(303, '/auth/login')
    }
}