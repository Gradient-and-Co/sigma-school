<script lang="ts">
	import type {CourseBrief, Lesson, LessonStat} from "$lib/types.ts";
	import {Marked} from "marked";
	import {markedHighlight} from "marked-highlight";
	import hljs from "highlight.js";
	import {apiRequest} from "$lib/api/api.ts";
	import {error} from "@sveltejs/kit";
	import type {ToastSettings} from "@skeletonlabs/skeleton";
	import { getToastStore } from '@skeletonlabs/skeleton';
	const toastStore = getToastStore();

	const marked = new Marked(markedHighlight({
		emptyLangClass: 'hljs',
		langPrefix: 'hljs language-',
		highlight(code, lang) {
			const language = hljs.getLanguage(lang) ? lang : 'plaintext';
			return hljs.highlight(code, { language }).value;
		}
	}));

	export let data;

	const course = data.course as CourseBrief;
	const lessons = data.lessons as Lesson[];
	const stats = data.stats as LessonStat[];

	let currIndex: number = 0;

	let answer = '';

	function toastCreateError(error: string): void {
		const t: ToastSettings = {
			message: error,
			background: 'variant-filled-error',
		};
		toastStore.trigger(t);
	}

	async function submit(index: number) {
		if (lessons[index].type == 'practice' && answer == '') {
			toastCreateError('select test answer option');
			return;
		}

		const token = data.token;
		const {error: err} = await apiRequest('/courses/' + course.id +
				'/lessons/' + lessons[index].id + '/stat', 'post', {
			tests: [{
				test_id: lessons[index].tests[0]?.id,
				answer: answer,
			}]
		}, token);
		if (err) {
			throw error(500, {message: err.error})
		}

		const {data: stat, error: statError} = await apiRequest('/courses/' + course.id +
				'/lessons/' + lessons[index].id + '/stat', 'get', undefined, token);
		if (statError) {
			throw error(500, {message: statError.error})
		}
		stats[index] = stat;

		if (lessons[index].type == 'practice' && answer != lessons[index].tests[0].answer) {
			toastCreateError('incorrect answer, try again');
			return;
		}
	}
</script>

<div class="grid grid-cols-12">
	<div class="col-span-3 p-4 pb-20 space-y-4 overflow-y-auto">
		<section>
			<p class="font-bold pl-4 text-2xl mb-4">LESSONS</p>
			<nav class="list-nav">
				<ul>
					{#each course.lessons as l, index}
						{#if l.type === 'practice'}
							<li class:variant-glass-success={stats[index].tests[0].score !== 0}
								class:variant-filled-surface={stats[index].tests[0].score === 0 && currIndex === index}>
								<button class="btn btn-md px-10 w-full" on:click={() => currIndex = index}>
									<span class="flex-auto p-2">{index + 1}. {l.title}</span>
								</button>
							</li>
						{:else}
							<li class:variant-glass-success={stats[index].score !== 0}
								class:variant-filled-surface={stats[index].score === 0 && currIndex === index}>
								<button class="btn btn-md px-10 w-full" on:click={() => currIndex = index}>
									<span class="flex-auto p-2">{index + 1}. {l.title}</span>
								</button>
							</li>
						{/if}
					{/each}
				</ul>
			</nav>
		</section>
	</div>

	<div class="col-span-8 p-4 pb-20 space-y-4 overflow-y-auto">
		{#if lessons[currIndex].type === 'theory'}
			<h1 class="text-3xl">{lessons[currIndex].title}</h1>
			<div class="preview">{@html marked.parse(lessons[currIndex].markdown)}</div>
		{:else if lessons[currIndex].type === 'practice'}
			<h1 class="text-3xl">{lessons[currIndex].title}</h1>
			<div class="preview">{@html marked.parse(lessons[currIndex].markdown)}</div>
			<div>
				{#each lessons[currIndex].tests[0].options as option}
					<div class="variant-ghost-surface flex items-center px-2 my-2">
						<input class="radio" type="radio" name="answer" bind:group={answer} value={option}/>
						<p class="mx-2 my-2">{option}</p>
					</div>
				{/each}
			</div>
		{:else if lessons[currIndex].type === 'video'}
			<h1 class="text-3xl">{lessons[currIndex].title}</h1>
			<video autoplay controls>
				<source src={lessons[currIndex].video_url} type="video/mp4">
				<track src={lessons[currIndex].video_url} kind="captions">
				Your browser does not support the video tag.
			</video>
		{/if}

		<div class="flex justify-between items-center">
			<button type="button" class="btn btn-md variant-ghost-surface px-10 flex-end" on:click={() => submit(currIndex)}>
				SUBMIT
			</button>
		</div>
		<div class="flex justify-between items-center">
			<button disabled={currIndex - 1 < 0} type="button" class="btn btn-md variant-ghost-surface px-10 flex-end" on:click={() => currIndex -= 1}>
				PREVIOUS
			</button>
			<button disabled={currIndex + 1 >= lessons.length} type="button" class="btn btn-md variant-ghost-surface px-10" on:click={() => currIndex += 1}>
				NEXT
			</button>
		</div>
	</div>
</div>

<style>
	.btn {
		text-align: left; /* Aligns text to the left inside the button */
	}
	span {
		justify-content: flex-start; /* Ensures flex items align to the left */
	}

	:global(.hljs) {
		background-color: #2f2826;
		margin-top: 1rem;
		margin-bottom: 1rem;
	}

	:global(p) {
		margin-top: 1rem;
		margin-bottom: 1rem;
	}

	:global(h1) {
		font-size: 2rem;
		margin-top: 1rem;
		margin-bottom: 1rem;
	}

	:global(h2) {
		font-size: 1.5rem;
		margin-top: 1rem;
		margin-bottom: 1rem;
	}

	:global(h3) {
		font-size: 1.2rem;
		margin-top: 1rem;
		margin-bottom: 1rem;
	}

	:global(h4) {
		font-size: 1.0rem;
		margin-top: 1rem;
		margin-bottom: 1rem;
	}
</style>