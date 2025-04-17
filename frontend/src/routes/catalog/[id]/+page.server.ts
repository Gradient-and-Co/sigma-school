import {error} from "@sveltejs/kit";
import {apiRequest} from "$lib/api/api.ts";

export const load = async ({ params}) => {
    const {data: course, error: err} = await apiRequest('/courses/' + params.id + '/brief');
    if (err) {
        throw error(500, {message: err.error})
    }

    return {course};
};