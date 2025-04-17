<script lang="ts">
	import type {CourseBrief} from "$lib/types.ts";

    export let data;
    const course = data.course as CourseBrief;
</script>

<div class="container h-full mx-auto w-5/6 text-token py-3 pt-7">
	<div class="grid grid-cols-12 gap-4">
		<div class="col-span-12 md:col-span-9 gap-4 pt-7">
			<div class="text-4xl">
				<h1>{course.name}</h1>
			</div>
			<p class="text-1xl">{course.description}</p>
			<div class="pt-5">
				<span class="text-2xl">Language: </span>
				<span class="text-1xl">{course.language}</span>
			</div>
			<div class="pt-5">
				<h1 class="text-2xl">Lessons</h1>
				<ol class="list list-decimal">
					{#each course.lessons as lesson, i}
						<li>
							<span>{i + 1}.</span>
							<span class="text-1xl">{lesson.title}</span>
						</li>
					{/each}
				</ol>
			</div>
		</div>
		<div class="col-span-3 md:col-span-3 gap-4">
			<img src={course.image_url} class="bg-black/50 w-full aspect-[9/9]" alt="course avatar" />
			<div class="flex-auto flex justify-center items-center pt-5">
				<h6 class="font-bold p-2" data-toc-ignore>{course.price} RUB</h6>
				<a class="btn btn-md variant-ghost-surface px-10" href={'/catalog/' + course.id + '/payment'}>
					BUY
				</a>
			</div>
		</div>
	</div>
</div>
