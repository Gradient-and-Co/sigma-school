import {error, redirect} from "@sveltejs/kit";
import {apiRequest} from "$lib/api/api.ts";

export const load = async ({locals, cookies}) => {
    const session = locals.session;
    if (!session) {
        throw redirect(302, '/auth/login')
    }

    const token = cookies.get('AuthorizationToken');
    const {data: courses, error: err} = await apiRequest('/users/me/courses', 'get', undefined, token);
    if (err) {
        throw error(500, {message: err.error})
    }

    const schools = new Map();
    for (const course of courses) {
        const { data: school, error: schoolError } = await apiRequest('/schools/' + course.school_id);
        if (schoolError) {
            throw error(500, {message: schoolError.error})
        }
        schools.set(course.id, school);
    }

    return {courses, schools};
};