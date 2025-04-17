import {apiRequest} from "$lib/api/api.ts";
import {redirect, error, fail} from "@sveltejs/kit";
import axios from "axios";

export const load = async ({cookies, locals, params}) => {
    const session = locals.session;
    if (!session) {
        throw redirect(302, '/auth/login')
    }

    const token = cookies.get('AuthorizationToken');
    const {data: lesson, error: err} = await apiRequest('/courses/' + params.course + '/lessons/' + params.lesson,
        'get', undefined, token);
    if (err) {
        throw error(500, {message: err.error})
    }

    if (lesson.type === 'theory') {
        lesson.theory_url = lesson.theory_url.replace('minio', 'localhost');
        const {data: theory} = await axios.get(lesson.theory_url);
        return {lesson, theory, token};
    } else if (lesson.type === 'practice') {
        if (lesson.tests.length == 0) {
            throw error(500);
        }
        lesson.tests[0].task_url = lesson.tests[0].task_url.replace('minio', 'localhost');
        const {data: task} = await axios.get(lesson.tests[0].task_url);
        return {lesson, task, token};
    }
    return {lesson, token};
};

export const actions = {
    editTheoryLesson: async ({request, cookies, params}) => {
        const formData = await request.formData();
        const title = String(formData.get('title'));
        const theory = String(formData.get('theory'));

        if (title == '') {
            return fail(400, {
                error: 'title cannot be empty',
            })
        }

        const token = cookies.get('AuthorizationToken');
        const {error} = await apiRequest('/courses/' + params.course + '/lessons/' + params.lesson, 'patch', {
            title: title,
            theory: theory,
        }, token)

        if (error) {
            return fail(error.status, {
                error: error.error,
                title: title,
                theory: theory,
            })
        }

        throw redirect(302, '/me/schools/' + params.school + '/' + params.course);
    },

    editPracticeLesson: async ({request, cookies, params}) => {
        const formData = await request.formData();
        const title = String(formData.get('title'));
        const task = String(formData.get('task'));
        const answer = formData.get('answer');
        const options = formData.getAll('options[]');

        if (answer == null) {
            return fail(400, {
                error: 'select correct answer',
            })
        }

        if (title == '') {
            return fail(400, {
                error: 'title cannot be empty',
            })
        }

        const token = cookies.get('AuthorizationToken');
        const {error} = await apiRequest('/courses/' + params.course + '/lessons/' + params.lesson, 'patch', {
            title: title,
            tests: [{
                task: task,
                options: options,
                answer: answer,
                level: 1,
                score: 1,
            }],
        }, token)

        if (error) {
            return fail(error.status, {
                error: error.error,
                title: title,
                task: task,
                options: options,
                answer: answer,
            })
        }

        throw redirect(302, '/me/schools/' + params.school + '/' + params.course);
    },

    editVideoLesson: async ({request, cookies, params}) => {
        const formData = await request.formData();
        const title = String(formData.get('title'));
        const videoUrl = String(formData.get('video_url'));
        console.log(videoUrl);

        if (title == '') {
            return fail(400, {
                error: 'title cannot be empty',
            })
        }

        const token = cookies.get('AuthorizationToken');
        const {error} = await apiRequest('/courses/' + params.course + '/lessons/' + params.lesson, 'patch', {
            title: title,
            video_url: videoUrl,
        }, token)

        if (error) {
            return fail(error.status, {
                error: error.error,
                title: title,
                video_url: videoUrl,
            })
        }

        throw redirect(302, '/me/schools/' + params.school + '/' + params.course);
    },
}
