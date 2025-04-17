import {error, redirect} from "@sveltejs/kit";
import {apiRequest} from "$lib/api/api.ts";

export const load = async ({ locals, params, cookies}) => {
    const session = locals.session;
    if (!session) {
        throw redirect(302, '/auth/login')
    } else {
        const {data: course, error: err} = await apiRequest('/courses/' + params.id + '/brief');
        if (err) {
            throw error(500, {message: err.error})
        }

        const token = cookies.get('AuthorizationToken');
        if (course.price === 0) {
            const {error: err} = await apiRequest('/users/me/courses/' + params.id, 'put', undefined, token);
            if (err) {
                throw error(500, {message: err.error})
            }
        } else {
            const {data: url, error: paymentError} = await apiRequest('/payment/courses/' + params.id, 'get', undefined, token);
            if (paymentError) {
                throw error(500, {message: paymentError.error})
            }
            throw redirect(302, url);
        }
        throw redirect(302, '/me/courses');
    }
};
