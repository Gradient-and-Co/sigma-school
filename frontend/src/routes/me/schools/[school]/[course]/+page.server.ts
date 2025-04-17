import {apiRequest} from "$lib/api/api.ts";
import {redirect, fail, error} from "@sveltejs/kit";

export const load = async ({cookies, locals, params}) => {
    const session = locals.session;
    if (!session) {
        throw redirect(302, '/auth/login')
    }

    const token = cookies.get('AuthorizationToken');
    const {data: course, error: err} = await apiRequest('/courses/' + params.course + '/brief', 'get', undefined, token);
    if (err) {
        throw error(500, {message: err.error})
    }

    const {data: school, error: schoolError} = await apiRequest('/schools/' + params.school, 'get', undefined, token);
    if (schoolError) {
        throw error(500, {message: schoolError.error})
    }

    return {course, school, token};
};

export const actions = {
    editCourse: async ({request, cookies, params}) => {
        const formData = await request.formData();
        const name = String(formData.get('name'));
        const description = String(formData.get('description'));
        const price = Number(String(formData.get('price')));
        const language = String(formData.get('language'));

        if (name == '') {
            return fail(400, {
                error: 'name cannot be empty',
            })
        }

        const token = cookies.get('AuthorizationToken');
        const {error} = await apiRequest('/schools/' + params.school + '/courses/' + params.course, 'patch', {
            name: name,
            description: description,
            level: 1,
            price: price,
            language: language,
        }, token)

        if (error) {
            return fail(error.status, {
                error: error.error,
                name: name,
                description: description,
                price: price,
                language: language,
            })
        }

        throw redirect(302, '/me/schools/' + params.school)
    }
}