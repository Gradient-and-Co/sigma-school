import {error} from "@sveltejs/kit";
import {apiRequest} from "$lib/api/api.ts";

export const load = async ({params, cookies}) => {
    const token = cookies.get('AuthorizationToken');
    const {data: courses, error: err} = await apiRequest('/schools/' + params.school + '/courses',
        'get', undefined, token);
    if (err) {
        throw error(500, {message: err.error})
    }

    const {data: school, error: schoolError} = await apiRequest('/schools/' + params.school,
        'get', undefined, token);
    if (schoolError) {
        throw error(500, {message: schoolError.error})
    }

    return {courses, school, token};
};