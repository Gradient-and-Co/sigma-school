<script lang="ts">
	import type {CourseBrief, LessonBrief, School} from "$lib/types.ts";

	export let form;
	export let data;

    import { getToastStore, type ToastSettings } from '@skeletonlabs/skeleton';
	import {apiRequest} from "$lib/api/api.ts";
	import {error} from "@sveltejs/kit";
    const toastStore = getToastStore();
	const course = data.course as CourseBrief;
	const school = data.school as School;

    function toastCreateError(error: string): void {
        const t: ToastSettings = {
            message: error,
            background: 'variant-filled-error',
        };
        toastStore.trigger(t);
    }

    if (form?.error) {
		toastCreateError(form?.error);
    }

	let newLessonTitle = '';
	let newLessonType;

	async function createLesson() {
		if (newLessonTitle == '' || newLessonType == '') {
			toastCreateError('lesson title and type cannot be empty');
		}

		const token = data.token;
		const {data: lesson, error: err} = await apiRequest('/courses/' + course.id + '/lessons', 'post', {
			title: newLessonTitle,
			type: newLessonType,
			score: 1,
			theory: '',
			video_url: '',
			tests: [{
				task: "Given equation 2 * x + 1 = 0. Find x.",
				options: ["2", "-0.5", "1", "no solution"],
				answer: "-0.5",
				level: 1,
				score: 1
			}]
		}, token);
		if (err) {
			throw error(500, {message: err.error})
		} else {
			const lessonBrief: LessonBrief = {
				id: lesson.id,
				title: lesson.title,
				type: lesson.type,
			};

			course.lessons = [...course.lessons, lessonBrief];
			newLessonTitle = '';
			newLessonType = 'THEORY';
		}
	}

	async function deleteLesson(id: string) {
		const token = data.token;
		const {error: err} = await apiRequest('/courses/' + course.id + '/lessons/' + id, 'delete',
				undefined, token);
		if (err) {
			throw error(500, {message: err.error})
		} else {
			for (let i = 0; i < course.lessons.length; i++) {
				if (course.lessons[i].id === id) {
					course.lessons.splice(i, 1);
					course.lessons = [...course.lessons];
					return;
				}
			}
		}
	}
</script>

<div class="flex justify-center h-full pt-5">
	<div class="card p-4 w-modal shadow-xl space-y-4">
		<form class="modal-form border border-surface-500 p-4 space-y-4 rounded-container-token"
		      method="POST" action="?/editCourse">
			<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
				<div class="col-span-1">
					<label class="label">
						<input class="input" type="text" name="name" value={form?.name ?? course.name}
							   placeholder="Course name..." />
					</label>
				</div>
				<div class="col-span-1">
					<label class="label mx-1">
						<select class="select variant-ghost-surface" name="language" value={form?.language ?? course.language}>
							<option value="english">English</option>
							<option value="russian">Russian</option>
						</select>
					</label>
				</div>
				<div class="col-span-2">
					<label class="label">
						<input class="input" type="number" name="price" value={form?.price ?? course.price}
							   placeholder="Course price..." />
					</label>
				</div>
				<div class="col-span-2">
					<label class="label">
						<textarea class="textarea" name="description" rows="3" placeholder="Course description...">{form?.description ?? course.description}</textarea>
					</label>
				</div>
				{#each course.lessons as lesson}
					<div class="col-span-2 variant-ghost-surface card w-modal shadow-xl p-3 flex justify-between items-center">
						<span class="text-1xl">{lesson.title}</span>
						<div class="flex justify-center">
							<label class="label mx-1">
								<select class="select variant-ghost-surface" value={lesson.type}>
									<option value="theory">THEORY</option>
									<option value="practice">PRACTICE</option>
									<option value="video">VIDEO</option>
								</select>
							</label>
							<a class="btn btn-md variant-ghost-surface mx-1" href={'/me/schools/' + school.id + '/' + course.id + '/' + lesson.id}>EDIT</a>
							<button type="button" class="btn btn-md variant-ghost-error mx-1" on:click={() => deleteLesson(lesson.id)}>✖</button>
						</div>
					</div>
				{/each}
				<div class="col-span-2 variant-ghost-surface card w-modal shadow-xl p-3 flex justify-between items-center">
					<input class="input w-1/2" type="text" name="description"
						   placeholder="Enter new lesson title..." bind:value={newLessonTitle} />
					<div class="flex justify-center">
						<label class="label mx-1">
							<select class="select variant-ghost-surface" bind:value={newLessonType}>
								<option value="theory">THEORY</option>
								<option value="practice">PRACTICE</option>
								<option value="video">VIDEO</option>
							</select>
						</label>
						<button type="button" class="btn btn-md variant-ghost-surface mx-1" on:click={createLesson}>CREATE</button>
					</div>
				</div>
			</div>
			<button class="input btn btn-md variant-ghost-primary" type="submit">SAVE</button>
		</form>
	</div>
</div>

<style>
    :global(body) {
        display: flex;
        flex-direction: column;
        min-height: 100vh;
    }
</style>
