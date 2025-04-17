import {apiRequest} from "$lib/api/api.ts";
import {redirect, fail} from "@sveltejs/kit";

export const actions = {
    createSchool: async ({request, cookies}) => {
        const formData = await request.formData();
        const name = String(formData.get('name'));
        const description = String(formData.get('description'));

        const token = cookies.get('AuthorizationToken');
        const {data, error} = await apiRequest('/schools', 'post', {
            name: name,
            description: description,
        }, token)

        if (error) {
            return fail(error.status, {
                error: error.error,
                name: name,
                description: description,
            })
        }

        throw redirect(302, '/me/schools/' + data.id)
    }
}