import {apiRequest} from "$lib/api/api.ts";
import {redirect, fail} from "@sveltejs/kit";

export const actions = {
    createCourse: async ({request, cookies, params}) => {
        const formData = await request.formData();
        const name = String(formData.get('name'));
        const description = String(formData.get('description'));
        const price = Number(String(formData.get('price')));
        const language = String(formData.get('language'));
        const imageUrl = String(formData.get('image_url'));

        const token = cookies.get('AuthorizationToken');
        const {error} = await apiRequest('/schools/' + params.school + '/courses', 'post', {
            name: name,
            description: description,
            level: 1,
            price: price,
            language: language,
            image_url: imageUrl,
        }, token)

        if (error) {
            return fail(error.status, {
                error: error.error,
                name: name,
                description: description,
                price: price,
                language: language,
                image_url: imageUrl,
            })
        }

        throw redirect(302, '/me/schools/' + params.school)
    }
}