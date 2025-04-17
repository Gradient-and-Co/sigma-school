import {error} from "@sveltejs/kit";
import {apiRequest} from "$lib/api/api.ts";

export const load = async () => {
    const {data: courses, error: err} = await apiRequest('/courses');
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